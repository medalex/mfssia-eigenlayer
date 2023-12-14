package operator

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/medalex/mfssia-eigenlayer/aggregator"
	cstaskmanager "github.com/medalex/mfssia-eigenlayer/contracts/bindings/MfssiaTaskManager"
	"github.com/medalex/mfssia-eigenlayer/core"
	"github.com/medalex/mfssia-eigenlayer/core/chainio"
	"github.com/medalex/mfssia-eigenlayer/metrics"
	"github.com/medalex/mfssia-eigenlayer/types"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signer"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const AVS_NAME = "incredible-squaring"
const SEM_VER = "0.0.1"

const SYSTEM_1_URL = ""
const SYSTEM_2_URL = ""
const DKG_URL = ""

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.EthClient
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	blsKeypair       *bls.KeyPair
	operatorId       bls.OperatorId
	operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *cstaskmanager.ContractMfssiaTaskManagerNewTaskCreated
	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	mfssiaServiceManagerAddr common.Address
}

func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	var logLevel logging.LogLevel

	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}

	logger, err := sdklogging.NewZapLogger(logLevel)

	if err != nil {
		return nil, err
	}

	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.EthClient

	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)

		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)

			return nil, err
		}

		ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)

		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)

			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)

		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)

			return nil, err
		}

		ethWsClient, err = eth.NewClient(c.EthWsUrl)

		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)

			return nil, err
		}
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")

	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}

	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)

	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)

		return nil, err
	}

	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethRpcClient.ChainID(context.Background())

	if err != nil {
		logger.Error("Cannot get chainId", "err", err)

		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	sgn, err := signer.NewPrivateKeyFromKeystoreSigner(c.EcdsaPrivateKeyStorePath, ecdsaKeyPassword, chainId)

	if err != nil {
		logger.Errorf("Cannot create signer", "err", err)

		return nil, err
	}

	avsWriter, err := chainio.NewAvsWriter(sgn, common.HexToAddress(c.AVSServiceManagerAddress),
		common.HexToAddress(c.BLSOperatorStateRetrieverAddress), ethRpcClient, logger,
	)

	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)

		return nil, err
	}

	avsServiceBindings, err := chainio.NewAvsServiceBindings(
		common.HexToAddress(c.AVSServiceManagerAddress),
		common.HexToAddress(c.BLSOperatorStateRetrieverAddress),
		ethRpcClient,
		logger,
	)

	if err != nil {
		return nil, err
	}

	blsRegistryCoordinatorAddr, err := avsServiceBindings.ServiceManager.RegistryCoordinator(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}

	stakeRegistryAddr, err := avsServiceBindings.ServiceManager.StakeRegistry(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}

	blsPubkeyRegistryAddr, err := avsServiceBindings.ServiceManager.BlsPubkeyRegistry(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}

	avsRegistryContractClient, err := sdkclients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordinatorAddr, common.HexToAddress(c.BLSOperatorStateRetrieverAddress), stakeRegistryAddr, blsPubkeyRegistryAddr, ethRpcClient, logger,
	)

	if err != nil {
		return nil, err
	}

	avsRegistryReader, err := sdkavsregistry.NewAvsRegistryReader(avsRegistryContractClient, logger, ethRpcClient)

	if err != nil {
		return nil, err
	}

	avsReader, err := chainio.NewAvsReader(avsRegistryReader, avsServiceBindings, logger)

	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	avsSubscriber, err := chainio.NewAvsSubscriber(common.HexToAddress(c.AVSServiceManagerAddress),
		common.HexToAddress(c.BLSOperatorStateRetrieverAddress), ethWsClient, logger,
	)

	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	slasherAddr, err := avsReader.AvsServiceBindings.ServiceManager.Slasher(&bind.CallOpts{})

	if err != nil {
		logger.Error("Cannot get slasher address", "err", err)
		return nil, err
	}

	elContractsClient, err := sdkclients.NewELContractsChainClient(slasherAddr, common.HexToAddress(c.BlsPublicKeyCompendiumAddress), ethRpcClient, ethWsClient, logger)

	if err != nil {
		logger.Error("Cannot create ELContractsChainClient", "err", err)
		return nil, err
	}

	eigenlayerWriter := sdkelcontracts.NewELChainWriter(elContractsClient, ethRpcClient, sgn, logger, eigenMetrics)

	if err != nil {
		logger.Error("Cannot create EigenLayerWriter", "err", err)
		return nil, err
	}

	eigenlayerReader, err := sdkelcontracts.NewELChainReader(elContractsClient, logger, ethRpcClient)

	if err != nil {
		logger.Error("Cannot create EigenLayerReader", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}

	economicMetricsCollector := economic.NewCollector(eigenlayerReader, avsRegistryReader, AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)

	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	operator := &Operator{
		config:                     c,
		logger:                     logger,
		metricsReg:                 reg,
		metrics:                    avsAndEigenMetrics,
		nodeApi:                    nodeApi,
		ethClient:                  ethRpcClient,
		avsWriter:                  avsWriter,
		avsReader:                  avsReader,
		avsSubscriber:              avsSubscriber,
		eigenlayerReader:           eigenlayerReader,
		eigenlayerWriter:           eigenlayerWriter,
		blsKeypair:                 blsKeyPair,
		operatorAddr:               common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr: c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:        aggregatorRpcClient,
		newTaskCreatedChan:         make(chan *cstaskmanager.ContractMfssiaTaskManagerNewTaskCreated),
		mfssiaServiceManagerAddr:   common.HexToAddress(c.AVSServiceManagerAddress),
		operatorId:                 [32]byte{0}, // this is set below
	}

	if c.RegisterOperatorOnStartup {
		operator.registerOperatorOnStartup(common.HexToAddress(c.BlsPublicKeyCompendiumAddress))
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := avsRegistryReader.GetOperatorId(context.Background(), operator.operatorAddr)

	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	operator.operatorId = operatorId

	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil
}

func (o *Operator) Start(ctx context.Context) error {
	operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(ctx, o.operatorAddr)

	if err != nil {
		o.logger.Error("Error checking if operator is registered", "err", err)
		return err
	}

	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
		// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
		return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}

	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}

	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub := o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.logger.Fatal("Error in metrics server", "err", err)
		case err := <-sub.Err():
			o.logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub = o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
		case newTaskCreatedLog := <-o.newTaskCreatedChan:
			o.metrics.IncNumTasksReceived()
			taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
		}
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractMfssiaTaskManagerNewTaskCreated) *cstaskmanager.IMfssiaTaskManagerTaskResponse {
	o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.logger.Info("Received new task",
		"system 1 value", newTaskCreatedLog.Task.System1Value,
		"system 2 value", newTaskCreatedLog.Task.System2Value,
		"dkg value", newTaskCreatedLog.Task.DkgValue,
		"taskIndex", newTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	var system1 = GetSystem1()
	var system2 = GetSystem2()
	var dkg = GetDkg()
	faultySystem := ""

	if dkg != system1 && dkg == system2 {
		faultySystem = "system1"
	}

	if dkg != system2 && dkg == system1 {
		faultySystem = "system2"
	}

	if dkg != system1 && dkg != system2 {
		faultySystem = "ALL"
	}

	taskResponse := &cstaskmanager.IMfssiaTaskManagerTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		FaultySystem:       faultySystem,
	}

	return taskResponse
}

func GetSystem1() string {
	res, err := http.Get(SYSTEM_1_URL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return string(resBody)
}

func GetSystem2() string {
	res, err := http.Get(SYSTEM_2_URL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return string(resBody)
}

func GetDkg() string {
	res, err := http.Get(DKG_URL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return string(resBody)
}

func (o *Operator) SignTaskResponse(taskResponse *cstaskmanager.IMfssiaTaskManagerTaskResponse) (*aggregator.SignedTaskResponse, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}
	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedTaskResponse := &aggregator.SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	erc20mock "github.com/medalex/mfssia-eigenlayer/contracts/bindings/ERC20Mock"
	csservicemanager "github.com/medalex/mfssia-eigenlayer/contracts/bindings/MfssiaServiceManager"
	cstaskmanager "github.com/medalex/mfssia-eigenlayer/contracts/bindings/MfssiaTaskManager"
)

type AvsServiceBindings struct {
	TaskManager    *cstaskmanager.ContractMfssiaTaskManager
	ServiceManager *csservicemanager.ContractMfssiaServiceManager
	ethClient      eth.EthClient
	logger         logging.Logger
}

func NewAvsServiceBindings(serviceManagerAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.EthClient, logger logging.Logger) (*AvsServiceBindings, error) {
	contractServiceManager, err := csservicemanager.NewContractMfssiaServiceManager(serviceManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}

	taskManagerAddr, err := contractServiceManager.MfssiaTaskManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}
	contractTaskManager, err := cstaskmanager.NewContractMfssiaTaskManager(taskManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IIncredibleSquaringTaskManager contract", "err", err)
		return nil, err
	}

	return &AvsServiceBindings{
		ServiceManager: contractServiceManager,
		TaskManager:    contractTaskManager,
		ethClient:      ethclient,
		logger:         logger,
	}, nil
}

func (b *AvsServiceBindings) GetErc20Mock(tokenAddr common.Address) (*erc20mock.ContractERC20Mock, error) {
	contractErc20Mock, err := erc20mock.NewContractERC20Mock(tokenAddr, b.ethClient)
	if err != nil {
		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return contractErc20Mock, nil
}

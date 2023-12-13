// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMfssiaTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BLSOperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type BLSOperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// BLSOperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type BLSOperatorStateRetrieverOperator struct {
	OperatorId [32]byte
	Stake      *big.Int
}

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IMfssiaTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTask struct {
	TaskCreatedBlock          uint32
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	System1Value              string
	System2Value              string
	DkgValue                  string
}

// IMfssiaTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	FaultySystem       string
}

// IMfssiaTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// ContractMfssiaTaskManagerMetaData contains all meta data concerning the ContractMfssiaTaskManager contract.
var ContractMfssiaTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsPubkeyRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSPubkeyRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"system1Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"system2Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dkgValue\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBLSOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"system1Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"system2Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dkgValue\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"faultySystem\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"system1Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"system2Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dkgValue\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"faultySystem\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"system1Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"system2Value\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dkgValue\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"faultySystem\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200566538038062005665833981016040819052620000359162000169565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001b0565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001b0565b6001600160a01b031660c0525063ffffffff1660e05250620001d7565b6001600160a01b03811681146200016657600080fd5b50565b600080604083850312156200017d57600080fd5b82516200018a8162000150565b602084015190925063ffffffff81168114620001a557600080fd5b809150509250929050565b600060208284031215620001c357600080fd5b8151620001d08162000150565b9392505050565b60805160a05160c05160e05161541c620002496000396000818161026401528181610509015261074e015260008181610319015281816115f00152611f0c0152600081816103f90152818161250d01526126ae015260008181610420015281816113bb0152612361015261541c6000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c8063683048351161010f5780638da5cb5b116100a2578063f5c9899d11610071578063f5c9899d14610507578063f63c5bab1461052d578063f8c8765e14610535578063fabc1cbc1461054857600080fd5b80638da5cb5b146104af578063a11c1a5d146104c0578063cefdc1d4146104d3578063f2fde38b146104f457600080fd5b806372d18e8d116100de57806372d18e8d1461046b5780637afa1eed14610479578063886f11951461048c5780638b00ce7c1461049f57600080fd5b806368304835146103f45780636d14a9871461041b5780636efb463614610442578063715018a61461046357600080fd5b80633561deb111610187578063595c6a6711610156578063595c6a671461038e5780635ac86ab7146103965780635c975abb146103c95780635decc3f5146103d157600080fd5b80633561deb1146103145780633563b0d11461033b5780634e2c1cc71461035b5780634f739f741461036e57600080fd5b80631ad43189116101c35780631ad431891461025f578063245a7bfc1461029b5780632cb223d5146102c65780632d89f6fc146102f457600080fd5b80630c48e57b146101f557806310d67a2f1461020a578063136439dd1461021d578063171f1d5b14610230575b600080fd5b6102086102033660046142de565b61055b565b005b61020861021836600461437a565b6109dc565b61020861022b36600461439e565b610a8f565b61024361023e3660046143b7565b610bce565b6040805192151583529015156020830152015b60405180910390f35b6102867f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610256565b609b546102ae906001600160a01b031681565b6040516001600160a01b039091168152602001610256565b6102e66102d4366004614408565b60996020526000908152604090205481565b604051908152602001610256565b6102e6610302366004614408565b60986020526000908152604090205481565b6102ae7f000000000000000000000000000000000000000000000000000000000000000081565b61034e610349366004614425565b610d58565b604051610256919061456c565b61020861036936600461457f565b6110d1565b61038161037c366004614660565b61178a565b6040516102569190614764565b610208611e0e565b6103b96103a436600461481f565b606654600160ff9092169190911b9081161490565b6040519015158152602001610256565b6066546102e6565b6103b96103df366004614408565b609a6020526000908152604090205460ff1681565b6102ae7f000000000000000000000000000000000000000000000000000000000000000081565b6102ae7f000000000000000000000000000000000000000000000000000000000000000081565b610455610450366004614842565b611ed5565b604051610256929190614902565b610208612941565b60975463ffffffff16610286565b609c546102ae906001600160a01b031681565b6065546102ae906001600160a01b031681565b6097546102869063ffffffff1681565b6033546001600160a01b03166102ae565b6102086104ce36600461494b565b612955565b6104e66104e1366004614a25565b612baf565b604051610256929190614a67565b61020861050236600461437a565b612d41565b7f0000000000000000000000000000000000000000000000000000000000000000610286565b610286606481565b610208610543366004614a80565b612db7565b61020861055636600461439e565b612f08565b609b546001600160a01b031633146105ba5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064015b60405180910390fd5b60006105c96020850185614408565b90503660006105db6040870187614adc565b909250905060006105f26040880160208901614408565b9050609860006106056020890189614408565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016106319190614b90565b60405160208183030381529060405280519060200120146106ba5760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e747261637400000060648201526084016105b1565b60006099816106cc60208a018a614408565b63ffffffff1663ffffffff16815260200190815260200160002054146107495760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016105b1565b6107737f000000000000000000000000000000000000000000000000000000000000000085614c83565b63ffffffff164363ffffffff1611156107e45760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016105b1565b6000866040516020016107f79190614ceb565b60405160208183030381529060405280519060200120905060008061081f8387878a8c611ed5565b9150915060005b8581101561091e578460ff168360200151828151811061084857610848614cfe565b602002602001015161085a9190614d14565b6001600160601b031660648460000151838151811061087b5761087b614cfe565b60200260200101516001600160601b03166108969190614d43565b101561090c576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016105b1565b8061091681614d62565b915050610826565b5060408051808201825263ffffffff4316815260208082018490529151909161094b918c91849101614d7d565b60405160208183030381529060405280519060200120609960008c60000160208101906109789190614408565b63ffffffff1663ffffffff168152602001908152602001600020819055507fd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb548a826040516109c7929190614d7d565b60405180910390a15050505050505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a539190614db0565b6001600160a01b0316336001600160a01b031614610a835760405162461bcd60e51b81526004016105b190614dcd565b610a8c81613064565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610ad7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afb9190614e17565b610b175760405162461bcd60e51b81526004016105b190614e39565b60665481811614610b905760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016105b1565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610c1657610c16614cfe565b60200201518951600160200201518a60200151600060028110610c3b57610c3b614cfe565b60200201518b60200151600160028110610c5757610c57614cfe565b602090810291909101518c518d830151604051610cb49a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610cd79190614e81565b9050610d4a610cf0610ce9888461315b565b86906131f2565b610cf8613286565b610d40610d3185610d2b604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061315b565b610d3a8c613346565b906131f2565b886201d4c06133d6565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbe9190614db0565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e249190614db0565b9050600085516001600160401b03811115610e4157610e41613ec3565b604051908082528060200260200182016040528015610e7457816020015b6060815260200190600190039081610e5f5790505b50905060005b86518110156110c6576000878281518110610e9757610e97614cfe565b016020015160405163889ae3e560e01b815260f89190911c6004820181905263ffffffff8916602483015291506000906001600160a01b0386169063889ae3e590604401600060405180830381865afa158015610ef8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f209190810190614ea3565b905080516001600160401b03811115610f3b57610f3b613ec3565b604051908082528060200260200182016040528015610f8057816020015b6040805180820190915260008082526020820152815260200190600190039081610f595790505b50848481518110610f9357610f93614cfe565b602002602001018190525060005b81518110156110b0576000828281518110610fbe57610fbe614cfe565b6020908102919091018101516040805180820182528281529051631b32722560e01b81526004810183905260ff8816602482015263ffffffff8e166044820152919350918201906001600160a01b038b1690631b32722590606401602060405180830381865afa158015611036573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105a9190614f33565b6001600160601b031681525086868151811061107857611078614cfe565b6020026020010151838151811061109157611091614cfe565b60200260200101819052505080806110a890614d62565b915050610fa1565b50505080806110be90614d62565b915050610e7a565b509695505050505050565b60006110e06020850185614408565b63ffffffff811660009081526099602052604090205490915061114f5760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b60648201526084016105b1565b8383604051602001611162929190614f5c565b60408051601f19818403018152918152815160209283012063ffffffff841660009081526099909352912054146112015760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e747261637400000060648201526084016105b1565b63ffffffff81166000908152609a602052604090205460ff16156112995760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a4016105b1565b60646112a86020850185614408565b6112b29190614c83565b63ffffffff164363ffffffff1611156113335760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e00000000000000000060648201526084016105b1565b600061133e866135fa565b9050600061134f6020870187614adc565b60405161135d929190614f9a565b6040519081900390208251602084012014905060018114156113b457604051339063ffffffff8516907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a3505050611784565b60006114297f00000000000000000000000000000000000000000000000000000000000000006113e760408b018b614adc565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103499250505060208c018c614408565b9050600085516001600160401b0381111561144657611446613ec3565b60405190808252806020026020018201604052801561146f578160200160208202803683370190505b50905060005b86518110156114cf576114a087828151811061149357611493614cfe565b6020026020010151613745565b8282815181106114b2576114b2614cfe565b6020908102919091010152806114c781614d62565b915050611475565b5060006114df60208b018b614408565b826040516020016114f1929190614faa565b6040516020818303038152906040528051906020012090508760200135811461159b5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a4016105b1565b600087516001600160401b038111156115b6576115b6613ec3565b6040519080825280602002602001820160405280156115df578160200160208202803683370190505b50905060005b8851811015611733577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663187548c86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561164c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116709190614db0565b6001600160a01b031663e8bb9ae685838151811061169057611690614cfe565b60200260200101516040518263ffffffff1660e01b81526004016116b691815260200190565b602060405180830381865afa1580156116d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f79190614db0565b82828151811061170957611709614cfe565b6001600160a01b03909216602092830291909101909101528061172b81614d62565b9150506115e5565b5063ffffffff87166000818152609a6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505050505b50505050565b6117b56040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118199190614db0565b90506118466040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516385020d4960e01b81526001600160a01b038a16906385020d4990611876908b9089908990600401614ff2565b600060405180830381865afa158015611893573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118bb919081019061503c565b815260405163e192e9ad60e01b81526001600160a01b0383169063e192e9ad906118ed908b908b908b906004016150ca565b600060405180830381865afa15801561190a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611932919081019061503c565b6040820152856001600160401b0381111561194f5761194f613ec3565b60405190808252806020026020018201604052801561198257816020015b606081526020019060019003908161196d5790505b50606082015260005b60ff8116871115611d1f576000856001600160401b038111156119b0576119b0613ec3565b6040519080825280602002602001820160405280156119d9578160200160208202803683370190505b5083606001518360ff16815181106119f3576119f3614cfe565b602002602001018190525060005b86811015611c1f5760008c6001600160a01b0316633064620d8a8a85818110611a2c57611a2c614cfe565b905060200201358e88600001518681518110611a4a57611a4a614cfe565b60200260200101516040518463ffffffff1660e01b8152600401611a879392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611aa4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ac891906150ea565b90508a8a8560ff16818110611adf57611adf614cfe565b6001600160c01b03841692013560f81c9190911c600190811614159050611c0c57856001600160a01b031663480858668a8a85818110611b2157611b21614cfe565b905060200201358d8d8860ff16818110611b3d57611b3d614cfe565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611b93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bb79190615113565b85606001518560ff1681518110611bd057611bd0614cfe565b60200260200101518481518110611be957611be9614cfe565b63ffffffff9092166020928302919091019091015282611c0881614d62565b9350505b5080611c1781614d62565b915050611a01565b506000816001600160401b03811115611c3a57611c3a613ec3565b604051908082528060200260200182016040528015611c63578160200160208202803683370190505b50905060005b82811015611ce45784606001518460ff1681518110611c8a57611c8a614cfe565b60200260200101518181518110611ca357611ca3614cfe565b6020026020010151828281518110611cbd57611cbd614cfe565b63ffffffff9092166020928302919091019091015280611cdc81614d62565b915050611c69565b508084606001518460ff1681518110611cff57611cff614cfe565b602002602001018190525050508080611d1790615130565b91505061198b565b506000896001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d849190614db0565b60405163eda1076360e01b81529091506001600160a01b0382169063eda1076390611db7908b908b908e90600401615150565b600060405180830381865afa158015611dd4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611dfc919081019061503c565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611e56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e7a9190614e17565b611e965760405162461bcd60e51b81526004016105b190614e39565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60408051808201909152606080825260208201526040805180820190915260008082526020820181905290815b868110156120f2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c1af6b24898984818110611f4b57611f4b614cfe565b9050013560f81c60f81b60f81c888860a001518581518110611f6f57611f6f614cfe565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611fcb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fef919061517a565b6001600160401b0319166120128660400151838151811061149357611493614cfe565b67ffffffffffffffff1916146120ae5760405162461bcd60e51b815260206004820152606160248201526000805160206153c783398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016105b1565b6120de856040015182815181106120c7576120c7614cfe565b6020026020010151836131f290919063ffffffff16565b9150806120ea81614d62565b915050611f02565b506040805180820190915260608082526020820152866001600160401b0381111561211f5761211f613ec3565b604051908082528060200260200182016040528015612148578160200160208202803683370190505b506020820152866001600160401b0381111561216657612166613ec3565b60405190808252806020026020018201604052801561218f578160200160208202803683370190505b5081526020850151516000906001600160401b038111156121b2576121b2613ec3565b6040519080825280602002602001820160405280156121db578160200160208202803683370190505b50905060008660200151516001600160401b038111156121fd576121fd613ec3565b604051908082528060200260200182016040528015612226578160200160208202803683370190505b509050600061226a8b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061378892505050565b905060005b8860200151518110156124d5576122958960200151828151811061149357611493614cfe565b8482815181106122a7576122a7614cfe565b6020908102919091010152801561235f57836122c46001836151a5565b815181106122d4576122d4614cfe565b602002602001015160001c8482815181106122f1576122f1614cfe565b602002602001015160001c1161235f576040805162461bcd60e51b81526020600482015260248101919091526000805160206153c783398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016105b1565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633064620d8583815181106123a0576123a0614cfe565b60200260200101518c8c6000015185815181106123bf576123bf614cfe565b60200260200101516040518463ffffffff1660e01b81526004016123fc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612419573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061243d91906150ea565b6001600160c01b031683828151811061245857612458614cfe565b6020026020010181815250506124c16124ba61248e8486858151811061248057612480614cfe565b6020026020010151166138f1565b6124b48c6020015185815181106124a7576124a7614cfe565b6020026020010151613922565b906139bd565b87906131f2565b9550806124cd81614d62565b91505061226f565b505060005b60ff81168a11156128155760008b8b8360ff168181106124fc576124fc614cfe565b9050013560f81c60f81b60f81c90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c56828c8c60c001518660ff168151811061255557612555614cfe565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156125b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125d59190614f33565b85602001518360ff16815181106125ee576125ee614cfe565b6001600160601b03909216602092830291909101820152850151805160ff841690811061261d5761261d614cfe565b602002602001015185600001518360ff168151811061263e5761263e614cfe565b60200260200101906001600160601b031690816001600160601b03168152505060005b8960200151518163ffffffff16101561280b5760006126a7858363ffffffff168151811061269157612691614cfe565b60200260200101518460ff161c60019081161490565b156127f8577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a43cde89848e898663ffffffff16815181106126f5576126f5614cfe565b60200260200101518f60e001518960ff168151811061271657612716614cfe565b60200260200101518663ffffffff168151811061273557612735614cfe565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612799573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127bd9190614f33565b8751805160ff87169081106127d4576127d4614cfe565b602002602001018181516127e891906151bc565b6001600160601b03169052506001015b5080612803816151e4565b915050612661565b50506001016124da565b505060008061282e8c868a606001518b60800151610bce565b915091508161289f5760405162461bcd60e51b815260206004820152604360248201526000805160206153c783398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016105b1565b806129005760405162461bcd60e51b815260206004820152603960248201526000805160206153c783398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016105b1565b505060008782604051602001612917929190614faa565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612949613aa2565b6129536000613afc565b565b609c546001600160a01b031633146129b95760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084016105b1565b6129fe6040518060c00160405280600063ffffffff168152602001600063ffffffff168152602001606081526020016060815260200160608152602001606081525090565b89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506060820152604080516020601f8a018190048102820181019092528881529089908990819084018382808284376000920191909152505050506080820152604080516020601f880181900481028201810190925286815290879087908190840183828082843760009201919091525050505060a082015263ffffffff4381168252841660208083019190915260408051601f850183900483028101830190915283815290849084908190840183828082843760009201919091525050505060408083019190915251612b0a908290602001615260565b60408051601f1981840301815282825280516020918201206097805463ffffffff9081166000908152609890945293909220555416907f02ce4afaf544492edb45b478f1e8f27989b08d79fb36777a33affaf90366be9490612b6d908490615260565b60405180910390a2609754612b899063ffffffff166001614c83565b6097805463ffffffff191663ffffffff9290921691909117905550505050505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612bea57612bea614cfe565b60209081029190910101526040516385020d4960e01b81526000906001600160a01b038816906385020d4990612c2690889086906004016152f3565b600060405180830381865afa158015612c43573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612c6b919081019061503c565b600081518110612c7d57612c7d614cfe565b6020908102919091010151604051633064620d60e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b03891690633064620d90606401602060405180830381865afa158015612ce9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d0d91906150ea565b6001600160c01b031690506000612d2382613b4e565b905081612d318a838a610d58565b9550955050505050935093915050565b612d49613aa2565b6001600160a01b038116612dae5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105b1565b610a8c81613afc565b600054610100900460ff1615808015612dd75750600054600160ff909116105b80612df15750303b158015612df1575060005460ff166001145b612e545760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105b1565b6000805460ff191660011790558015612e77576000805461ff0019166101001790555b612e82856000613bab565b612e8b84613afc565b609b80546001600160a01b038086166001600160a01b031992831617909255609c8054928516929091169190911790558015612f01576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f7f9190614db0565b6001600160a01b0316336001600160a01b031614612faf5760405162461bcd60e51b81526004016105b190614dcd565b60665419811960665419161461302d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016105b1565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610bc3565b6001600160a01b0381166130f25760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016105b1565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613177613dbf565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156131aa576131ac565bfe5b50806131ea5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016105b1565b505092915050565b604080518082019091526000808252602082015261320e613ddd565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156131aa5750806131ea5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016105b1565b61328e613dfb565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806133766000805160206153a783398151915286614e81565b90505b61338281613c95565b90935091506000805160206153a78339815191528283098314156133bc576040805180820190915290815260208101919091529392505050565b6000805160206153a7833981519152600182089050613379565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613408613e20565b60005b60028110156135cd576000613421826006614d43565b905084826002811061343557613435614cfe565b60200201515183613447836000615347565b600c811061345757613457614cfe565b602002015284826002811061346e5761346e614cfe565b602002015160200151838260016134859190615347565b600c811061349557613495614cfe565b60200201528382600281106134ac576134ac614cfe565b60200201515151836134bf836002615347565b600c81106134cf576134cf614cfe565b60200201528382600281106134e6576134e6614cfe565b60200201515160016020020151836134ff836003615347565b600c811061350f5761350f614cfe565b602002015283826002811061352657613526614cfe565b60200201516020015160006002811061354157613541614cfe565b602002015183613552836004615347565b600c811061356257613562614cfe565b602002015283826002811061357957613579614cfe565b60200201516020015160016002811061359457613594614cfe565b6020020151836135a5836005615347565b600c81106135b5576135b5614cfe565b602002015250806135c581614d62565b91505061340b565b506135d6613e3f565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6060600061360b60a0840184614adc565b604051613619929190614f9a565b60405190819003902061362f6060850185614adc565b60405161363d929190614f9a565b604051908190039020149050600061365860a0850185614adc565b604051613666929190614f9a565b60405190819003902061367c6080860186614adc565b60405161368a929190614f9a565b60405180910390201490506060826136c0575060408051808201909152600781526673797374656d3160c81b602082015261373d565b816136e9575060408051808201909152600781526639bcb9ba32b69960c91b602082015261373d565b60405162461bcd60e51b815260206004820152602360248201527f426f74682073797374656d7320646f206d617463682074686520444b472076616044820152626c756560e81b60648201526084016105b1565b949350505050565b60008160000151826020015160405160200161376b929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b6000610100825111156137fc5760405162461bcd60e51b815260206004820152603660248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a206044820152756279746573417272617920697320746f6f206c6f6e6760501b60648201526084016105b1565b815161380a57506000919050565b6000808360008151811061382057613820614cfe565b0160200151600160f89190911c81901b92505b84518110156138e85784818151811061384e5761384e614cfe565b0160200151600160f89190911c1b9150828216156138d45760405162461bcd60e51b815260206004820152603a60248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a2060448201527f72657065617420656e74727920696e206279746573417272617900000000000060648201526084016105b1565b918117916138e181614d62565b9050613833565b50909392505050565b6000805b821561391c576139066001846151a5565b90921691806139148161535f565b9150506138f5565b92915050565b6040805180820190915260008082526020820152815115801561394757506020820151155b15613965575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206153a783398151915284602001516139989190614e81565b6139b0906000805160206153a78339815191526151a5565b905292915050565b919050565b60408051808201909152600080825260208201526102008261ffff1610613a195760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016105b1565b8161ffff1660011415613a2d57508161391c565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161115613a9757600161ffff871660ff83161c81161415613a7a57613a7784846131f2565b93505b613a8483846131f2565b92506201fffe600192831b169101613a49565b509195945050505050565b6033546001600160a01b031633146129535760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105b1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b610100811015613ba4576001811b915083821615613b9457828160f81b604051602001613b82929190615377565b60405160208183030381529060405292505b613b9d81614d62565b9050613b54565b5050919050565b6065546001600160a01b0316158015613bcc57506001600160a01b03821615155b613c4e5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016105b1565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613c9182613064565b5050565b600080806000805160206153a783398151915260036000805160206153a7833981519152866000805160206153a7833981519152888909090890506000613d0b827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206153a7833981519152613d17565b91959194509092505050565b600080613d22613e3f565b613d2a613e5d565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156131aa575082613db45760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016105b1565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613e0e613e7b565b8152602001613e1b613e7b565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b600060c08284031215613eab57600080fd5b50919050565b600060408284031215613eab57600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613efb57613efb613ec3565b60405290565b60405161010081016001600160401b0381118282101715613efb57613efb613ec3565b604051601f8201601f191681016001600160401b0381118282101715613f4c57613f4c613ec3565b604052919050565b60006001600160401b03821115613f6d57613f6d613ec3565b5060051b60200190565b63ffffffff81168114610a8c57600080fd5b80356139b881613f77565b600082601f830112613fa557600080fd5b81356020613fba613fb583613f54565b613f24565b82815260059290921b84018101918181019086841115613fd957600080fd5b8286015b848110156110c6578035613ff081613f77565b8352918301918301613fdd565b60006040828403121561400f57600080fd5b614017613ed9565b9050813581526020820135602082015292915050565b600082601f83011261403e57600080fd5b8135602061404e613fb583613f54565b82815260069290921b8401810191818101908684111561406d57600080fd5b8286015b848110156110c6576140838882613ffd565b835291830191604001614071565b600082601f8301126140a257600080fd5b604051604081018181106001600160401b03821117156140c4576140c4613ec3565b80604052508060408401858111156140db57600080fd5b845b81811015613a975780358352602092830192016140dd565b60006080828403121561410757600080fd5b61410f613ed9565b905061411b8383614091565b815261412a8360408401614091565b602082015292915050565b600082601f83011261414657600080fd5b81356020614156613fb583613f54565b82815260059290921b8401810191818101908684111561417557600080fd5b8286015b848110156110c65780356001600160401b038111156141985760008081fd5b6141a68986838b0101613f94565b845250918301918301614179565b600061018082840312156141c757600080fd5b6141cf613f01565b905081356001600160401b03808211156141e857600080fd5b6141f485838601613f94565b8352602084013591508082111561420a57600080fd5b6142168583860161402d565b6020840152604084013591508082111561422f57600080fd5b61423b8583860161402d565b604084015261424d85606086016140f5565b606084015261425f8560e08601613ffd565b608084015261012084013591508082111561427957600080fd5b61428585838601613f94565b60a084015261014084013591508082111561429f57600080fd5b6142ab85838601613f94565b60c08401526101608401359150808211156142c557600080fd5b506142d284828501614135565b60e08301525092915050565b6000806000606084860312156142f357600080fd5b83356001600160401b038082111561430a57600080fd5b61431687838801613e99565b9450602086013591508082111561432c57600080fd5b61433887838801613eb1565b9350604086013591508082111561434e57600080fd5b5061435b868287016141b4565b9150509250925092565b6001600160a01b0381168114610a8c57600080fd5b60006020828403121561438c57600080fd5b813561439781614365565b9392505050565b6000602082840312156143b057600080fd5b5035919050565b60008060008061012085870312156143ce57600080fd5b843593506143df8660208701613ffd565b92506143ee86606087016140f5565b91506143fd8660e08701613ffd565b905092959194509250565b60006020828403121561441a57600080fd5b813561439781613f77565b60008060006060848603121561443a57600080fd5b833561444581614365565b92506020848101356001600160401b038082111561446257600080fd5b818701915087601f83011261447657600080fd5b81358181111561448857614488613ec3565b61449a601f8201601f19168501613f24565b915080825288848285010111156144b057600080fd5b80848401858401376000848284010152508094505050506144d360408501613f89565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b8681101561455e578385038a52825180518087529087019087870190845b81811015614549578351805184528a01516001600160601b03168a84015292890192604090920191600101614519565b50509a87019a955050918501916001016144fb565b509298975050505050505050565b60208152600061439760208301846144dc565b60008060008060a0858703121561459557600080fd5b84356001600160401b03808211156145ac57600080fd5b6145b888838901613e99565b955060208701359150808211156145ce57600080fd5b6145da88838901613eb1565b94506145e98860408901613eb1565b935060808701359150808211156145ff57600080fd5b5061460c8782880161402d565b91505092959194509250565b60008083601f84011261462a57600080fd5b5081356001600160401b0381111561464157600080fd5b60208301915083602082850101111561465957600080fd5b9250929050565b6000806000806000806080878903121561467957600080fd5b863561468481614365565b9550602087013561469481613f77565b945060408701356001600160401b03808211156146b057600080fd5b6146bc8a838b01614618565b909650945060608901359150808211156146d557600080fd5b818901915089601f8301126146e957600080fd5b8135818111156146f857600080fd5b8a60208260051b850101111561470d57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561475957815163ffffffff1687529582019590820190600101614737565b509495945050505050565b60006020808352835160808285015261478060a0850182614723565b905081850151601f198086840301604087015261479d8383614723565b925060408701519150808684030160608701526147ba8383614723565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561481157848783030184526147ff828751614723565b958801959388019391506001016147e5565b509998505050505050505050565b60006020828403121561483157600080fd5b813560ff8116811461439757600080fd5b60008060008060006080868803121561485a57600080fd5b8535945060208601356001600160401b038082111561487857600080fd5b61488489838a01614618565b90965094506040880135915061489982613f77565b909250606087013590808211156148af57600080fd5b506148bc888289016141b4565b9150509295509295909350565b600081518084526020808501945080840160005b838110156147595781516001600160601b0316875295820195908201906001016148dd565b604081526000835160408084015261491d60808401826148c9565b90506020850151603f1984830301606085015261493a82826148c9565b925050508260208301529392505050565b600080600080600080600080600060a08a8c03121561496957600080fd5b89356001600160401b038082111561498057600080fd5b61498c8d838e01614618565b909b50995060208c01359150808211156149a557600080fd5b6149b18d838e01614618565b909950975060408c01359150808211156149ca57600080fd5b6149d68d838e01614618565b909750955060608c013591506149eb82613f77565b90935060808b01359080821115614a0157600080fd5b50614a0e8c828d01614618565b915080935050809150509295985092959850929598565b600080600060608486031215614a3a57600080fd5b8335614a4581614365565b9250602084013591506040840135614a5c81613f77565b809150509250925092565b82815260406020820152600061373d60408301846144dc565b60008060008060808587031215614a9657600080fd5b8435614aa181614365565b93506020850135614ab181614365565b92506040850135614ac181614365565b91506060850135614ad181614365565b939692955090935050565b6000808335601e19843603018112614af357600080fd5b8301803591506001600160401b03821115614b0d57600080fd5b60200191503681900382131561465957600080fd5b6000808335601e19843603018112614b3957600080fd5b83016020810192503590506001600160401b03811115614b5857600080fd5b80360383131561465957600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815260008235614ba181613f77565b63ffffffff808216602085015260208501359150614bbe82613f77565b80821660408501525050614bd56040840184614b22565b60c06060850152614bea60e085018284614b67565b915050614bfa6060850185614b22565b601f1980868503016080870152614c12848385614b67565b9350614c216080880188614b22565b93509150808685030160a0870152614c3a848484614b67565b9350614c4960a0880188614b22565b93509150808685030160c087015250614c63838383614b67565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff808316818516808303821115614ca257614ca2614c6d565b01949350505050565b60008135614cb881613f77565b63ffffffff168352614ccd6020830183614b22565b60406020860152614ce2604086018284614b67565b95945050505050565b6020815260006143976020830184614cab565b634e487b7160e01b600052603260045260246000fd5b60006001600160601b0380831681851681830481118215151615614d3a57614d3a614c6d565b02949350505050565b6000816000190483118215151615614d5d57614d5d614c6d565b500290565b6000600019821415614d7657614d76614c6d565b5060010190565b606081526000614d906060830185614cab565b905063ffffffff8351166020830152602083015160408301529392505050565b600060208284031215614dc257600080fd5b815161439781614365565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614e2957600080fd5b8151801515811461439757600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b600082614e9e57634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614eb657600080fd5b82516001600160401b03811115614ecc57600080fd5b8301601f81018513614edd57600080fd5b8051614eeb613fb582613f54565b81815260059190911b82018301908381019087831115614f0a57600080fd5b928401925b82841015614f2857835182529284019290840190614f0f565b979650505050505050565b600060208284031215614f4557600080fd5b81516001600160601b038116811461439757600080fd5b606081526000614f6f6060830185614cab565b90508235614f7c81613f77565b63ffffffff8116602084015250602083013560408301529392505050565b8183823760009101908152919050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015614fe557815185529382019390820190600101614fc9565b5092979650505050505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561501f57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561504f57600080fd5b82516001600160401b0381111561506557600080fd5b8301601f8101851361507657600080fd5b8051615084613fb582613f54565b81815260059190911b820183019083810190878311156150a357600080fd5b928401925b82841015614f285783516150bb81613f77565b825292840192908401906150a8565b63ffffffff84168152604060208201526000614ce2604083018486614b67565b6000602082840312156150fc57600080fd5b81516001600160c01b038116811461439757600080fd5b60006020828403121561512557600080fd5b815161439781613f77565b600060ff821660ff81141561514757615147614c6d565b60010192915050565b604081526000615164604083018587614b67565b905063ffffffff83166020830152949350505050565b60006020828403121561518c57600080fd5b815167ffffffffffffffff198116811461439757600080fd5b6000828210156151b7576151b7614c6d565b500390565b60006001600160601b03838116908316818110156151dc576151dc614c6d565b039392505050565b600063ffffffff808316818114156151fe576151fe614c6d565b6001019392505050565b60005b8381101561522357818101518382015260200161520b565b838111156117845750506000910152565b6000815180845261524c816020860160208601615208565b601f01601f19169290920160200192915050565b60208152600063ffffffff80845116602084015280602085015116604084015250604083015160c0606084015261529a60e0840182615234565b90506060840151601f19808584030160808601526152b88383615234565b925060808601519150808584030160a08601526152d58383615234565b925060a08601519150808584030160c086015250614ce28282615234565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561533a5784518352938301939183019160010161531e565b5090979650505050505050565b6000821982111561535a5761535a614c6d565b500190565b600061ffff808316818114156151fe576151fe614c6d565b60008351615389818460208801615208565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220cd85e38a80c91cdd64325d212b2745ad52f04870b65e88f7bec33b11f886396c64736f6c634300080c0033",
}

// ContractMfssiaTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMfssiaTaskManagerMetaData.ABI instead.
var ContractMfssiaTaskManagerABI = ContractMfssiaTaskManagerMetaData.ABI

// ContractMfssiaTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMfssiaTaskManagerMetaData.Bin instead.
var ContractMfssiaTaskManagerBin = ContractMfssiaTaskManagerMetaData.Bin

// DeployContractMfssiaTaskManager deploys a new Ethereum contract, binding an instance of ContractMfssiaTaskManager to it.
func DeployContractMfssiaTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractMfssiaTaskManager, error) {
	parsed, err := ContractMfssiaTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMfssiaTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMfssiaTaskManager{ContractMfssiaTaskManagerCaller: ContractMfssiaTaskManagerCaller{contract: contract}, ContractMfssiaTaskManagerTransactor: ContractMfssiaTaskManagerTransactor{contract: contract}, ContractMfssiaTaskManagerFilterer: ContractMfssiaTaskManagerFilterer{contract: contract}}, nil
}

// ContractMfssiaTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractMfssiaTaskManager struct {
	ContractMfssiaTaskManagerCaller     // Read-only binding to the contract
	ContractMfssiaTaskManagerTransactor // Write-only binding to the contract
	ContractMfssiaTaskManagerFilterer   // Log filterer for contract events
}

// ContractMfssiaTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMfssiaTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMfssiaTaskManagerSession struct {
	Contract     *ContractMfssiaTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractMfssiaTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMfssiaTaskManagerCallerSession struct {
	Contract *ContractMfssiaTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractMfssiaTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMfssiaTaskManagerTransactorSession struct {
	Contract     *ContractMfssiaTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractMfssiaTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerRaw struct {
	Contract *ContractMfssiaTaskManager // Generic contract binding to access the raw methods on
}

// ContractMfssiaTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerCallerRaw struct {
	Contract *ContractMfssiaTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractMfssiaTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerTransactorRaw struct {
	Contract *ContractMfssiaTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMfssiaTaskManager creates a new instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManager(address common.Address, backend bind.ContractBackend) (*ContractMfssiaTaskManager, error) {
	contract, err := bindContractMfssiaTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManager{ContractMfssiaTaskManagerCaller: ContractMfssiaTaskManagerCaller{contract: contract}, ContractMfssiaTaskManagerTransactor: ContractMfssiaTaskManagerTransactor{contract: contract}, ContractMfssiaTaskManagerFilterer: ContractMfssiaTaskManagerFilterer{contract: contract}}, nil
}

// NewContractMfssiaTaskManagerCaller creates a new read-only instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractMfssiaTaskManagerCaller, error) {
	contract, err := bindContractMfssiaTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerCaller{contract: contract}, nil
}

// NewContractMfssiaTaskManagerTransactor creates a new write-only instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractMfssiaTaskManagerTransactor, error) {
	contract, err := bindContractMfssiaTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTransactor{contract: contract}, nil
}

// NewContractMfssiaTaskManagerFilterer creates a new log filterer instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractMfssiaTaskManagerFilterer, error) {
	contract, err := bindContractMfssiaTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerFilterer{contract: contract}, nil
}

// bindContractMfssiaTaskManager binds a generic wrapper to an already deployed contract.
func bindContractMfssiaTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMfssiaTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMfssiaTaskManager.Contract.ContractMfssiaTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.ContractMfssiaTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.ContractMfssiaTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMfssiaTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Aggregator(&_ContractMfssiaTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Aggregator(&_ContractMfssiaTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskHashes(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskHashes(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskResponses(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskResponses(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) BlsPubkeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "blsPubkeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.BlsPubkeyRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.BlsPubkeyRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.CheckSignatures(&_ContractMfssiaTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.CheckSignatures(&_ContractMfssiaTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Generator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Generator(&_ContractMfssiaTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Generator(&_ContractMfssiaTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(BLSOperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(BLSOperatorStateRetrieverCheckSignaturesIndices)).(*BLSOperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractMfssiaTaskManager.Contract.GetCheckSignaturesIndices(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractMfssiaTaskManager.Contract.GetCheckSignaturesIndices(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]BLSOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]BLSOperatorStateRetrieverOperator)).(*[][]BLSOperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]BLSOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]BLSOperatorStateRetrieverOperator)).(*[][]BLSOperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState0(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState0(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractMfssiaTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractMfssiaTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.LatestTaskNum(&_ContractMfssiaTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.LatestTaskNum(&_ContractMfssiaTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Owner() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Owner(&_ContractMfssiaTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Owner(&_ContractMfssiaTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.Paused(&_ContractMfssiaTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.Paused(&_ContractMfssiaTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractMfssiaTaskManager.Contract.Paused0(&_ContractMfssiaTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractMfssiaTaskManager.Contract.Paused0(&_ContractMfssiaTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.PauserRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.PauserRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.RegistryCoordinator(&_ContractMfssiaTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.RegistryCoordinator(&_ContractMfssiaTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.StakeRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.StakeRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TaskNumber(&_ContractMfssiaTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TaskNumber(&_ContractMfssiaTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractMfssiaTaskManager.Contract.TrySignatureAndApkVerification(&_ContractMfssiaTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractMfssiaTaskManager.Contract.TrySignatureAndApkVerification(&_ContractMfssiaTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa11c1a5d.
//
// Solidity: function createNewTask(string system1Value, string system2Value, string dkgValue, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, system1Value string, system2Value string, dkgValue string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "createNewTask", system1Value, system2Value, dkgValue, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa11c1a5d.
//
// Solidity: function createNewTask(string system1Value, string system2Value, string dkgValue, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) CreateNewTask(system1Value string, system2Value string, dkgValue string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.CreateNewTask(&_ContractMfssiaTaskManager.TransactOpts, system1Value, system2Value, dkgValue, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa11c1a5d.
//
// Solidity: function createNewTask(string system1Value, string system2Value, string dkgValue, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) CreateNewTask(system1Value string, system2Value string, dkgValue string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.CreateNewTask(&_ContractMfssiaTaskManager.TransactOpts, system1Value, system2Value, dkgValue, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Initialize(&_ContractMfssiaTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Initialize(&_ContractMfssiaTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Pause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Pause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.PauseAll(&_ContractMfssiaTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.PauseAll(&_ContractMfssiaTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4e2c1cc7.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,string,string,string) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4e2c1cc7.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,string,string,string) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RaiseAndResolveChallenge(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RaiseAndResolveChallenge(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4e2c1cc7.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,string,string,string) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) RaiseAndResolveChallenge(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RaiseAndResolveChallenge(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RenounceOwnership(&_ContractMfssiaTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RenounceOwnership(&_ContractMfssiaTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x0c48e57b.
//
// Solidity: function respondToTask((uint32,uint32,bytes,string,string,string) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x0c48e57b.
//
// Solidity: function respondToTask((uint32,uint32,bytes,string,string,string) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RespondToTask(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RespondToTask(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x0c48e57b.
//
// Solidity: function respondToTask((uint32,uint32,bytes,string,string,string) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) RespondToTask(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RespondToTask(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.SetPauserRegistry(&_ContractMfssiaTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.SetPauserRegistry(&_ContractMfssiaTaskManager.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.TransferOwnership(&_ContractMfssiaTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.TransferOwnership(&_ContractMfssiaTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Unpause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Unpause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// ContractMfssiaTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerInitializedIterator struct {
	Event *ContractMfssiaTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerInitialized represents a Initialized event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerInitializedIterator{contract: _ContractMfssiaTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerInitialized)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractMfssiaTaskManagerInitialized, error) {
	event := new(ContractMfssiaTaskManagerInitialized)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerNewTaskCreatedIterator struct {
	Event *ContractMfssiaTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IMfssiaTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x02ce4afaf544492edb45b478f1e8f27989b08d79fb36777a33affaf90366be94.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32,uint32,bytes,string,string,string) task)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractMfssiaTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerNewTaskCreatedIterator{contract: _ContractMfssiaTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x02ce4afaf544492edb45b478f1e8f27989b08d79fb36777a33affaf90366be94.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32,uint32,bytes,string,string,string) task)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerNewTaskCreated)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x02ce4afaf544492edb45b478f1e8f27989b08d79fb36777a33affaf90366be94.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32,uint32,bytes,string,string,string) task)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractMfssiaTaskManagerNewTaskCreated, error) {
	event := new(ContractMfssiaTaskManagerNewTaskCreated)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerOwnershipTransferredIterator struct {
	Event *ContractMfssiaTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractMfssiaTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerOwnershipTransferredIterator{contract: _ContractMfssiaTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerOwnershipTransferred)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractMfssiaTaskManagerOwnershipTransferred, error) {
	event := new(ContractMfssiaTaskManagerOwnershipTransferred)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPausedIterator struct {
	Event *ContractMfssiaTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerPaused represents a Paused event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractMfssiaTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerPausedIterator{contract: _ContractMfssiaTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerPaused)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParsePaused(log types.Log) (*ContractMfssiaTaskManagerPaused, error) {
	event := new(ContractMfssiaTaskManagerPaused)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPauserRegistrySetIterator struct {
	Event *ContractMfssiaTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerPauserRegistrySetIterator{contract: _ContractMfssiaTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerPauserRegistrySet)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractMfssiaTaskManagerPauserRegistrySet, error) {
	event := new(ContractMfssiaTaskManagerPauserRegistrySet)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractMfssiaTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractMfssiaTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractMfssiaTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractMfssiaTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskCompletedIterator struct {
	Event *ContractMfssiaTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractMfssiaTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskCompletedIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskCompleted)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractMfssiaTaskManagerTaskCompleted, error) {
	event := new(ContractMfssiaTaskManagerTaskCompleted)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskRespondedIterator struct {
	Event *ContractMfssiaTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskResponded represents a TaskResponded event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskResponded struct {
	TaskResponse         IMfssiaTaskManagerTaskResponse
	TaskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskRespondedIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskResponded)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractMfssiaTaskManagerTaskResponded, error) {
	event := new(ContractMfssiaTaskManagerTaskResponded)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerUnpausedIterator struct {
	Event *ContractMfssiaTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMfssiaTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMfssiaTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMfssiaTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerUnpaused represents a Unpaused event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractMfssiaTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerUnpausedIterator{contract: _ContractMfssiaTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerUnpaused)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractMfssiaTaskManagerUnpaused, error) {
	event := new(ContractMfssiaTaskManagerUnpaused)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

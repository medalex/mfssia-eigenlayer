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
	NumberToBeSquared         *big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IMfssiaTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

// IMfssiaTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// ContractMfssiaTaskManagerMetaData contains all meta data concerning the ContractMfssiaTaskManager contract.
var ContractMfssiaTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsPubkeyRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSPubkeyRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBLSOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200529638038062005296833981016040819052620000359162000169565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001b0565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001b0565b6001600160a01b031660c0525063ffffffff1660e05250620001d7565b6001600160a01b03811681146200016657600080fd5b50565b600080604083850312156200017d57600080fd5b82516200018a8162000150565b602084015190925063ffffffff81168114620001a557600080fd5b809150509250929050565b600060208284031215620001c357600080fd5b8151620001d08162000150565b9392505050565b60805160a05160c05160e05161504d620002496000396000818161025101528181610509015261159501526000818161030601528181611d2401526120970152600081816103e601528181612698015261283901526000818161043301528181611ae901526124ec015261504d6000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c80636b532e9e1161010f5780638b00ce7c116100a2578063f5c9899d11610071578063f5c9899d14610507578063f63c5bab1461052d578063f8c8765e14610535578063fabc1cbc1461054857600080fd5b80638b00ce7c146104b25780638da5cb5b146104c2578063cefdc1d4146104d3578063f2fde38b146104f457600080fd5b8063715018a6116100de578063715018a61461047657806372d18e8d1461047e5780637afa1eed1461048c578063886f11951461049f57600080fd5b80636b532e9e146104085780636b92787e1461041b5780636d14a9871461042e5780636efb46361461045557600080fd5b80633563b0d1116101875780635baec9a0116101565780635baec9a0146103a35780635c975abb146103b65780635decc3f5146103be57806368304835146103e157600080fd5b80633563b0d1146103285780634f739f7414610348578063595c6a67146103685780635ac86ab71461037057600080fd5b8063245a7bfc116101c3578063245a7bfc146102885780632cb223d5146102b35780632d89f6fc146102e15780633561deb11461030157600080fd5b806310d67a2f146101f5578063136439dd1461020a578063171f1d5b1461021d5780631ad431891461024c575b600080fd5b610208610203366004613c94565b61055b565b005b610208610218366004613cb8565b610617565b61023061022b366004613e36565b610756565b6040805192151583529015156020830152015b60405180910390f35b6102737f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610243565b609b5461029b906001600160a01b031681565b6040516001600160a01b039091168152602001610243565b6102d36102c1366004613ea4565b60996020526000908152604090205481565b604051908152602001610243565b6102d36102ef366004613ea4565b60986020526000908152604090205481565b61029b7f000000000000000000000000000000000000000000000000000000000000000081565b61033b610336366004613ec1565b6108e0565b6040516102439190614008565b61035b610356366004614063565b610c59565b6040516102439190614167565b6102086112dd565b61039361037e366004614222565b606654600160ff9092169190911b9081161490565b6040519015158152602001610243565b6102086103b1366004614508565b6113a4565b6066546102d3565b6103936103cc366004613ea4565b609a6020526000908152604090205460ff1681565b61029b7f000000000000000000000000000000000000000000000000000000000000000081565b61020861041636600461457c565b611823565b610208610429366004614602565b611ebf565b61029b7f000000000000000000000000000000000000000000000000000000000000000081565b61046861046336600461465d565b612060565b60405161024392919061471d565b610208612acc565b60975463ffffffff16610273565b609c5461029b906001600160a01b031681565b60655461029b906001600160a01b031681565b6097546102739063ffffffff1681565b6033546001600160a01b031661029b565b6104e66104e1366004614766565b612ae0565b6040516102439291906147a8565b610208610502366004613c94565b612c72565b7f0000000000000000000000000000000000000000000000000000000000000000610273565b610273606481565b6102086105433660046147c9565b612ce8565b610208610556366004613cb8565b612e39565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d29190614825565b6001600160a01b0316336001600160a01b03161461060b5760405162461bcd60e51b815260040161060290614842565b60405180910390fd5b61061481612f95565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561065f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610683919061488c565b61069f5760405162461bcd60e51b8152600401610602906148ae565b606654818116146107185760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610602565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061079e5761079e6148f6565b60200201518951600160200201518a602001516000600281106107c3576107c36148f6565b60200201518b602001516001600281106107df576107df6148f6565b602090810291909101518c518d83015160405161083c9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61085f919061490c565b90506108d2610878610871888461308c565b8690613123565b6108806131b7565b6108c86108b9856108b3604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061308c565b6108c28c613277565b90613123565b886201d4c0613307565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610922573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109469190614825565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610988573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ac9190614825565b9050600085516001600160401b038111156109c9576109c9613cd1565b6040519080825280602002602001820160405280156109fc57816020015b60608152602001906001900390816109e75790505b50905060005b8651811015610c4e576000878281518110610a1f57610a1f6148f6565b016020015160405163889ae3e560e01b815260f89190911c6004820181905263ffffffff8916602483015291506000906001600160a01b0386169063889ae3e590604401600060405180830381865afa158015610a80573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610aa8919081019061492e565b905080516001600160401b03811115610ac357610ac3613cd1565b604051908082528060200260200182016040528015610b0857816020015b6040805180820190915260008082526020820152815260200190600190039081610ae15790505b50848481518110610b1b57610b1b6148f6565b602002602001018190525060005b8151811015610c38576000828281518110610b4657610b466148f6565b6020908102919091018101516040805180820182528281529051631b32722560e01b81526004810183905260ff8816602482015263ffffffff8e166044820152919350918201906001600160a01b038b1690631b32722590606401602060405180830381865afa158015610bbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be291906149be565b6001600160601b0316815250868681518110610c0057610c006148f6565b60200260200101518381518110610c1957610c196148f6565b6020026020010181905250508080610c30906149fd565b915050610b29565b5050508080610c46906149fd565b915050610a02565b509695505050505050565b610c846040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce89190614825565b9050610d156040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516385020d4960e01b81526001600160a01b038a16906385020d4990610d45908b9089908990600401614a18565b600060405180830381865afa158015610d62573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d8a9190810190614a62565b815260405163e192e9ad60e01b81526001600160a01b0383169063e192e9ad90610dbc908b908b908b90600401614b19565b600060405180830381865afa158015610dd9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e019190810190614a62565b6040820152856001600160401b03811115610e1e57610e1e613cd1565b604051908082528060200260200182016040528015610e5157816020015b6060815260200190600190039081610e3c5790505b50606082015260005b60ff81168711156111ee576000856001600160401b03811115610e7f57610e7f613cd1565b604051908082528060200260200182016040528015610ea8578160200160208202803683370190505b5083606001518360ff1681518110610ec257610ec26148f6565b602002602001018190525060005b868110156110ee5760008c6001600160a01b0316633064620d8a8a85818110610efb57610efb6148f6565b905060200201358e88600001518681518110610f1957610f196148f6565b60200260200101516040518463ffffffff1660e01b8152600401610f569392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610f73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f979190614b42565b90508a8a8560ff16818110610fae57610fae6148f6565b6001600160c01b03841692013560f81c9190911c6001908116141590506110db57856001600160a01b031663480858668a8a85818110610ff057610ff06148f6565b905060200201358d8d8860ff1681811061100c5761100c6148f6565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611062573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110869190614b6b565b85606001518560ff168151811061109f5761109f6148f6565b602002602001015184815181106110b8576110b86148f6565b63ffffffff90921660209283029190910190910152826110d7816149fd565b9350505b50806110e6816149fd565b915050610ed0565b506000816001600160401b0381111561110957611109613cd1565b604051908082528060200260200182016040528015611132578160200160208202803683370190505b50905060005b828110156111b35784606001518460ff1681518110611159576111596148f6565b60200260200101518181518110611172576111726148f6565b602002602001015182828151811061118c5761118c6148f6565b63ffffffff90921660209283029190910190910152806111ab816149fd565b915050611138565b508084606001518460ff16815181106111ce576111ce6148f6565b6020026020010181905250505080806111e690614b88565b915050610e5a565b506000896001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa15801561122f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112539190614825565b60405163eda1076360e01b81529091506001600160a01b0382169063eda1076390611286908b908b908e90600401614ba8565b600060405180830381865afa1580156112a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112cb9190810190614a62565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611325573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611349919061488c565b6113655760405162461bcd60e51b8152600401610602906148ae565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b609b546001600160a01b031633146113fe5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610602565b60006114106040850160208601613ea4565b90503660006114226040870187614bd2565b909250905060006114396080880160608901613ea4565b90506098600061144c6020890189613ea4565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016114789190614c18565b60405160208183030381529060405280519060200120146115015760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610602565b600060998161151360208a018a613ea4565b63ffffffff1663ffffffff16815260200190815260200160002054146115905760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610602565b6115ba7f000000000000000000000000000000000000000000000000000000000000000085614cb9565b63ffffffff164363ffffffff16111561162b5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610602565b60008660405160200161163e9190614cff565b6040516020818303038152906040528051906020012090506000806116668387878a8c612060565b9150915060005b85811015611765578460ff168360200151828151811061168f5761168f6148f6565b60200260200101516116a19190614d0d565b6001600160601b03166064846000015183815181106116c2576116c26148f6565b60200260200101516001600160601b03166116dd9190614d3c565b1015611753576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610602565b8061175d816149fd565b91505061166d565b5060408051808201825263ffffffff43168152602080820184905291519091611792918c91849101614d5b565b60405160208183030381529060405280519060200120609960008c60000160208101906117bf9190613ea4565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a8260405161180e929190614d5b565b60405180910390a15050505050505050505050565b60006118326020850185613ea4565b63ffffffff81166000908152609960205260409020549091508535906118a45760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610602565b84846040516020016118b7929190614d87565b60408051601f19818403018152918152815160209283012063ffffffff851660009081526099909352912054146119565760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610602565b63ffffffff82166000908152609a602052604090205460ff16156119ee5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610602565b60646119fd6020860186613ea4565b611a079190614cb9565b63ffffffff164363ffffffff161115611a885760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610602565b6000611a948280614d3c565b9050602086013581146001811415611ae257604051339063ffffffff8616907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a350505050611eb9565b6000611b5a7f0000000000000000000000000000000000000000000000000000000000000000611b1560408c018c614bd2565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103369250505060408d0160208e01613ea4565b9050600086516001600160401b03811115611b7757611b77613cd1565b604051908082528060200260200182016040528015611ba0578160200160208202803683370190505b50905060005b8751811015611c0057611bd1888281518110611bc457611bc46148f6565b602002602001015161352b565b828281518110611be357611be36148f6565b602090810291909101015280611bf8816149fd565b915050611ba6565b506000611c1360408c0160208d01613ea4565b82604051602001611c25929190614dbd565b60405160208183030381529060405280519060200120905088602001358114611ccf5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610602565b600088516001600160401b03811115611cea57611cea613cd1565b604051908082528060200260200182016040528015611d13578160200160208202803683370190505b50905060005b8951811015611e67577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663187548c86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611da49190614825565b6001600160a01b031663e8bb9ae6858381518110611dc457611dc46148f6565b60200260200101516040518263ffffffff1660e01b8152600401611dea91815260200190565b602060405180830381865afa158015611e07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e2b9190614825565b828281518110611e3d57611e3d6148f6565b6001600160a01b039092166020928302919091019091015280611e5f816149fd565b915050611d19565b5063ffffffff88166000818152609a6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a350505050505050505b50505050565b609c546001600160a01b03163314611f235760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610602565b611f5a604051806080016040528060008152602001600063ffffffff16815260200160608152602001600063ffffffff1681525090565b84815263ffffffff438116602080840191909152908516606083015260408051601f850183900483028101830190915283815290849084908190840183828082843760009201919091525050505060408083019190915251611fc0908290602001614e31565b60408051601f1981840301815282825280516020918201206097805463ffffffff9081166000908152609890945293909220555416907f1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c4890612023908490614e31565b60405180910390a260975461203f9063ffffffff166001614cb9565b6097805463ffffffff191663ffffffff929092169190911790555050505050565b60408051808201909152606080825260208201526040805180820190915260008082526020820181905290815b8681101561227d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c1af6b248989848181106120d6576120d66148f6565b9050013560f81c60f81b60f81c888860a0015185815181106120fa576120fa6148f6565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612156573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061217a9190614e96565b6001600160401b03191661219d86604001518381518110611bc457611bc46148f6565b67ffffffffffffffff1916146122395760405162461bcd60e51b81526020600482015260616024820152600080516020614ff883398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610602565b61226985604001518281518110612252576122526148f6565b60200260200101518361312390919063ffffffff16565b915080612275816149fd565b91505061208d565b506040805180820190915260608082526020820152866001600160401b038111156122aa576122aa613cd1565b6040519080825280602002602001820160405280156122d3578160200160208202803683370190505b506020820152866001600160401b038111156122f1576122f1613cd1565b60405190808252806020026020018201604052801561231a578160200160208202803683370190505b5081526020850151516000906001600160401b0381111561233d5761233d613cd1565b604051908082528060200260200182016040528015612366578160200160208202803683370190505b50905060008660200151516001600160401b0381111561238857612388613cd1565b6040519080825280602002602001820160405280156123b1578160200160208202803683370190505b50905060006123f58b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061356e92505050565b905060005b8860200151518110156126605761242089602001518281518110611bc457611bc46148f6565b848281518110612432576124326148f6565b602090810291909101015280156124ea578361244f600183614ec1565b8151811061245f5761245f6148f6565b602002602001015160001c84828151811061247c5761247c6148f6565b602002602001015160001c116124ea576040805162461bcd60e51b8152602060048201526024810191909152600080516020614ff883398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610602565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633064620d85838151811061252b5761252b6148f6565b60200260200101518c8c60000151858151811061254a5761254a6148f6565b60200260200101516040518463ffffffff1660e01b81526004016125879392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156125a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c89190614b42565b6001600160c01b03168382815181106125e3576125e36148f6565b60200260200101818152505061264c6126456126198486858151811061260b5761260b6148f6565b6020026020010151166136d7565b61263f8c602001518581518110612632576126326148f6565b6020026020010151613708565b906137a3565b8790613123565b955080612658816149fd565b9150506123fa565b505060005b60ff81168a11156129a05760008b8b8360ff16818110612687576126876148f6565b9050013560f81c60f81b60f81c90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c56828c8c60c001518660ff16815181106126e0576126e06148f6565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561273c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061276091906149be565b85602001518360ff1681518110612779576127796148f6565b6001600160601b03909216602092830291909101820152850151805160ff84169081106127a8576127a86148f6565b602002602001015185600001518360ff16815181106127c9576127c96148f6565b60200260200101906001600160601b031690816001600160601b03168152505060005b8960200151518163ffffffff161015612996576000612832858363ffffffff168151811061281c5761281c6148f6565b60200260200101518460ff161c60019081161490565b15612983577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a43cde89848e898663ffffffff1681518110612880576128806148f6565b60200260200101518f60e001518960ff16815181106128a1576128a16148f6565b60200260200101518663ffffffff16815181106128c0576128c06148f6565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612924573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061294891906149be565b8751805160ff871690811061295f5761295f6148f6565b602002602001018181516129739190614ed8565b6001600160601b03169052506001015b508061298e81614f00565b9150506127ec565b5050600101612665565b50506000806129b98c868a606001518b60800151610756565b9150915081612a2a5760405162461bcd60e51b81526020600482015260436024820152600080516020614ff883398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610602565b80612a8b5760405162461bcd60e51b81526020600482015260396024820152600080516020614ff883398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610602565b505060008782604051602001612aa2929190614dbd565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612ad4613888565b612ade60006138e2565b565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612b1b57612b1b6148f6565b60209081029190910101526040516385020d4960e01b81526000906001600160a01b038816906385020d4990612b579088908690600401614f24565b600060405180830381865afa158015612b74573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612b9c9190810190614a62565b600081518110612bae57612bae6148f6565b6020908102919091010151604051633064620d60e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b03891690633064620d90606401602060405180830381865afa158015612c1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c3e9190614b42565b6001600160c01b031690506000612c5482613934565b905081612c628a838a6108e0565b9550955050505050935093915050565b612c7a613888565b6001600160a01b038116612cdf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610602565b610614816138e2565b600054610100900460ff1615808015612d085750600054600160ff909116105b80612d225750303b158015612d22575060005460ff166001145b612d855760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610602565b6000805460ff191660011790558015612da8576000805461ff0019166101001790555b612db3856000613991565b612dbc846138e2565b609b80546001600160a01b038086166001600160a01b031992831617909255609c8054928516929091169190911790558015612e32576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eb09190614825565b6001600160a01b0316336001600160a01b031614612ee05760405162461bcd60e51b815260040161060290614842565b606654198119606654191614612f5e5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610602565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161074b565b6001600160a01b0381166130235760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610602565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526130a8613ba5565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156130db576130dd565bfe5b508061311b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610602565b505092915050565b604080518082019091526000808252602082015261313f613bc3565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156130db57508061311b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610602565b6131bf613be1565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806132a7600080516020614fd88339815191528661490c565b90505b6132b381613a7b565b9093509150600080516020614fd88339815191528283098314156132ed576040805180820190915290815260208101919091529392505050565b600080516020614fd88339815191526001820890506132aa565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613339613c06565b60005b60028110156134fe576000613352826006614d3c565b9050848260028110613366576133666148f6565b60200201515183613378836000614f78565b600c8110613388576133886148f6565b602002015284826002811061339f5761339f6148f6565b602002015160200151838260016133b69190614f78565b600c81106133c6576133c66148f6565b60200201528382600281106133dd576133dd6148f6565b60200201515151836133f0836002614f78565b600c8110613400576134006148f6565b6020020152838260028110613417576134176148f6565b6020020151516001602002015183613430836003614f78565b600c8110613440576134406148f6565b6020020152838260028110613457576134576148f6565b602002015160200151600060028110613472576134726148f6565b602002015183613483836004614f78565b600c8110613493576134936148f6565b60200201528382600281106134aa576134aa6148f6565b6020020151602001516001600281106134c5576134c56148f6565b6020020151836134d6836005614f78565b600c81106134e6576134e66148f6565b602002015250806134f6816149fd565b91505061333c565b50613507613c25565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600081600001518260200151604051602001613551929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b6000610100825111156135e25760405162461bcd60e51b815260206004820152603660248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a206044820152756279746573417272617920697320746f6f206c6f6e6760501b6064820152608401610602565b81516135f057506000919050565b60008083600081518110613606576136066148f6565b0160200151600160f89190911c81901b92505b84518110156136ce57848181518110613634576136346148f6565b0160200151600160f89190911c1b9150828216156136ba5760405162461bcd60e51b815260206004820152603a60248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a2060448201527f72657065617420656e74727920696e20627974657341727261790000000000006064820152608401610602565b918117916136c7816149fd565b9050613619565b50909392505050565b6000805b8215613702576136ec600184614ec1565b90921691806136fa81614f90565b9150506136db565b92915050565b6040805180820190915260008082526020820152815115801561372d57506020820151155b1561374b575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020614fd8833981519152846020015161377e919061490c565b61379690600080516020614fd8833981519152614ec1565b905292915050565b919050565b60408051808201909152600080825260208201526102008261ffff16106137ff5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610602565b8161ffff1660011415613813575081613702565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16111561387d57600161ffff871660ff83161c811614156138605761385d8484613123565b93505b61386a8384613123565b92506201fffe600192831b16910161382f565b509195945050505050565b6033546001600160a01b03163314612ade5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610602565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b61010081101561398a576001811b91508382161561397a57828160f81b604051602001613968929190614fa8565b60405160208183030381529060405292505b613983816149fd565b905061393a565b5050919050565b6065546001600160a01b03161580156139b257506001600160a01b03821615155b613a345760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610602565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613a7782612f95565b5050565b60008080600080516020614fd88339815191526003600080516020614fd883398151915286600080516020614fd8833981519152888909090890506000613af1827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020614fd8833981519152613afd565b91959194509092505050565b600080613b08613c25565b613b10613c43565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156130db575082613b9a5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610602565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613bf4613c61565b8152602001613c01613c61565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461061457600080fd5b600060208284031215613ca657600080fd5b8135613cb181613c7f565b9392505050565b600060208284031215613cca57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613d0957613d09613cd1565b60405290565b60405161010081016001600160401b0381118282101715613d0957613d09613cd1565b604051601f8201601f191681016001600160401b0381118282101715613d5a57613d5a613cd1565b604052919050565b600060408284031215613d7457600080fd5b613d7c613ce7565b9050813581526020820135602082015292915050565b600082601f830112613da357600080fd5b604051604081018181106001600160401b0382111715613dc557613dc5613cd1565b8060405250806040840185811115613ddc57600080fd5b845b8181101561387d578035835260209283019201613dde565b600060808284031215613e0857600080fd5b613e10613ce7565b9050613e1c8383613d92565b8152613e2b8360408401613d92565b602082015292915050565b6000806000806101208587031215613e4d57600080fd5b84359350613e5e8660208701613d62565b9250613e6d8660608701613df6565b9150613e7c8660e08701613d62565b905092959194509250565b63ffffffff8116811461061457600080fd5b803561379e81613e87565b600060208284031215613eb657600080fd5b8135613cb181613e87565b600080600060608486031215613ed657600080fd5b8335613ee181613c7f565b92506020848101356001600160401b0380821115613efe57600080fd5b818701915087601f830112613f1257600080fd5b813581811115613f2457613f24613cd1565b613f36601f8201601f19168501613d32565b91508082528884828501011115613f4c57600080fd5b8084840185840137600084828401015250809450505050613f6f60408501613e99565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015613ffa578385038a52825180518087529087019087870190845b81811015613fe5578351805184528a01516001600160601b03168a84015292890192604090920191600101613fb5565b50509a87019a95505091850191600101613f97565b509298975050505050505050565b602081526000613cb16020830184613f78565b60008083601f84011261402d57600080fd5b5081356001600160401b0381111561404457600080fd5b60208301915083602082850101111561405c57600080fd5b9250929050565b6000806000806000806080878903121561407c57600080fd5b863561408781613c7f565b9550602087013561409781613e87565b945060408701356001600160401b03808211156140b357600080fd5b6140bf8a838b0161401b565b909650945060608901359150808211156140d857600080fd5b818901915089601f8301126140ec57600080fd5b8135818111156140fb57600080fd5b8a60208260051b850101111561411057600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561415c57815163ffffffff168752958201959082019060010161413a565b509495945050505050565b60006020808352835160808285015261418360a0850182614126565b905081850151601f19808684030160408701526141a08383614126565b925060408701519150808684030160608701526141bd8383614126565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156142145784878303018452614202828751614126565b958801959388019391506001016141e8565b509998505050505050505050565b60006020828403121561423457600080fd5b813560ff81168114613cb157600080fd5b60006080828403121561425757600080fd5b50919050565b60006040828403121561425757600080fd5b60006001600160401b0382111561428857614288613cd1565b5060051b60200190565b600082601f8301126142a357600080fd5b813560206142b86142b38361426f565b613d32565b82815260059290921b840181019181810190868411156142d757600080fd5b8286015b84811015610c4e5780356142ee81613e87565b83529183019183016142db565b600082601f83011261430c57600080fd5b8135602061431c6142b38361426f565b82815260069290921b8401810191818101908684111561433b57600080fd5b8286015b84811015610c4e576143518882613d62565b83529183019160400161433f565b600082601f83011261437057600080fd5b813560206143806142b38361426f565b82815260059290921b8401810191818101908684111561439f57600080fd5b8286015b84811015610c4e5780356001600160401b038111156143c25760008081fd5b6143d08986838b0101614292565b8452509183019183016143a3565b600061018082840312156143f157600080fd5b6143f9613d0f565b905081356001600160401b038082111561441257600080fd5b61441e85838601614292565b8352602084013591508082111561443457600080fd5b614440858386016142fb565b6020840152604084013591508082111561445957600080fd5b614465858386016142fb565b60408401526144778560608601613df6565b60608401526144898560e08601613d62565b60808401526101208401359150808211156144a357600080fd5b6144af85838601614292565b60a08401526101408401359150808211156144c957600080fd5b6144d585838601614292565b60c08401526101608401359150808211156144ef57600080fd5b506144fc8482850161435f565b60e08301525092915050565b60008060006080848603121561451d57600080fd5b83356001600160401b038082111561453457600080fd5b61454087838801614245565b945061454f876020880161425d565b9350606086013591508082111561456557600080fd5b50614572868287016143de565b9150509250925092565b60008060008060c0858703121561459257600080fd5b84356001600160401b03808211156145a957600080fd5b6145b588838901614245565b95506145c4886020890161425d565b94506145d3886060890161425d565b935060a08701359150808211156145e957600080fd5b506145f6878288016142fb565b91505092959194509250565b6000806000806060858703121561461857600080fd5b84359350602085013561462a81613e87565b925060408501356001600160401b0381111561464557600080fd5b6146518782880161401b565b95989497509550505050565b60008060008060006080868803121561467557600080fd5b8535945060208601356001600160401b038082111561469357600080fd5b61469f89838a0161401b565b9096509450604088013591506146b482613e87565b909250606087013590808211156146ca57600080fd5b506146d7888289016143de565b9150509295509295909350565b600081518084526020808501945080840160005b8381101561415c5781516001600160601b0316875295820195908201906001016146f8565b604081526000835160408084015261473860808401826146e4565b90506020850151603f1984830301606085015261475582826146e4565b925050508260208301529392505050565b60008060006060848603121561477b57600080fd5b833561478681613c7f565b925060208401359150604084013561479d81613e87565b809150509250925092565b8281526040602082015260006147c16040830184613f78565b949350505050565b600080600080608085870312156147df57600080fd5b84356147ea81613c7f565b935060208501356147fa81613c7f565b9250604085013561480a81613c7f565b9150606085013561481a81613c7f565b939692955090935050565b60006020828403121561483757600080fd5b8151613cb181613c7f565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561489e57600080fd5b81518015158114613cb157600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261492957634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561494157600080fd5b82516001600160401b0381111561495757600080fd5b8301601f8101851361496857600080fd5b80516149766142b38261426f565b81815260059190911b8201830190838101908783111561499557600080fd5b928401925b828410156149b35783518252928401929084019061499a565b979650505050505050565b6000602082840312156149d057600080fd5b81516001600160601b0381168114613cb157600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415614a1157614a116149e7565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115614a4557600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215614a7557600080fd5b82516001600160401b03811115614a8b57600080fd5b8301601f81018513614a9c57600080fd5b8051614aaa6142b38261426f565b81815260059190911b82018301908381019087831115614ac957600080fd5b928401925b828410156149b3578351614ae181613e87565b82529284019290840190614ace565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000614b39604083018486614af0565b95945050505050565b600060208284031215614b5457600080fd5b81516001600160c01b0381168114613cb157600080fd5b600060208284031215614b7d57600080fd5b8151613cb181613e87565b600060ff821660ff811415614b9f57614b9f6149e7565b60010192915050565b604081526000614bbc604083018587614af0565b905063ffffffff83166020830152949350505050565b6000808335601e19843603018112614be957600080fd5b8301803591506001600160401b03821115614c0357600080fd5b60200191503681900382131561405c57600080fd5b602081528135602082015260006020830135614c3381613e87565b63ffffffff81166040840152506040830135601e19843603018112614c5757600080fd5b830180356001600160401b03811115614c6f57600080fd5b803603851315614c7e57600080fd5b60806060850152614c9660a085018260208501614af0565b915050614ca560608501613e99565b63ffffffff81166080850152509392505050565b600063ffffffff808316818516808303821115614cd857614cd86149e7565b01949350505050565b8035614cec81613e87565b63ffffffff168252602090810135910152565b604081016137028284614ce1565b60006001600160601b0380831681851681830481118215151615614d3357614d336149e7565b02949350505050565b6000816000190483118215151615614d5657614d566149e7565b500290565b60808101614d698285614ce1565b63ffffffff8351166040830152602083015160608301529392505050565b60808101614d958285614ce1565b8235614da081613e87565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015614df857815185529382019390820190600101614ddc565b5092979650505050505050565b60005b83811015614e20578181015183820152602001614e08565b83811115611eb95750506000910152565b60208152815160208201526000602083015163ffffffff8082166040850152604085015191506080606085015281518060a0860152614e778160c0870160208601614e05565b60609590950151166080840152505060c0601f909201601f1916010190565b600060208284031215614ea857600080fd5b815167ffffffffffffffff1981168114613cb157600080fd5b600082821015614ed357614ed36149e7565b500390565b60006001600160601b0383811690831681811015614ef857614ef86149e7565b039392505050565b600063ffffffff80831681811415614f1a57614f1a6149e7565b6001019392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015614f6b57845183529383019391830191600101614f4f565b5090979650505050505050565b60008219821115614f8b57614f8b6149e7565b500190565b600061ffff80831681811415614f1a57614f1a6149e7565b60008351614fba818460208801614e05565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220a8c5164f9a02610c587b9a735362a05cb5a5ead51d88db8aea53c32a38d764a164736f6c634300080c0033",
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

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "createNewTask", numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) CreateNewTask(numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.CreateNewTask(&_ContractMfssiaTaskManager.TransactOpts, numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) CreateNewTask(numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.CreateNewTask(&_ContractMfssiaTaskManager.TransactOpts, numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
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

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RaiseAndResolveChallenge(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RaiseAndResolveChallenge(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
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

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RespondToTask(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RespondToTask(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
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

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
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

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
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

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskRespondedIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
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

// ParseTaskResponded is a log parse operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
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

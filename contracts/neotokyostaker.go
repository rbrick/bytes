// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// NeoTokyoStakerLPPosition is an auto generated low-level Go binding around an user-defined struct.
type NeoTokyoStakerLPPosition struct {
	Amount          *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	Multiplier      *big.Int
}

// NeoTokyoStakerPoolConfigurationInput is an auto generated low-level Go binding around an user-defined struct.
type NeoTokyoStakerPoolConfigurationInput struct {
	AssetType     uint8
	DaoTax        *big.Int
	RewardWindows []NeoTokyoStakerRewardWindow
}

// NeoTokyoStakerRewardWindow is an auto generated low-level Go binding around an user-defined struct.
type NeoTokyoStakerRewardWindow struct {
	StartTime *big.Int
	Reward    *big.Int
}

// NeoTokyoStakerStakedS1CitizenOutput is an auto generated low-level Go binding around an user-defined struct.
type NeoTokyoStakerStakedS1CitizenOutput struct {
	CitizenId       *big.Int
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	StakedVaultId   *big.Int
	HasVault        bool
}

// NeoTokyoStakerStakedS2CitizenOutput is an auto generated low-level Go binding around an user-defined struct.
type NeoTokyoStakerStakedS2CitizenOutput struct {
	CitizenId       *big.Int
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
}

// NeoTokyoStakerStakerPosition is an auto generated low-level Go binding around an user-defined struct.
type NeoTokyoStakerStakerPosition struct {
	StakedS1Citizens []NeoTokyoStakerStakedS1CitizenOutput
	StakedS2Citizens []NeoTokyoStakerStakedS2CitizenOutput
	StakedLPPosition NeoTokyoStakerLPPosition
}

// NeoTokyoStakerMetaData contains all meta data concerning the NeoTokyoStaker contract.
var NeoTokyoStakerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bytes\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_s1Citizen\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_s2Citizen\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sbt\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_noVaultCap\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"managedRight\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"managerRight\",\"type\":\"bytes32\"}],\"name\":\"ManagementUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"circumstance\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"PermitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timelockOption\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOrTokenId\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOrTokenId\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BYTES\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIGURE_CAPS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIGURE_CREDITS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIGURE_LP\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIGURE_POOLS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIGURE_TIMELOCKS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IDENTITY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NO_VAULT_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NT_STAKED_CITIZEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"S1_CITIZEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"S2_CITIZEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNIVERSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_RIGHT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultedCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unvaultedCap\",\"type\":\"uint256\"}],\"name\":\"configureCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_identityCreditYields\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_points\",\"type\":\"uint256[]\"}],\"name\":\"configureIdentityCreditPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_citizenRewardRates\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_vaultRewardRates\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_identityCreditYields\",\"type\":\"string[]\"}],\"name\":\"configureIdentityCreditYields\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"configureLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"daoTax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"startTime\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"reward\",\"type\":\"uint128\"}],\"internalType\":\"structNeoTokyoStaker.RewardWindow[]\",\"name\":\"rewardWindows\",\"type\":\"tuple[]\"}],\"internalType\":\"structNeoTokyoStaker.PoolConfigurationInput[]\",\"name\":\"_inputs\",\"type\":\"tuple[]\"}],\"name\":\"configurePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"_assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"_timelockIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_encodedSettings\",\"type\":\"uint256[]\"}],\"name\":\"configureTimelockOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_vaultCreditMultipliers\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_multipliers\",\"type\":\"uint256[]\"}],\"name\":\"configureVaultCreditMultipliers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"getConfiguredVaultMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_citizenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"getCreditYield\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"_assetType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"getPendingPoolReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"_assetType\",\"type\":\"uint8\"}],\"name\":\"getStakerPosition\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerPositions\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedVaultId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVault\",\"type\":\"bool\"}],\"internalType\":\"structNeoTokyoStaker.StakedS1CitizenOutput[]\",\"name\":\"stakedS1Citizens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"}],\"internalType\":\"structNeoTokyoStaker.StakedS2CitizenOutput[]\",\"name\":\"stakedS2Citizens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"internalType\":\"structNeoTokyoStaker.LPPosition\",\"name\":\"stakedLPPosition\",\"type\":\"tuple\"}],\"internalType\":\"structNeoTokyoStaker.StakerPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"_assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"}],\"name\":\"getTotalEmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_circumstance\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hasRight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_circumstance\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hasRightUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"identityCreditPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"identityCreditYield\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"lastRewardTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"managerRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"permissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"rewardAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_managedRight\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_managerRight\",\"type\":\"bytes32\"}],\"name\":\"setManagerRight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_circumstance\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expirationTime\",\"type\":\"uint256\"}],\"name\":\"setPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"_assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_timelockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedS1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedVaultId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVault\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedS2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerLPPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timelockOptions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"vaultCreditMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNeoTokyoStaker.AssetType\",\"name\":\"_assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NeoTokyoStakerABI is the input ABI used to generate the binding from.
// Deprecated: Use NeoTokyoStakerMetaData.ABI instead.
var NeoTokyoStakerABI = NeoTokyoStakerMetaData.ABI

// NeoTokyoStaker is an auto generated Go binding around an Ethereum contract.
type NeoTokyoStaker struct {
	NeoTokyoStakerCaller     // Read-only binding to the contract
	NeoTokyoStakerTransactor // Write-only binding to the contract
	NeoTokyoStakerFilterer   // Log filterer for contract events
}

// NeoTokyoStakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NeoTokyoStakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoTokyoStakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NeoTokyoStakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoTokyoStakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NeoTokyoStakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoTokyoStakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NeoTokyoStakerSession struct {
	Contract     *NeoTokyoStaker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeoTokyoStakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NeoTokyoStakerCallerSession struct {
	Contract *NeoTokyoStakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NeoTokyoStakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NeoTokyoStakerTransactorSession struct {
	Contract     *NeoTokyoStakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NeoTokyoStakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NeoTokyoStakerRaw struct {
	Contract *NeoTokyoStaker // Generic contract binding to access the raw methods on
}

// NeoTokyoStakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NeoTokyoStakerCallerRaw struct {
	Contract *NeoTokyoStakerCaller // Generic read-only contract binding to access the raw methods on
}

// NeoTokyoStakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NeoTokyoStakerTransactorRaw struct {
	Contract *NeoTokyoStakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeoTokyoStaker creates a new instance of NeoTokyoStaker, bound to a specific deployed contract.
func NewNeoTokyoStaker(address common.Address, backend bind.ContractBackend) (*NeoTokyoStaker, error) {
	contract, err := bindNeoTokyoStaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStaker{NeoTokyoStakerCaller: NeoTokyoStakerCaller{contract: contract}, NeoTokyoStakerTransactor: NeoTokyoStakerTransactor{contract: contract}, NeoTokyoStakerFilterer: NeoTokyoStakerFilterer{contract: contract}}, nil
}

// NewNeoTokyoStakerCaller creates a new read-only instance of NeoTokyoStaker, bound to a specific deployed contract.
func NewNeoTokyoStakerCaller(address common.Address, caller bind.ContractCaller) (*NeoTokyoStakerCaller, error) {
	contract, err := bindNeoTokyoStaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerCaller{contract: contract}, nil
}

// NewNeoTokyoStakerTransactor creates a new write-only instance of NeoTokyoStaker, bound to a specific deployed contract.
func NewNeoTokyoStakerTransactor(address common.Address, transactor bind.ContractTransactor) (*NeoTokyoStakerTransactor, error) {
	contract, err := bindNeoTokyoStaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerTransactor{contract: contract}, nil
}

// NewNeoTokyoStakerFilterer creates a new log filterer instance of NeoTokyoStaker, bound to a specific deployed contract.
func NewNeoTokyoStakerFilterer(address common.Address, filterer bind.ContractFilterer) (*NeoTokyoStakerFilterer, error) {
	contract, err := bindNeoTokyoStaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerFilterer{contract: contract}, nil
}

// bindNeoTokyoStaker binds a generic wrapper to an already deployed contract.
func bindNeoTokyoStaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NeoTokyoStakerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeoTokyoStaker *NeoTokyoStakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NeoTokyoStaker.Contract.NeoTokyoStakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeoTokyoStaker *NeoTokyoStakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.NeoTokyoStakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeoTokyoStaker *NeoTokyoStakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.NeoTokyoStakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeoTokyoStaker *NeoTokyoStakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NeoTokyoStaker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeoTokyoStaker *NeoTokyoStakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeoTokyoStaker *NeoTokyoStakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.contract.Transact(opts, method, params...)
}

// BYTES is a free data retrieval call binding the contract method 0xe88343e6.
//
// Solidity: function BYTES() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) BYTES(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "BYTES")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BYTES is a free data retrieval call binding the contract method 0xe88343e6.
//
// Solidity: function BYTES() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) BYTES() (common.Address, error) {
	return _NeoTokyoStaker.Contract.BYTES(&_NeoTokyoStaker.CallOpts)
}

// BYTES is a free data retrieval call binding the contract method 0xe88343e6.
//
// Solidity: function BYTES() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) BYTES() (common.Address, error) {
	return _NeoTokyoStaker.Contract.BYTES(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURECAPS is a free data retrieval call binding the contract method 0x36f0af78.
//
// Solidity: function CONFIGURE_CAPS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) CONFIGURECAPS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "CONFIGURE_CAPS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGURECAPS is a free data retrieval call binding the contract method 0x36f0af78.
//
// Solidity: function CONFIGURE_CAPS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) CONFIGURECAPS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURECAPS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURECAPS is a free data retrieval call binding the contract method 0x36f0af78.
//
// Solidity: function CONFIGURE_CAPS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) CONFIGURECAPS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURECAPS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURECREDITS is a free data retrieval call binding the contract method 0x50e63475.
//
// Solidity: function CONFIGURE_CREDITS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) CONFIGURECREDITS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "CONFIGURE_CREDITS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGURECREDITS is a free data retrieval call binding the contract method 0x50e63475.
//
// Solidity: function CONFIGURE_CREDITS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) CONFIGURECREDITS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURECREDITS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURECREDITS is a free data retrieval call binding the contract method 0x50e63475.
//
// Solidity: function CONFIGURE_CREDITS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) CONFIGURECREDITS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURECREDITS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURELP is a free data retrieval call binding the contract method 0x6b7e290b.
//
// Solidity: function CONFIGURE_LP() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) CONFIGURELP(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "CONFIGURE_LP")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGURELP is a free data retrieval call binding the contract method 0x6b7e290b.
//
// Solidity: function CONFIGURE_LP() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) CONFIGURELP() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURELP(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURELP is a free data retrieval call binding the contract method 0x6b7e290b.
//
// Solidity: function CONFIGURE_LP() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) CONFIGURELP() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURELP(&_NeoTokyoStaker.CallOpts)
}

// CONFIGUREPOOLS is a free data retrieval call binding the contract method 0x8551c918.
//
// Solidity: function CONFIGURE_POOLS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) CONFIGUREPOOLS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "CONFIGURE_POOLS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGUREPOOLS is a free data retrieval call binding the contract method 0x8551c918.
//
// Solidity: function CONFIGURE_POOLS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) CONFIGUREPOOLS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGUREPOOLS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGUREPOOLS is a free data retrieval call binding the contract method 0x8551c918.
//
// Solidity: function CONFIGURE_POOLS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) CONFIGUREPOOLS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGUREPOOLS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURETIMELOCKS is a free data retrieval call binding the contract method 0xe52b622a.
//
// Solidity: function CONFIGURE_TIMELOCKS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) CONFIGURETIMELOCKS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "CONFIGURE_TIMELOCKS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGURETIMELOCKS is a free data retrieval call binding the contract method 0xe52b622a.
//
// Solidity: function CONFIGURE_TIMELOCKS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) CONFIGURETIMELOCKS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURETIMELOCKS(&_NeoTokyoStaker.CallOpts)
}

// CONFIGURETIMELOCKS is a free data retrieval call binding the contract method 0xe52b622a.
//
// Solidity: function CONFIGURE_TIMELOCKS() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) CONFIGURETIMELOCKS() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.CONFIGURETIMELOCKS(&_NeoTokyoStaker.CallOpts)
}

// IDENTITY is a free data retrieval call binding the contract method 0x150c4aea.
//
// Solidity: function IDENTITY() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) IDENTITY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "IDENTITY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IDENTITY is a free data retrieval call binding the contract method 0x150c4aea.
//
// Solidity: function IDENTITY() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) IDENTITY() (common.Address, error) {
	return _NeoTokyoStaker.Contract.IDENTITY(&_NeoTokyoStaker.CallOpts)
}

// IDENTITY is a free data retrieval call binding the contract method 0x150c4aea.
//
// Solidity: function IDENTITY() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) IDENTITY() (common.Address, error) {
	return _NeoTokyoStaker.Contract.IDENTITY(&_NeoTokyoStaker.CallOpts)
}

// LP is a free data retrieval call binding the contract method 0xb6fccf8a.
//
// Solidity: function LP() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) LP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "LP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LP is a free data retrieval call binding the contract method 0xb6fccf8a.
//
// Solidity: function LP() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) LP() (common.Address, error) {
	return _NeoTokyoStaker.Contract.LP(&_NeoTokyoStaker.CallOpts)
}

// LP is a free data retrieval call binding the contract method 0xb6fccf8a.
//
// Solidity: function LP() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) LP() (common.Address, error) {
	return _NeoTokyoStaker.Contract.LP(&_NeoTokyoStaker.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) MANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) MANAGER() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.MANAGER(&_NeoTokyoStaker.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) MANAGER() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.MANAGER(&_NeoTokyoStaker.CallOpts)
}

// NOVAULTCAP is a free data retrieval call binding the contract method 0x8d6389a1.
//
// Solidity: function NO_VAULT_CAP() view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) NOVAULTCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "NO_VAULT_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NOVAULTCAP is a free data retrieval call binding the contract method 0x8d6389a1.
//
// Solidity: function NO_VAULT_CAP() view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) NOVAULTCAP() (*big.Int, error) {
	return _NeoTokyoStaker.Contract.NOVAULTCAP(&_NeoTokyoStaker.CallOpts)
}

// NOVAULTCAP is a free data retrieval call binding the contract method 0x8d6389a1.
//
// Solidity: function NO_VAULT_CAP() view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) NOVAULTCAP() (*big.Int, error) {
	return _NeoTokyoStaker.Contract.NOVAULTCAP(&_NeoTokyoStaker.CallOpts)
}

// NTSTAKEDCITIZEN is a free data retrieval call binding the contract method 0x42d0f0b6.
//
// Solidity: function NT_STAKED_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) NTSTAKEDCITIZEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "NT_STAKED_CITIZEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NTSTAKEDCITIZEN is a free data retrieval call binding the contract method 0x42d0f0b6.
//
// Solidity: function NT_STAKED_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) NTSTAKEDCITIZEN() (common.Address, error) {
	return _NeoTokyoStaker.Contract.NTSTAKEDCITIZEN(&_NeoTokyoStaker.CallOpts)
}

// NTSTAKEDCITIZEN is a free data retrieval call binding the contract method 0x42d0f0b6.
//
// Solidity: function NT_STAKED_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) NTSTAKEDCITIZEN() (common.Address, error) {
	return _NeoTokyoStaker.Contract.NTSTAKEDCITIZEN(&_NeoTokyoStaker.CallOpts)
}

// S1CITIZEN is a free data retrieval call binding the contract method 0x31100365.
//
// Solidity: function S1_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) S1CITIZEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "S1_CITIZEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// S1CITIZEN is a free data retrieval call binding the contract method 0x31100365.
//
// Solidity: function S1_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) S1CITIZEN() (common.Address, error) {
	return _NeoTokyoStaker.Contract.S1CITIZEN(&_NeoTokyoStaker.CallOpts)
}

// S1CITIZEN is a free data retrieval call binding the contract method 0x31100365.
//
// Solidity: function S1_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) S1CITIZEN() (common.Address, error) {
	return _NeoTokyoStaker.Contract.S1CITIZEN(&_NeoTokyoStaker.CallOpts)
}

// S2CITIZEN is a free data retrieval call binding the contract method 0x8904486e.
//
// Solidity: function S2_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) S2CITIZEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "S2_CITIZEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// S2CITIZEN is a free data retrieval call binding the contract method 0x8904486e.
//
// Solidity: function S2_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) S2CITIZEN() (common.Address, error) {
	return _NeoTokyoStaker.Contract.S2CITIZEN(&_NeoTokyoStaker.CallOpts)
}

// S2CITIZEN is a free data retrieval call binding the contract method 0x8904486e.
//
// Solidity: function S2_CITIZEN() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) S2CITIZEN() (common.Address, error) {
	return _NeoTokyoStaker.Contract.S2CITIZEN(&_NeoTokyoStaker.CallOpts)
}

// UNIVERSAL is a free data retrieval call binding the contract method 0x17f5ebb4.
//
// Solidity: function UNIVERSAL() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) UNIVERSAL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "UNIVERSAL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNIVERSAL is a free data retrieval call binding the contract method 0x17f5ebb4.
//
// Solidity: function UNIVERSAL() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) UNIVERSAL() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.UNIVERSAL(&_NeoTokyoStaker.CallOpts)
}

// UNIVERSAL is a free data retrieval call binding the contract method 0x17f5ebb4.
//
// Solidity: function UNIVERSAL() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) UNIVERSAL() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.UNIVERSAL(&_NeoTokyoStaker.CallOpts)
}

// VAULT is a free data retrieval call binding the contract method 0x411557d1.
//
// Solidity: function VAULT() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) VAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VAULT is a free data retrieval call binding the contract method 0x411557d1.
//
// Solidity: function VAULT() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) VAULT() (common.Address, error) {
	return _NeoTokyoStaker.Contract.VAULT(&_NeoTokyoStaker.CallOpts)
}

// VAULT is a free data retrieval call binding the contract method 0x411557d1.
//
// Solidity: function VAULT() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) VAULT() (common.Address, error) {
	return _NeoTokyoStaker.Contract.VAULT(&_NeoTokyoStaker.CallOpts)
}

// VAULTCAP is a free data retrieval call binding the contract method 0x76d8b246.
//
// Solidity: function VAULT_CAP() view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) VAULTCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "VAULT_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VAULTCAP is a free data retrieval call binding the contract method 0x76d8b246.
//
// Solidity: function VAULT_CAP() view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) VAULTCAP() (*big.Int, error) {
	return _NeoTokyoStaker.Contract.VAULTCAP(&_NeoTokyoStaker.CallOpts)
}

// VAULTCAP is a free data retrieval call binding the contract method 0x76d8b246.
//
// Solidity: function VAULT_CAP() view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) VAULTCAP() (*big.Int, error) {
	return _NeoTokyoStaker.Contract.VAULTCAP(&_NeoTokyoStaker.CallOpts)
}

// ZERORIGHT is a free data retrieval call binding the contract method 0xa625776e.
//
// Solidity: function ZERO_RIGHT() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) ZERORIGHT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "ZERO_RIGHT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZERORIGHT is a free data retrieval call binding the contract method 0xa625776e.
//
// Solidity: function ZERO_RIGHT() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) ZERORIGHT() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.ZERORIGHT(&_NeoTokyoStaker.CallOpts)
}

// ZERORIGHT is a free data retrieval call binding the contract method 0xa625776e.
//
// Solidity: function ZERO_RIGHT() view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) ZERORIGHT() ([32]byte, error) {
	return _NeoTokyoStaker.Contract.ZERORIGHT(&_NeoTokyoStaker.CallOpts)
}

// GetConfiguredVaultMultiplier is a free data retrieval call binding the contract method 0x00a90e85.
//
// Solidity: function getConfiguredVaultMultiplier(uint256 _vaultId) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) GetConfiguredVaultMultiplier(opts *bind.CallOpts, _vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "getConfiguredVaultMultiplier", _vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfiguredVaultMultiplier is a free data retrieval call binding the contract method 0x00a90e85.
//
// Solidity: function getConfiguredVaultMultiplier(uint256 _vaultId) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) GetConfiguredVaultMultiplier(_vaultId *big.Int) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.GetConfiguredVaultMultiplier(&_NeoTokyoStaker.CallOpts, _vaultId)
}

// GetConfiguredVaultMultiplier is a free data retrieval call binding the contract method 0x00a90e85.
//
// Solidity: function getConfiguredVaultMultiplier(uint256 _vaultId) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) GetConfiguredVaultMultiplier(_vaultId *big.Int) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.GetConfiguredVaultMultiplier(&_NeoTokyoStaker.CallOpts, _vaultId)
}

// GetCreditYield is a free data retrieval call binding the contract method 0xdd6fea0b.
//
// Solidity: function getCreditYield(uint256 _citizenId, uint256 _vaultId) view returns(string)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) GetCreditYield(opts *bind.CallOpts, _citizenId *big.Int, _vaultId *big.Int) (string, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "getCreditYield", _citizenId, _vaultId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCreditYield is a free data retrieval call binding the contract method 0xdd6fea0b.
//
// Solidity: function getCreditYield(uint256 _citizenId, uint256 _vaultId) view returns(string)
func (_NeoTokyoStaker *NeoTokyoStakerSession) GetCreditYield(_citizenId *big.Int, _vaultId *big.Int) (string, error) {
	return _NeoTokyoStaker.Contract.GetCreditYield(&_NeoTokyoStaker.CallOpts, _citizenId, _vaultId)
}

// GetCreditYield is a free data retrieval call binding the contract method 0xdd6fea0b.
//
// Solidity: function getCreditYield(uint256 _citizenId, uint256 _vaultId) view returns(string)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) GetCreditYield(_citizenId *big.Int, _vaultId *big.Int) (string, error) {
	return _NeoTokyoStaker.Contract.GetCreditYield(&_NeoTokyoStaker.CallOpts, _citizenId, _vaultId)
}

// GetPendingPoolReward is a free data retrieval call binding the contract method 0xd6d7e5fb.
//
// Solidity: function getPendingPoolReward(uint8 _assetType, address _recipient) view returns(uint256, uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) GetPendingPoolReward(opts *bind.CallOpts, _assetType uint8, _recipient common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "getPendingPoolReward", _assetType, _recipient)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPendingPoolReward is a free data retrieval call binding the contract method 0xd6d7e5fb.
//
// Solidity: function getPendingPoolReward(uint8 _assetType, address _recipient) view returns(uint256, uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) GetPendingPoolReward(_assetType uint8, _recipient common.Address) (*big.Int, *big.Int, error) {
	return _NeoTokyoStaker.Contract.GetPendingPoolReward(&_NeoTokyoStaker.CallOpts, _assetType, _recipient)
}

// GetPendingPoolReward is a free data retrieval call binding the contract method 0xd6d7e5fb.
//
// Solidity: function getPendingPoolReward(uint8 _assetType, address _recipient) view returns(uint256, uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) GetPendingPoolReward(_assetType uint8, _recipient common.Address) (*big.Int, *big.Int, error) {
	return _NeoTokyoStaker.Contract.GetPendingPoolReward(&_NeoTokyoStaker.CallOpts, _assetType, _recipient)
}

// GetStakerPosition is a free data retrieval call binding the contract method 0x9ff260eb.
//
// Solidity: function getStakerPosition(address _staker, uint8 _assetType) view returns(uint256[])
func (_NeoTokyoStaker *NeoTokyoStakerCaller) GetStakerPosition(opts *bind.CallOpts, _staker common.Address, _assetType uint8) ([]*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "getStakerPosition", _staker, _assetType)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakerPosition is a free data retrieval call binding the contract method 0x9ff260eb.
//
// Solidity: function getStakerPosition(address _staker, uint8 _assetType) view returns(uint256[])
func (_NeoTokyoStaker *NeoTokyoStakerSession) GetStakerPosition(_staker common.Address, _assetType uint8) ([]*big.Int, error) {
	return _NeoTokyoStaker.Contract.GetStakerPosition(&_NeoTokyoStaker.CallOpts, _staker, _assetType)
}

// GetStakerPosition is a free data retrieval call binding the contract method 0x9ff260eb.
//
// Solidity: function getStakerPosition(address _staker, uint8 _assetType) view returns(uint256[])
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) GetStakerPosition(_staker common.Address, _assetType uint8) ([]*big.Int, error) {
	return _NeoTokyoStaker.Contract.GetStakerPosition(&_NeoTokyoStaker.CallOpts, _staker, _assetType)
}

// GetStakerPositions is a free data retrieval call binding the contract method 0x667a5154.
//
// Solidity: function getStakerPositions(address _staker) view returns(((uint256,uint256,uint256,uint256,uint256,bool)[],(uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)))
func (_NeoTokyoStaker *NeoTokyoStakerCaller) GetStakerPositions(opts *bind.CallOpts, _staker common.Address) (NeoTokyoStakerStakerPosition, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "getStakerPositions", _staker)

	if err != nil {
		return *new(NeoTokyoStakerStakerPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(NeoTokyoStakerStakerPosition)).(*NeoTokyoStakerStakerPosition)

	return out0, err

}

// GetStakerPositions is a free data retrieval call binding the contract method 0x667a5154.
//
// Solidity: function getStakerPositions(address _staker) view returns(((uint256,uint256,uint256,uint256,uint256,bool)[],(uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)))
func (_NeoTokyoStaker *NeoTokyoStakerSession) GetStakerPositions(_staker common.Address) (NeoTokyoStakerStakerPosition, error) {
	return _NeoTokyoStaker.Contract.GetStakerPositions(&_NeoTokyoStaker.CallOpts, _staker)
}

// GetStakerPositions is a free data retrieval call binding the contract method 0x667a5154.
//
// Solidity: function getStakerPositions(address _staker) view returns(((uint256,uint256,uint256,uint256,uint256,bool)[],(uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)))
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) GetStakerPositions(_staker common.Address) (NeoTokyoStakerStakerPosition, error) {
	return _NeoTokyoStaker.Contract.GetStakerPositions(&_NeoTokyoStaker.CallOpts, _staker)
}

// GetTotalEmissions is a free data retrieval call binding the contract method 0x2b76b11a.
//
// Solidity: function getTotalEmissions(uint8 _assetType, uint256 _from) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) GetTotalEmissions(opts *bind.CallOpts, _assetType uint8, _from *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "getTotalEmissions", _assetType, _from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalEmissions is a free data retrieval call binding the contract method 0x2b76b11a.
//
// Solidity: function getTotalEmissions(uint8 _assetType, uint256 _from) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) GetTotalEmissions(_assetType uint8, _from *big.Int) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.GetTotalEmissions(&_NeoTokyoStaker.CallOpts, _assetType, _from)
}

// GetTotalEmissions is a free data retrieval call binding the contract method 0x2b76b11a.
//
// Solidity: function getTotalEmissions(uint8 _assetType, uint256 _from) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) GetTotalEmissions(_assetType uint8, _from *big.Int) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.GetTotalEmissions(&_NeoTokyoStaker.CallOpts, _assetType, _from)
}

// HasRight is a free data retrieval call binding the contract method 0x8681d49c.
//
// Solidity: function hasRight(address _address, bytes32 _circumstance, bytes32 _right) view returns(bool)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) HasRight(opts *bind.CallOpts, _address common.Address, _circumstance [32]byte, _right [32]byte) (bool, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "hasRight", _address, _circumstance, _right)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRight is a free data retrieval call binding the contract method 0x8681d49c.
//
// Solidity: function hasRight(address _address, bytes32 _circumstance, bytes32 _right) view returns(bool)
func (_NeoTokyoStaker *NeoTokyoStakerSession) HasRight(_address common.Address, _circumstance [32]byte, _right [32]byte) (bool, error) {
	return _NeoTokyoStaker.Contract.HasRight(&_NeoTokyoStaker.CallOpts, _address, _circumstance, _right)
}

// HasRight is a free data retrieval call binding the contract method 0x8681d49c.
//
// Solidity: function hasRight(address _address, bytes32 _circumstance, bytes32 _right) view returns(bool)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) HasRight(_address common.Address, _circumstance [32]byte, _right [32]byte) (bool, error) {
	return _NeoTokyoStaker.Contract.HasRight(&_NeoTokyoStaker.CallOpts, _address, _circumstance, _right)
}

// HasRightUntil is a free data retrieval call binding the contract method 0x66a0e54d.
//
// Solidity: function hasRightUntil(address _address, bytes32 _circumstance, bytes32 _right) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) HasRightUntil(opts *bind.CallOpts, _address common.Address, _circumstance [32]byte, _right [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "hasRightUntil", _address, _circumstance, _right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HasRightUntil is a free data retrieval call binding the contract method 0x66a0e54d.
//
// Solidity: function hasRightUntil(address _address, bytes32 _circumstance, bytes32 _right) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) HasRightUntil(_address common.Address, _circumstance [32]byte, _right [32]byte) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.HasRightUntil(&_NeoTokyoStaker.CallOpts, _address, _circumstance, _right)
}

// HasRightUntil is a free data retrieval call binding the contract method 0x66a0e54d.
//
// Solidity: function hasRightUntil(address _address, bytes32 _circumstance, bytes32 _right) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) HasRightUntil(_address common.Address, _circumstance [32]byte, _right [32]byte) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.HasRightUntil(&_NeoTokyoStaker.CallOpts, _address, _circumstance, _right)
}

// IdentityCreditPoints is a free data retrieval call binding the contract method 0x65bdb43a.
//
// Solidity: function identityCreditPoints(string ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) IdentityCreditPoints(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "identityCreditPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdentityCreditPoints is a free data retrieval call binding the contract method 0x65bdb43a.
//
// Solidity: function identityCreditPoints(string ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) IdentityCreditPoints(arg0 string) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.IdentityCreditPoints(&_NeoTokyoStaker.CallOpts, arg0)
}

// IdentityCreditPoints is a free data retrieval call binding the contract method 0x65bdb43a.
//
// Solidity: function identityCreditPoints(string ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) IdentityCreditPoints(arg0 string) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.IdentityCreditPoints(&_NeoTokyoStaker.CallOpts, arg0)
}

// IdentityCreditYield is a free data retrieval call binding the contract method 0xeda50dbd.
//
// Solidity: function identityCreditYield(uint256 , string ) view returns(string)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) IdentityCreditYield(opts *bind.CallOpts, arg0 *big.Int, arg1 string) (string, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "identityCreditYield", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IdentityCreditYield is a free data retrieval call binding the contract method 0xeda50dbd.
//
// Solidity: function identityCreditYield(uint256 , string ) view returns(string)
func (_NeoTokyoStaker *NeoTokyoStakerSession) IdentityCreditYield(arg0 *big.Int, arg1 string) (string, error) {
	return _NeoTokyoStaker.Contract.IdentityCreditYield(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// IdentityCreditYield is a free data retrieval call binding the contract method 0xeda50dbd.
//
// Solidity: function identityCreditYield(uint256 , string ) view returns(string)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) IdentityCreditYield(arg0 *big.Int, arg1 string) (string, error) {
	return _NeoTokyoStaker.Contract.IdentityCreditYield(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// LastRewardTime is a free data retrieval call binding the contract method 0xde8e7d25.
//
// Solidity: function lastRewardTime(address , uint8 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) LastRewardTime(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "lastRewardTime", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardTime is a free data retrieval call binding the contract method 0xde8e7d25.
//
// Solidity: function lastRewardTime(address , uint8 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) LastRewardTime(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.LastRewardTime(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// LastRewardTime is a free data retrieval call binding the contract method 0xde8e7d25.
//
// Solidity: function lastRewardTime(address , uint8 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) LastRewardTime(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.LastRewardTime(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// LpLocked is a free data retrieval call binding the contract method 0xcf53154c.
//
// Solidity: function lpLocked() view returns(bool)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) LpLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "lpLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LpLocked is a free data retrieval call binding the contract method 0xcf53154c.
//
// Solidity: function lpLocked() view returns(bool)
func (_NeoTokyoStaker *NeoTokyoStakerSession) LpLocked() (bool, error) {
	return _NeoTokyoStaker.Contract.LpLocked(&_NeoTokyoStaker.CallOpts)
}

// LpLocked is a free data retrieval call binding the contract method 0xcf53154c.
//
// Solidity: function lpLocked() view returns(bool)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) LpLocked() (bool, error) {
	return _NeoTokyoStaker.Contract.LpLocked(&_NeoTokyoStaker.CallOpts)
}

// ManagerRight is a free data retrieval call binding the contract method 0xc5b16c59.
//
// Solidity: function managerRight(bytes32 ) view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) ManagerRight(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "managerRight", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ManagerRight is a free data retrieval call binding the contract method 0xc5b16c59.
//
// Solidity: function managerRight(bytes32 ) view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerSession) ManagerRight(arg0 [32]byte) ([32]byte, error) {
	return _NeoTokyoStaker.Contract.ManagerRight(&_NeoTokyoStaker.CallOpts, arg0)
}

// ManagerRight is a free data retrieval call binding the contract method 0xc5b16c59.
//
// Solidity: function managerRight(bytes32 ) view returns(bytes32)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) ManagerRight(arg0 [32]byte) ([32]byte, error) {
	return _NeoTokyoStaker.Contract.ManagerRight(&_NeoTokyoStaker.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerSession) Owner() (common.Address, error) {
	return _NeoTokyoStaker.Contract.Owner(&_NeoTokyoStaker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) Owner() (common.Address, error) {
	return _NeoTokyoStaker.Contract.Owner(&_NeoTokyoStaker.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x483ba44e.
//
// Solidity: function permissions(address , bytes32 , bytes32 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) Permissions(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "permissions", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Permissions is a free data retrieval call binding the contract method 0x483ba44e.
//
// Solidity: function permissions(address , bytes32 , bytes32 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) Permissions(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.Permissions(&_NeoTokyoStaker.CallOpts, arg0, arg1, arg2)
}

// Permissions is a free data retrieval call binding the contract method 0x483ba44e.
//
// Solidity: function permissions(address , bytes32 , bytes32 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) Permissions(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.Permissions(&_NeoTokyoStaker.CallOpts, arg0, arg1, arg2)
}

// RewardAccrued is a free data retrieval call binding the contract method 0x02bd0de4.
//
// Solidity: function rewardAccrued(address , uint8 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) RewardAccrued(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "rewardAccrued", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardAccrued is a free data retrieval call binding the contract method 0x02bd0de4.
//
// Solidity: function rewardAccrued(address , uint8 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) RewardAccrued(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.RewardAccrued(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// RewardAccrued is a free data retrieval call binding the contract method 0x02bd0de4.
//
// Solidity: function rewardAccrued(address , uint8 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) RewardAccrued(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.RewardAccrued(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// StakedS1 is a free data retrieval call binding the contract method 0x8b1f4740.
//
// Solidity: function stakedS1(address , uint256 ) view returns(uint256 stakedBytes, uint256 timelockEndTime, uint256 points, uint256 stakedVaultId, bool hasVault)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) StakedS1(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	StakedVaultId   *big.Int
	HasVault        bool
}, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "stakedS1", arg0, arg1)

	outstruct := new(struct {
		StakedBytes     *big.Int
		TimelockEndTime *big.Int
		Points          *big.Int
		StakedVaultId   *big.Int
		HasVault        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedBytes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TimelockEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Points = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StakedVaultId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HasVault = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakedS1 is a free data retrieval call binding the contract method 0x8b1f4740.
//
// Solidity: function stakedS1(address , uint256 ) view returns(uint256 stakedBytes, uint256 timelockEndTime, uint256 points, uint256 stakedVaultId, bool hasVault)
func (_NeoTokyoStaker *NeoTokyoStakerSession) StakedS1(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	StakedVaultId   *big.Int
	HasVault        bool
}, error) {
	return _NeoTokyoStaker.Contract.StakedS1(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// StakedS1 is a free data retrieval call binding the contract method 0x8b1f4740.
//
// Solidity: function stakedS1(address , uint256 ) view returns(uint256 stakedBytes, uint256 timelockEndTime, uint256 points, uint256 stakedVaultId, bool hasVault)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) StakedS1(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	StakedVaultId   *big.Int
	HasVault        bool
}, error) {
	return _NeoTokyoStaker.Contract.StakedS1(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// StakedS2 is a free data retrieval call binding the contract method 0x669896ce.
//
// Solidity: function stakedS2(address , uint256 ) view returns(uint256 stakedBytes, uint256 timelockEndTime, uint256 points)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) StakedS2(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
}, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "stakedS2", arg0, arg1)

	outstruct := new(struct {
		StakedBytes     *big.Int
		TimelockEndTime *big.Int
		Points          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedBytes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TimelockEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Points = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakedS2 is a free data retrieval call binding the contract method 0x669896ce.
//
// Solidity: function stakedS2(address , uint256 ) view returns(uint256 stakedBytes, uint256 timelockEndTime, uint256 points)
func (_NeoTokyoStaker *NeoTokyoStakerSession) StakedS2(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
}, error) {
	return _NeoTokyoStaker.Contract.StakedS2(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// StakedS2 is a free data retrieval call binding the contract method 0x669896ce.
//
// Solidity: function stakedS2(address , uint256 ) view returns(uint256 stakedBytes, uint256 timelockEndTime, uint256 points)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) StakedS2(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBytes     *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
}, error) {
	return _NeoTokyoStaker.Contract.StakedS2(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// StakerLPPosition is a free data retrieval call binding the contract method 0xa327660e.
//
// Solidity: function stakerLPPosition(address ) view returns(uint256 amount, uint256 timelockEndTime, uint256 points, uint256 multiplier)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) StakerLPPosition(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount          *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	Multiplier      *big.Int
}, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "stakerLPPosition", arg0)

	outstruct := new(struct {
		Amount          *big.Int
		TimelockEndTime *big.Int
		Points          *big.Int
		Multiplier      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TimelockEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Points = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Multiplier = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakerLPPosition is a free data retrieval call binding the contract method 0xa327660e.
//
// Solidity: function stakerLPPosition(address ) view returns(uint256 amount, uint256 timelockEndTime, uint256 points, uint256 multiplier)
func (_NeoTokyoStaker *NeoTokyoStakerSession) StakerLPPosition(arg0 common.Address) (struct {
	Amount          *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	Multiplier      *big.Int
}, error) {
	return _NeoTokyoStaker.Contract.StakerLPPosition(&_NeoTokyoStaker.CallOpts, arg0)
}

// StakerLPPosition is a free data retrieval call binding the contract method 0xa327660e.
//
// Solidity: function stakerLPPosition(address ) view returns(uint256 amount, uint256 timelockEndTime, uint256 points, uint256 multiplier)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) StakerLPPosition(arg0 common.Address) (struct {
	Amount          *big.Int
	TimelockEndTime *big.Int
	Points          *big.Int
	Multiplier      *big.Int
}, error) {
	return _NeoTokyoStaker.Contract.StakerLPPosition(&_NeoTokyoStaker.CallOpts, arg0)
}

// TimelockOptions is a free data retrieval call binding the contract method 0xcf1606ff.
//
// Solidity: function timelockOptions(uint8 , uint256 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) TimelockOptions(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "timelockOptions", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimelockOptions is a free data retrieval call binding the contract method 0xcf1606ff.
//
// Solidity: function timelockOptions(uint8 , uint256 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) TimelockOptions(arg0 uint8, arg1 *big.Int) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.TimelockOptions(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// TimelockOptions is a free data retrieval call binding the contract method 0xcf1606ff.
//
// Solidity: function timelockOptions(uint8 , uint256 ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) TimelockOptions(arg0 uint8, arg1 *big.Int) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.TimelockOptions(&_NeoTokyoStaker.CallOpts, arg0, arg1)
}

// VaultCreditMultiplier is a free data retrieval call binding the contract method 0xe221d1b5.
//
// Solidity: function vaultCreditMultiplier(string ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCaller) VaultCreditMultiplier(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _NeoTokyoStaker.contract.Call(opts, &out, "vaultCreditMultiplier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCreditMultiplier is a free data retrieval call binding the contract method 0xe221d1b5.
//
// Solidity: function vaultCreditMultiplier(string ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) VaultCreditMultiplier(arg0 string) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.VaultCreditMultiplier(&_NeoTokyoStaker.CallOpts, arg0)
}

// VaultCreditMultiplier is a free data retrieval call binding the contract method 0xe221d1b5.
//
// Solidity: function vaultCreditMultiplier(string ) view returns(uint256)
func (_NeoTokyoStaker *NeoTokyoStakerCallerSession) VaultCreditMultiplier(arg0 string) (*big.Int, error) {
	return _NeoTokyoStaker.Contract.VaultCreditMultiplier(&_NeoTokyoStaker.CallOpts, arg0)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _recipient) returns(uint256, uint256)
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ClaimReward(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "claimReward", _recipient)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _recipient) returns(uint256, uint256)
func (_NeoTokyoStaker *NeoTokyoStakerSession) ClaimReward(_recipient common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ClaimReward(&_NeoTokyoStaker.TransactOpts, _recipient)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _recipient) returns(uint256, uint256)
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ClaimReward(_recipient common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ClaimReward(&_NeoTokyoStaker.TransactOpts, _recipient)
}

// ConfigureCaps is a paid mutator transaction binding the contract method 0x7f7a0431.
//
// Solidity: function configureCaps(uint256 _vaultedCap, uint256 _unvaultedCap) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigureCaps(opts *bind.TransactOpts, _vaultedCap *big.Int, _unvaultedCap *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configureCaps", _vaultedCap, _unvaultedCap)
}

// ConfigureCaps is a paid mutator transaction binding the contract method 0x7f7a0431.
//
// Solidity: function configureCaps(uint256 _vaultedCap, uint256 _unvaultedCap) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigureCaps(_vaultedCap *big.Int, _unvaultedCap *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureCaps(&_NeoTokyoStaker.TransactOpts, _vaultedCap, _unvaultedCap)
}

// ConfigureCaps is a paid mutator transaction binding the contract method 0x7f7a0431.
//
// Solidity: function configureCaps(uint256 _vaultedCap, uint256 _unvaultedCap) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigureCaps(_vaultedCap *big.Int, _unvaultedCap *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureCaps(&_NeoTokyoStaker.TransactOpts, _vaultedCap, _unvaultedCap)
}

// ConfigureIdentityCreditPoints is a paid mutator transaction binding the contract method 0xfab17133.
//
// Solidity: function configureIdentityCreditPoints(string[] _identityCreditYields, uint256[] _points) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigureIdentityCreditPoints(opts *bind.TransactOpts, _identityCreditYields []string, _points []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configureIdentityCreditPoints", _identityCreditYields, _points)
}

// ConfigureIdentityCreditPoints is a paid mutator transaction binding the contract method 0xfab17133.
//
// Solidity: function configureIdentityCreditPoints(string[] _identityCreditYields, uint256[] _points) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigureIdentityCreditPoints(_identityCreditYields []string, _points []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureIdentityCreditPoints(&_NeoTokyoStaker.TransactOpts, _identityCreditYields, _points)
}

// ConfigureIdentityCreditPoints is a paid mutator transaction binding the contract method 0xfab17133.
//
// Solidity: function configureIdentityCreditPoints(string[] _identityCreditYields, uint256[] _points) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigureIdentityCreditPoints(_identityCreditYields []string, _points []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureIdentityCreditPoints(&_NeoTokyoStaker.TransactOpts, _identityCreditYields, _points)
}

// ConfigureIdentityCreditYields is a paid mutator transaction binding the contract method 0x213e9b46.
//
// Solidity: function configureIdentityCreditYields(uint256[] _citizenRewardRates, string[] _vaultRewardRates, string[] _identityCreditYields) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigureIdentityCreditYields(opts *bind.TransactOpts, _citizenRewardRates []*big.Int, _vaultRewardRates []string, _identityCreditYields []string) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configureIdentityCreditYields", _citizenRewardRates, _vaultRewardRates, _identityCreditYields)
}

// ConfigureIdentityCreditYields is a paid mutator transaction binding the contract method 0x213e9b46.
//
// Solidity: function configureIdentityCreditYields(uint256[] _citizenRewardRates, string[] _vaultRewardRates, string[] _identityCreditYields) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigureIdentityCreditYields(_citizenRewardRates []*big.Int, _vaultRewardRates []string, _identityCreditYields []string) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureIdentityCreditYields(&_NeoTokyoStaker.TransactOpts, _citizenRewardRates, _vaultRewardRates, _identityCreditYields)
}

// ConfigureIdentityCreditYields is a paid mutator transaction binding the contract method 0x213e9b46.
//
// Solidity: function configureIdentityCreditYields(uint256[] _citizenRewardRates, string[] _vaultRewardRates, string[] _identityCreditYields) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigureIdentityCreditYields(_citizenRewardRates []*big.Int, _vaultRewardRates []string, _identityCreditYields []string) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureIdentityCreditYields(&_NeoTokyoStaker.TransactOpts, _citizenRewardRates, _vaultRewardRates, _identityCreditYields)
}

// ConfigureLP is a paid mutator transaction binding the contract method 0x8b595c38.
//
// Solidity: function configureLP(address _lp) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigureLP(opts *bind.TransactOpts, _lp common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configureLP", _lp)
}

// ConfigureLP is a paid mutator transaction binding the contract method 0x8b595c38.
//
// Solidity: function configureLP(address _lp) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigureLP(_lp common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureLP(&_NeoTokyoStaker.TransactOpts, _lp)
}

// ConfigureLP is a paid mutator transaction binding the contract method 0x8b595c38.
//
// Solidity: function configureLP(address _lp) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigureLP(_lp common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureLP(&_NeoTokyoStaker.TransactOpts, _lp)
}

// ConfigurePools is a paid mutator transaction binding the contract method 0xc25409b5.
//
// Solidity: function configurePools((uint8,uint256,(uint128,uint128)[])[] _inputs) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigurePools(opts *bind.TransactOpts, _inputs []NeoTokyoStakerPoolConfigurationInput) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configurePools", _inputs)
}

// ConfigurePools is a paid mutator transaction binding the contract method 0xc25409b5.
//
// Solidity: function configurePools((uint8,uint256,(uint128,uint128)[])[] _inputs) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigurePools(_inputs []NeoTokyoStakerPoolConfigurationInput) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigurePools(&_NeoTokyoStaker.TransactOpts, _inputs)
}

// ConfigurePools is a paid mutator transaction binding the contract method 0xc25409b5.
//
// Solidity: function configurePools((uint8,uint256,(uint128,uint128)[])[] _inputs) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigurePools(_inputs []NeoTokyoStakerPoolConfigurationInput) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigurePools(&_NeoTokyoStaker.TransactOpts, _inputs)
}

// ConfigureTimelockOptions is a paid mutator transaction binding the contract method 0xac11e40e.
//
// Solidity: function configureTimelockOptions(uint8 _assetType, uint256[] _timelockIds, uint256[] _encodedSettings) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigureTimelockOptions(opts *bind.TransactOpts, _assetType uint8, _timelockIds []*big.Int, _encodedSettings []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configureTimelockOptions", _assetType, _timelockIds, _encodedSettings)
}

// ConfigureTimelockOptions is a paid mutator transaction binding the contract method 0xac11e40e.
//
// Solidity: function configureTimelockOptions(uint8 _assetType, uint256[] _timelockIds, uint256[] _encodedSettings) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigureTimelockOptions(_assetType uint8, _timelockIds []*big.Int, _encodedSettings []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureTimelockOptions(&_NeoTokyoStaker.TransactOpts, _assetType, _timelockIds, _encodedSettings)
}

// ConfigureTimelockOptions is a paid mutator transaction binding the contract method 0xac11e40e.
//
// Solidity: function configureTimelockOptions(uint8 _assetType, uint256[] _timelockIds, uint256[] _encodedSettings) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigureTimelockOptions(_assetType uint8, _timelockIds []*big.Int, _encodedSettings []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureTimelockOptions(&_NeoTokyoStaker.TransactOpts, _assetType, _timelockIds, _encodedSettings)
}

// ConfigureVaultCreditMultipliers is a paid mutator transaction binding the contract method 0xc8257634.
//
// Solidity: function configureVaultCreditMultipliers(string[] _vaultCreditMultipliers, uint256[] _multipliers) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) ConfigureVaultCreditMultipliers(opts *bind.TransactOpts, _vaultCreditMultipliers []string, _multipliers []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "configureVaultCreditMultipliers", _vaultCreditMultipliers, _multipliers)
}

// ConfigureVaultCreditMultipliers is a paid mutator transaction binding the contract method 0xc8257634.
//
// Solidity: function configureVaultCreditMultipliers(string[] _vaultCreditMultipliers, uint256[] _multipliers) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) ConfigureVaultCreditMultipliers(_vaultCreditMultipliers []string, _multipliers []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureVaultCreditMultipliers(&_NeoTokyoStaker.TransactOpts, _vaultCreditMultipliers, _multipliers)
}

// ConfigureVaultCreditMultipliers is a paid mutator transaction binding the contract method 0xc8257634.
//
// Solidity: function configureVaultCreditMultipliers(string[] _vaultCreditMultipliers, uint256[] _multipliers) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) ConfigureVaultCreditMultipliers(_vaultCreditMultipliers []string, _multipliers []*big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.ConfigureVaultCreditMultipliers(&_NeoTokyoStaker.TransactOpts, _vaultCreditMultipliers, _multipliers)
}

// LockLP is a paid mutator transaction binding the contract method 0x763d7353.
//
// Solidity: function lockLP() returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) LockLP(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "lockLP")
}

// LockLP is a paid mutator transaction binding the contract method 0x763d7353.
//
// Solidity: function lockLP() returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) LockLP() (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.LockLP(&_NeoTokyoStaker.TransactOpts)
}

// LockLP is a paid mutator transaction binding the contract method 0x763d7353.
//
// Solidity: function lockLP() returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) LockLP() (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.LockLP(&_NeoTokyoStaker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.RenounceOwnership(&_NeoTokyoStaker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.RenounceOwnership(&_NeoTokyoStaker.TransactOpts)
}

// SetManagerRight is a paid mutator transaction binding the contract method 0xcc2af308.
//
// Solidity: function setManagerRight(bytes32 _managedRight, bytes32 _managerRight) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) SetManagerRight(opts *bind.TransactOpts, _managedRight [32]byte, _managerRight [32]byte) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "setManagerRight", _managedRight, _managerRight)
}

// SetManagerRight is a paid mutator transaction binding the contract method 0xcc2af308.
//
// Solidity: function setManagerRight(bytes32 _managedRight, bytes32 _managerRight) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) SetManagerRight(_managedRight [32]byte, _managerRight [32]byte) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.SetManagerRight(&_NeoTokyoStaker.TransactOpts, _managedRight, _managerRight)
}

// SetManagerRight is a paid mutator transaction binding the contract method 0xcc2af308.
//
// Solidity: function setManagerRight(bytes32 _managedRight, bytes32 _managerRight) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) SetManagerRight(_managedRight [32]byte, _managerRight [32]byte) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.SetManagerRight(&_NeoTokyoStaker.TransactOpts, _managedRight, _managerRight)
}

// SetPermit is a paid mutator transaction binding the contract method 0xcf64d4c2.
//
// Solidity: function setPermit(address _address, bytes32 _circumstance, bytes32 _right, uint256 _expirationTime) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) SetPermit(opts *bind.TransactOpts, _address common.Address, _circumstance [32]byte, _right [32]byte, _expirationTime *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "setPermit", _address, _circumstance, _right, _expirationTime)
}

// SetPermit is a paid mutator transaction binding the contract method 0xcf64d4c2.
//
// Solidity: function setPermit(address _address, bytes32 _circumstance, bytes32 _right, uint256 _expirationTime) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) SetPermit(_address common.Address, _circumstance [32]byte, _right [32]byte, _expirationTime *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.SetPermit(&_NeoTokyoStaker.TransactOpts, _address, _circumstance, _right, _expirationTime)
}

// SetPermit is a paid mutator transaction binding the contract method 0xcf64d4c2.
//
// Solidity: function setPermit(address _address, bytes32 _circumstance, bytes32 _right, uint256 _expirationTime) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) SetPermit(_address common.Address, _circumstance [32]byte, _right [32]byte, _expirationTime *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.SetPermit(&_NeoTokyoStaker.TransactOpts, _address, _circumstance, _right, _expirationTime)
}

// Stake is a paid mutator transaction binding the contract method 0x32bfbc97.
//
// Solidity: function stake(uint8 _assetType, uint256 _timelockId, uint256 , uint256 , uint256 ) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) Stake(opts *bind.TransactOpts, _assetType uint8, _timelockId *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "stake", _assetType, _timelockId, arg2, arg3, arg4)
}

// Stake is a paid mutator transaction binding the contract method 0x32bfbc97.
//
// Solidity: function stake(uint8 _assetType, uint256 _timelockId, uint256 , uint256 , uint256 ) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) Stake(_assetType uint8, _timelockId *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.Stake(&_NeoTokyoStaker.TransactOpts, _assetType, _timelockId, arg2, arg3, arg4)
}

// Stake is a paid mutator transaction binding the contract method 0x32bfbc97.
//
// Solidity: function stake(uint8 _assetType, uint256 _timelockId, uint256 , uint256 , uint256 ) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) Stake(_assetType uint8, _timelockId *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.Stake(&_NeoTokyoStaker.TransactOpts, _assetType, _timelockId, arg2, arg3, arg4)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.TransferOwnership(&_NeoTokyoStaker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.TransferOwnership(&_NeoTokyoStaker.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _assetType, uint256 ) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactor) Withdraw(opts *bind.TransactOpts, _assetType uint8, arg1 *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.contract.Transact(opts, "withdraw", _assetType, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _assetType, uint256 ) returns()
func (_NeoTokyoStaker *NeoTokyoStakerSession) Withdraw(_assetType uint8, arg1 *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.Withdraw(&_NeoTokyoStaker.TransactOpts, _assetType, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _assetType, uint256 ) returns()
func (_NeoTokyoStaker *NeoTokyoStakerTransactorSession) Withdraw(_assetType uint8, arg1 *big.Int) (*types.Transaction, error) {
	return _NeoTokyoStaker.Contract.Withdraw(&_NeoTokyoStaker.TransactOpts, _assetType, arg1)
}

// NeoTokyoStakerClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the NeoTokyoStaker contract.
type NeoTokyoStakerClaimIterator struct {
	Event *NeoTokyoStakerClaim // Event containing the contract specifics and raw log

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
func (it *NeoTokyoStakerClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoTokyoStakerClaim)
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
		it.Event = new(NeoTokyoStakerClaim)
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
func (it *NeoTokyoStakerClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoTokyoStakerClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoTokyoStakerClaim represents a Claim event raised by the NeoTokyoStaker contract.
type NeoTokyoStakerClaim struct {
	Recipient common.Address
	Reward    *big.Int
	Tax       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed recipient, uint256 reward, uint256 tax)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) FilterClaim(opts *bind.FilterOpts, recipient []common.Address) (*NeoTokyoStakerClaimIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.FilterLogs(opts, "Claim", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerClaimIterator{contract: _NeoTokyoStaker.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed recipient, uint256 reward, uint256 tax)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *NeoTokyoStakerClaim, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.WatchLogs(opts, "Claim", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoTokyoStakerClaim)
				if err := _NeoTokyoStaker.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed recipient, uint256 reward, uint256 tax)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) ParseClaim(log types.Log) (*NeoTokyoStakerClaim, error) {
	event := new(NeoTokyoStakerClaim)
	if err := _NeoTokyoStaker.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoTokyoStakerManagementUpdatedIterator is returned from FilterManagementUpdated and is used to iterate over the raw logs and unpacked data for ManagementUpdated events raised by the NeoTokyoStaker contract.
type NeoTokyoStakerManagementUpdatedIterator struct {
	Event *NeoTokyoStakerManagementUpdated // Event containing the contract specifics and raw log

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
func (it *NeoTokyoStakerManagementUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoTokyoStakerManagementUpdated)
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
		it.Event = new(NeoTokyoStakerManagementUpdated)
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
func (it *NeoTokyoStakerManagementUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoTokyoStakerManagementUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoTokyoStakerManagementUpdated represents a ManagementUpdated event raised by the NeoTokyoStaker contract.
type NeoTokyoStakerManagementUpdated struct {
	Manager      common.Address
	ManagedRight [32]byte
	ManagerRight [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterManagementUpdated is a free log retrieval operation binding the contract event 0xad26b90be8a18bd2262e914f6fd4919c42f9dd6a0d07a15fa728ec603a836a88.
//
// Solidity: event ManagementUpdated(address indexed manager, bytes32 indexed managedRight, bytes32 indexed managerRight)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) FilterManagementUpdated(opts *bind.FilterOpts, manager []common.Address, managedRight [][32]byte, managerRight [][32]byte) (*NeoTokyoStakerManagementUpdatedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}
	var managedRightRule []interface{}
	for _, managedRightItem := range managedRight {
		managedRightRule = append(managedRightRule, managedRightItem)
	}
	var managerRightRule []interface{}
	for _, managerRightItem := range managerRight {
		managerRightRule = append(managerRightRule, managerRightItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.FilterLogs(opts, "ManagementUpdated", managerRule, managedRightRule, managerRightRule)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerManagementUpdatedIterator{contract: _NeoTokyoStaker.contract, event: "ManagementUpdated", logs: logs, sub: sub}, nil
}

// WatchManagementUpdated is a free log subscription operation binding the contract event 0xad26b90be8a18bd2262e914f6fd4919c42f9dd6a0d07a15fa728ec603a836a88.
//
// Solidity: event ManagementUpdated(address indexed manager, bytes32 indexed managedRight, bytes32 indexed managerRight)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) WatchManagementUpdated(opts *bind.WatchOpts, sink chan<- *NeoTokyoStakerManagementUpdated, manager []common.Address, managedRight [][32]byte, managerRight [][32]byte) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}
	var managedRightRule []interface{}
	for _, managedRightItem := range managedRight {
		managedRightRule = append(managedRightRule, managedRightItem)
	}
	var managerRightRule []interface{}
	for _, managerRightItem := range managerRight {
		managerRightRule = append(managerRightRule, managerRightItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.WatchLogs(opts, "ManagementUpdated", managerRule, managedRightRule, managerRightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoTokyoStakerManagementUpdated)
				if err := _NeoTokyoStaker.contract.UnpackLog(event, "ManagementUpdated", log); err != nil {
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

// ParseManagementUpdated is a log parse operation binding the contract event 0xad26b90be8a18bd2262e914f6fd4919c42f9dd6a0d07a15fa728ec603a836a88.
//
// Solidity: event ManagementUpdated(address indexed manager, bytes32 indexed managedRight, bytes32 indexed managerRight)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) ParseManagementUpdated(log types.Log) (*NeoTokyoStakerManagementUpdated, error) {
	event := new(NeoTokyoStakerManagementUpdated)
	if err := _NeoTokyoStaker.contract.UnpackLog(event, "ManagementUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoTokyoStakerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NeoTokyoStaker contract.
type NeoTokyoStakerOwnershipTransferredIterator struct {
	Event *NeoTokyoStakerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NeoTokyoStakerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoTokyoStakerOwnershipTransferred)
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
		it.Event = new(NeoTokyoStakerOwnershipTransferred)
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
func (it *NeoTokyoStakerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoTokyoStakerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoTokyoStakerOwnershipTransferred represents a OwnershipTransferred event raised by the NeoTokyoStaker contract.
type NeoTokyoStakerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NeoTokyoStakerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerOwnershipTransferredIterator{contract: _NeoTokyoStaker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NeoTokyoStakerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoTokyoStakerOwnershipTransferred)
				if err := _NeoTokyoStaker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) ParseOwnershipTransferred(log types.Log) (*NeoTokyoStakerOwnershipTransferred, error) {
	event := new(NeoTokyoStakerOwnershipTransferred)
	if err := _NeoTokyoStaker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoTokyoStakerPermitUpdatedIterator is returned from FilterPermitUpdated and is used to iterate over the raw logs and unpacked data for PermitUpdated events raised by the NeoTokyoStaker contract.
type NeoTokyoStakerPermitUpdatedIterator struct {
	Event *NeoTokyoStakerPermitUpdated // Event containing the contract specifics and raw log

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
func (it *NeoTokyoStakerPermitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoTokyoStakerPermitUpdated)
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
		it.Event = new(NeoTokyoStakerPermitUpdated)
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
func (it *NeoTokyoStakerPermitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoTokyoStakerPermitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoTokyoStakerPermitUpdated represents a PermitUpdated event raised by the NeoTokyoStaker contract.
type NeoTokyoStakerPermitUpdated struct {
	Updater        common.Address
	Updatee        common.Address
	Circumstance   [32]byte
	Role           [32]byte
	ExpirationTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPermitUpdated is a free log retrieval operation binding the contract event 0x71b8ef6d2e182fa6ca30442059cc10398330b3e0561fd4ecc7232b62a8678cb6.
//
// Solidity: event PermitUpdated(address indexed updater, address indexed updatee, bytes32 circumstance, bytes32 indexed role, uint256 expirationTime)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) FilterPermitUpdated(opts *bind.FilterOpts, updater []common.Address, updatee []common.Address, role [][32]byte) (*NeoTokyoStakerPermitUpdatedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var updateeRule []interface{}
	for _, updateeItem := range updatee {
		updateeRule = append(updateeRule, updateeItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.FilterLogs(opts, "PermitUpdated", updaterRule, updateeRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerPermitUpdatedIterator{contract: _NeoTokyoStaker.contract, event: "PermitUpdated", logs: logs, sub: sub}, nil
}

// WatchPermitUpdated is a free log subscription operation binding the contract event 0x71b8ef6d2e182fa6ca30442059cc10398330b3e0561fd4ecc7232b62a8678cb6.
//
// Solidity: event PermitUpdated(address indexed updater, address indexed updatee, bytes32 circumstance, bytes32 indexed role, uint256 expirationTime)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) WatchPermitUpdated(opts *bind.WatchOpts, sink chan<- *NeoTokyoStakerPermitUpdated, updater []common.Address, updatee []common.Address, role [][32]byte) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var updateeRule []interface{}
	for _, updateeItem := range updatee {
		updateeRule = append(updateeRule, updateeItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.WatchLogs(opts, "PermitUpdated", updaterRule, updateeRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoTokyoStakerPermitUpdated)
				if err := _NeoTokyoStaker.contract.UnpackLog(event, "PermitUpdated", log); err != nil {
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

// ParsePermitUpdated is a log parse operation binding the contract event 0x71b8ef6d2e182fa6ca30442059cc10398330b3e0561fd4ecc7232b62a8678cb6.
//
// Solidity: event PermitUpdated(address indexed updater, address indexed updatee, bytes32 circumstance, bytes32 indexed role, uint256 expirationTime)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) ParsePermitUpdated(log types.Log) (*NeoTokyoStakerPermitUpdated, error) {
	event := new(NeoTokyoStakerPermitUpdated)
	if err := _NeoTokyoStaker.contract.UnpackLog(event, "PermitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoTokyoStakerStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the NeoTokyoStaker contract.
type NeoTokyoStakerStakeIterator struct {
	Event *NeoTokyoStakerStake // Event containing the contract specifics and raw log

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
func (it *NeoTokyoStakerStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoTokyoStakerStake)
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
		it.Event = new(NeoTokyoStakerStake)
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
func (it *NeoTokyoStakerStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoTokyoStakerStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoTokyoStakerStake represents a Stake event raised by the NeoTokyoStaker contract.
type NeoTokyoStakerStake struct {
	Staker          common.Address
	Asset           common.Address
	TimelockOption  *big.Int
	AmountOrTokenId *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed staker, address indexed asset, uint256 timelockOption, uint256 amountOrTokenId)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) FilterStake(opts *bind.FilterOpts, staker []common.Address, asset []common.Address) (*NeoTokyoStakerStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.FilterLogs(opts, "Stake", stakerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerStakeIterator{contract: _NeoTokyoStaker.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed staker, address indexed asset, uint256 timelockOption, uint256 amountOrTokenId)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *NeoTokyoStakerStake, staker []common.Address, asset []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.WatchLogs(opts, "Stake", stakerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoTokyoStakerStake)
				if err := _NeoTokyoStaker.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed staker, address indexed asset, uint256 timelockOption, uint256 amountOrTokenId)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) ParseStake(log types.Log) (*NeoTokyoStakerStake, error) {
	event := new(NeoTokyoStakerStake)
	if err := _NeoTokyoStaker.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoTokyoStakerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the NeoTokyoStaker contract.
type NeoTokyoStakerWithdrawIterator struct {
	Event *NeoTokyoStakerWithdraw // Event containing the contract specifics and raw log

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
func (it *NeoTokyoStakerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoTokyoStakerWithdraw)
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
		it.Event = new(NeoTokyoStakerWithdraw)
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
func (it *NeoTokyoStakerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoTokyoStakerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoTokyoStakerWithdraw represents a Withdraw event raised by the NeoTokyoStaker contract.
type NeoTokyoStakerWithdraw struct {
	Caller          common.Address
	Asset           common.Address
	AmountOrTokenId *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed caller, address indexed asset, uint256 amountOrTokenId)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, asset []common.Address) (*NeoTokyoStakerWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.FilterLogs(opts, "Withdraw", callerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &NeoTokyoStakerWithdrawIterator{contract: _NeoTokyoStaker.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed caller, address indexed asset, uint256 amountOrTokenId)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NeoTokyoStakerWithdraw, caller []common.Address, asset []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NeoTokyoStaker.contract.WatchLogs(opts, "Withdraw", callerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoTokyoStakerWithdraw)
				if err := _NeoTokyoStaker.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed caller, address indexed asset, uint256 amountOrTokenId)
func (_NeoTokyoStaker *NeoTokyoStakerFilterer) ParseWithdraw(log types.Log) (*NeoTokyoStakerWithdraw, error) {
	event := new(NeoTokyoStakerWithdraw)
	if err := _NeoTokyoStaker.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


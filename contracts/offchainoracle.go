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

// OffchainOracleMetaData contains all meta data concerning the OffchainOracle contract.
var OffchainOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"_multiWrapper\",\"type\":\"address\"},{\"internalType\":\"contractIOracle[]\",\"name\":\"existingOracles\",\"type\":\"address[]\"},{\"internalType\":\"enumOffchainOracle.OracleType[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"existingConnectors\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"wBase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"\",\"outputs\":null,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"addConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"allConnectors\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSrcWrappers\",\"type\":\"bool\"}],\"name\":\"getRateToEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSrcWrappers\",\"type\":\"bool\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"customConnectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"thresholdFilter\",\"type\":\"uint256\"}],\"name\":\"getRateToEthWithCustomConnectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSrcWrappers\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"thresholdFilter\",\"type\":\"uint256\"}],\"name\":\"getRateToEthWithThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"customConnectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"thresholdFilter\",\"type\":\"uint256\"}],\"name\":\"getRateWithCustomConnectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"thresholdFilter\",\"type\":\"uint256\"}],\"name\":\"getRateWithThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiWrapper\",\"outputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"contractIOracle[]\",\"name\":\"allOracles\",\"type\":\"address[]\"},{\"internalType\":\"enumOffchainOracle.OracleType[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"removeConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"removeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"_multiWrapper\",\"type\":\"address\"}],\"name\":\"setMultiWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OffchainOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OffchainOracleMetaData.ABI instead.
var OffchainOracleABI = OffchainOracleMetaData.ABI

// OffchainOracle is an auto generated Go binding around an Ethereum contract.
type OffchainOracle struct {
	OffchainOracleCaller     // Read-only binding to the contract
	OffchainOracleTransactor // Write-only binding to the contract
	OffchainOracleFilterer   // Log filterer for contract events
}

// OffchainOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OffchainOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OffchainOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OffchainOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OffchainOracleSession struct {
	Contract     *OffchainOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OffchainOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OffchainOracleCallerSession struct {
	Contract *OffchainOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OffchainOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OffchainOracleTransactorSession struct {
	Contract     *OffchainOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OffchainOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OffchainOracleRaw struct {
	Contract *OffchainOracle // Generic contract binding to access the raw methods on
}

// OffchainOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OffchainOracleCallerRaw struct {
	Contract *OffchainOracleCaller // Generic read-only contract binding to access the raw methods on
}

// OffchainOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OffchainOracleTransactorRaw struct {
	Contract *OffchainOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOffchainOracle creates a new instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracle(address common.Address, backend bind.ContractBackend) (*OffchainOracle, error) {
	contract, err := bindOffchainOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffchainOracle{OffchainOracleCaller: OffchainOracleCaller{contract: contract}, OffchainOracleTransactor: OffchainOracleTransactor{contract: contract}, OffchainOracleFilterer: OffchainOracleFilterer{contract: contract}}, nil
}

// NewOffchainOracleCaller creates a new read-only instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracleCaller(address common.Address, caller bind.ContractCaller) (*OffchainOracleCaller, error) {
	contract, err := bindOffchainOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleCaller{contract: contract}, nil
}

// NewOffchainOracleTransactor creates a new write-only instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OffchainOracleTransactor, error) {
	contract, err := bindOffchainOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleTransactor{contract: contract}, nil
}

// NewOffchainOracleFilterer creates a new log filterer instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OffchainOracleFilterer, error) {
	contract, err := bindOffchainOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleFilterer{contract: contract}, nil
}

// bindOffchainOracle binds a generic wrapper to an already deployed contract.
func bindOffchainOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OffchainOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainOracle *OffchainOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainOracle.Contract.OffchainOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainOracle *OffchainOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainOracle.Contract.OffchainOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainOracle *OffchainOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainOracle.Contract.OffchainOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainOracle *OffchainOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainOracle *OffchainOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainOracle *OffchainOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainOracle.Contract.contract.Transact(opts, method, params...)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OffchainOracle *OffchainOracleCaller) Connectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "connectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OffchainOracle *OffchainOracleSession) Connectors() ([]common.Address, error) {
	return _OffchainOracle.Contract.Connectors(&_OffchainOracle.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OffchainOracle *OffchainOracleCallerSession) Connectors() ([]common.Address, error) {
	return _OffchainOracle.Contract.Connectors(&_OffchainOracle.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRate(opts *bind.CallOpts, srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRate", srcToken, dstToken, useWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRate(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRate(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRateToEth(opts *bind.CallOpts, srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRateToEth", srcToken, useSrcWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEth(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEth(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers)
}

// GetRateToEthWithCustomConnectors is a free data retrieval call binding the contract method 0xade8b048.
//
// Solidity: function getRateToEthWithCustomConnectors(address srcToken, bool useSrcWrappers, address[] customConnectors, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRateToEthWithCustomConnectors(opts *bind.CallOpts, srcToken common.Address, useSrcWrappers bool, customConnectors []common.Address, thresholdFilter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRateToEthWithCustomConnectors", srcToken, useSrcWrappers, customConnectors, thresholdFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateToEthWithCustomConnectors is a free data retrieval call binding the contract method 0xade8b048.
//
// Solidity: function getRateToEthWithCustomConnectors(address srcToken, bool useSrcWrappers, address[] customConnectors, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRateToEthWithCustomConnectors(srcToken common.Address, useSrcWrappers bool, customConnectors []common.Address, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEthWithCustomConnectors(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers, customConnectors, thresholdFilter)
}

// GetRateToEthWithCustomConnectors is a free data retrieval call binding the contract method 0xade8b048.
//
// Solidity: function getRateToEthWithCustomConnectors(address srcToken, bool useSrcWrappers, address[] customConnectors, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRateToEthWithCustomConnectors(srcToken common.Address, useSrcWrappers bool, customConnectors []common.Address, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEthWithCustomConnectors(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers, customConnectors, thresholdFilter)
}

// GetRateToEthWithThreshold is a free data retrieval call binding the contract method 0x78159aae.
//
// Solidity: function getRateToEthWithThreshold(address srcToken, bool useSrcWrappers, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRateToEthWithThreshold(opts *bind.CallOpts, srcToken common.Address, useSrcWrappers bool, thresholdFilter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRateToEthWithThreshold", srcToken, useSrcWrappers, thresholdFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateToEthWithThreshold is a free data retrieval call binding the contract method 0x78159aae.
//
// Solidity: function getRateToEthWithThreshold(address srcToken, bool useSrcWrappers, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRateToEthWithThreshold(srcToken common.Address, useSrcWrappers bool, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEthWithThreshold(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers, thresholdFilter)
}

// GetRateToEthWithThreshold is a free data retrieval call binding the contract method 0x78159aae.
//
// Solidity: function getRateToEthWithThreshold(address srcToken, bool useSrcWrappers, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRateToEthWithThreshold(srcToken common.Address, useSrcWrappers bool, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEthWithThreshold(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers, thresholdFilter)
}

// GetRateWithCustomConnectors is a free data retrieval call binding the contract method 0x6f9293b9.
//
// Solidity: function getRateWithCustomConnectors(address srcToken, address dstToken, bool useWrappers, address[] customConnectors, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRateWithCustomConnectors(opts *bind.CallOpts, srcToken common.Address, dstToken common.Address, useWrappers bool, customConnectors []common.Address, thresholdFilter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRateWithCustomConnectors", srcToken, dstToken, useWrappers, customConnectors, thresholdFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateWithCustomConnectors is a free data retrieval call binding the contract method 0x6f9293b9.
//
// Solidity: function getRateWithCustomConnectors(address srcToken, address dstToken, bool useWrappers, address[] customConnectors, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRateWithCustomConnectors(srcToken common.Address, dstToken common.Address, useWrappers bool, customConnectors []common.Address, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateWithCustomConnectors(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers, customConnectors, thresholdFilter)
}

// GetRateWithCustomConnectors is a free data retrieval call binding the contract method 0x6f9293b9.
//
// Solidity: function getRateWithCustomConnectors(address srcToken, address dstToken, bool useWrappers, address[] customConnectors, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRateWithCustomConnectors(srcToken common.Address, dstToken common.Address, useWrappers bool, customConnectors []common.Address, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateWithCustomConnectors(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers, customConnectors, thresholdFilter)
}

// GetRateWithThreshold is a free data retrieval call binding the contract method 0x6744d6c7.
//
// Solidity: function getRateWithThreshold(address srcToken, address dstToken, bool useWrappers, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRateWithThreshold(opts *bind.CallOpts, srcToken common.Address, dstToken common.Address, useWrappers bool, thresholdFilter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRateWithThreshold", srcToken, dstToken, useWrappers, thresholdFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateWithThreshold is a free data retrieval call binding the contract method 0x6744d6c7.
//
// Solidity: function getRateWithThreshold(address srcToken, address dstToken, bool useWrappers, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRateWithThreshold(srcToken common.Address, dstToken common.Address, useWrappers bool, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateWithThreshold(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers, thresholdFilter)
}

// GetRateWithThreshold is a free data retrieval call binding the contract method 0x6744d6c7.
//
// Solidity: function getRateWithThreshold(address srcToken, address dstToken, bool useWrappers, uint256 thresholdFilter) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRateWithThreshold(srcToken common.Address, dstToken common.Address, useWrappers bool, thresholdFilter *big.Int) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateWithThreshold(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers, thresholdFilter)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OffchainOracle *OffchainOracleCaller) MultiWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "multiWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OffchainOracle *OffchainOracleSession) MultiWrapper() (common.Address, error) {
	return _OffchainOracle.Contract.MultiWrapper(&_OffchainOracle.CallOpts)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OffchainOracle *OffchainOracleCallerSession) MultiWrapper() (common.Address, error) {
	return _OffchainOracle.Contract.MultiWrapper(&_OffchainOracle.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OffchainOracle *OffchainOracleCaller) Oracles(opts *bind.CallOpts) (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "oracles")

	outstruct := new(struct {
		AllOracles  []common.Address
		OracleTypes []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllOracles = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OracleTypes = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OffchainOracle *OffchainOracleSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OffchainOracle.Contract.Oracles(&_OffchainOracle.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OffchainOracle *OffchainOracleCallerSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OffchainOracle.Contract.Oracles(&_OffchainOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainOracle *OffchainOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainOracle *OffchainOracleSession) Owner() (common.Address, error) {
	return _OffchainOracle.Contract.Owner(&_OffchainOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainOracle *OffchainOracleCallerSession) Owner() (common.Address, error) {
	return _OffchainOracle.Contract.Owner(&_OffchainOracle.CallOpts)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactor) AddConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "addConnector", connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddConnector(&_OffchainOracle.TransactOpts, connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddConnector(&_OffchainOracle.TransactOpts, connector)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactor) AddOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "addOracle", oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactor) RemoveConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "removeConnector", connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveConnector(&_OffchainOracle.TransactOpts, connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveConnector(&_OffchainOracle.TransactOpts, connector)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactor) RemoveOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "removeOracle", oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OffchainOracle *OffchainOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OffchainOracle *OffchainOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _OffchainOracle.Contract.RenounceOwnership(&_OffchainOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OffchainOracle *OffchainOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OffchainOracle.Contract.RenounceOwnership(&_OffchainOracle.TransactOpts)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OffchainOracle *OffchainOracleTransactor) SetMultiWrapper(opts *bind.TransactOpts, _multiWrapper common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "setMultiWrapper", _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OffchainOracle *OffchainOracleSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.SetMultiWrapper(&_OffchainOracle.TransactOpts, _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.SetMultiWrapper(&_OffchainOracle.TransactOpts, _multiWrapper)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OffchainOracle *OffchainOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OffchainOracle *OffchainOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.TransferOwnership(&_OffchainOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.TransferOwnership(&_OffchainOracle.TransactOpts, newOwner)
}


// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolswaptest

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

// IPoolManagerSwapParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerSwapParams struct {
	ZeroForOne        bool
	AmountSpecified   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// PoolSwapTestTestSettings is an auto generated low-level Go binding around an user-defined struct.
type PoolSwapTestTestSettings struct {
	TakeClaims      bool
	SettleUsingBurn bool
}

// PoolswaptestMetaData contains all meta data concerning the Poolswaptest contract.
var PoolswaptestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"testSettings\",\"type\":\"tuple\",\"internalType\":\"structPoolSwapTest.TestSettings\",\"components\":[{\"name\":\"takeClaims\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"settleUsingBurn\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unlockCallback\",\"inputs\":[{\"name\":\"rawData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"NoSwapOccurred\",\"inputs\":[]}]",
}

// PoolswaptestABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolswaptestMetaData.ABI instead.
var PoolswaptestABI = PoolswaptestMetaData.ABI

// Poolswaptest is an auto generated Go binding around an Ethereum contract.
type Poolswaptest struct {
	PoolswaptestCaller     // Read-only binding to the contract
	PoolswaptestTransactor // Write-only binding to the contract
	PoolswaptestFilterer   // Log filterer for contract events
}

// PoolswaptestCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolswaptestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolswaptestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolswaptestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolswaptestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolswaptestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolswaptestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolswaptestSession struct {
	Contract     *Poolswaptest     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolswaptestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolswaptestCallerSession struct {
	Contract *PoolswaptestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PoolswaptestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolswaptestTransactorSession struct {
	Contract     *PoolswaptestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PoolswaptestRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolswaptestRaw struct {
	Contract *Poolswaptest // Generic contract binding to access the raw methods on
}

// PoolswaptestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolswaptestCallerRaw struct {
	Contract *PoolswaptestCaller // Generic read-only contract binding to access the raw methods on
}

// PoolswaptestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolswaptestTransactorRaw struct {
	Contract *PoolswaptestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolswaptest creates a new instance of Poolswaptest, bound to a specific deployed contract.
func NewPoolswaptest(address common.Address, backend bind.ContractBackend) (*Poolswaptest, error) {
	contract, err := bindPoolswaptest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolswaptest{PoolswaptestCaller: PoolswaptestCaller{contract: contract}, PoolswaptestTransactor: PoolswaptestTransactor{contract: contract}, PoolswaptestFilterer: PoolswaptestFilterer{contract: contract}}, nil
}

// NewPoolswaptestCaller creates a new read-only instance of Poolswaptest, bound to a specific deployed contract.
func NewPoolswaptestCaller(address common.Address, caller bind.ContractCaller) (*PoolswaptestCaller, error) {
	contract, err := bindPoolswaptest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolswaptestCaller{contract: contract}, nil
}

// NewPoolswaptestTransactor creates a new write-only instance of Poolswaptest, bound to a specific deployed contract.
func NewPoolswaptestTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolswaptestTransactor, error) {
	contract, err := bindPoolswaptest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolswaptestTransactor{contract: contract}, nil
}

// NewPoolswaptestFilterer creates a new log filterer instance of Poolswaptest, bound to a specific deployed contract.
func NewPoolswaptestFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolswaptestFilterer, error) {
	contract, err := bindPoolswaptest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolswaptestFilterer{contract: contract}, nil
}

// bindPoolswaptest binds a generic wrapper to an already deployed contract.
func bindPoolswaptest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolswaptestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolswaptest *PoolswaptestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolswaptest.Contract.PoolswaptestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolswaptest *PoolswaptestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolswaptest.Contract.PoolswaptestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolswaptest *PoolswaptestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolswaptest.Contract.PoolswaptestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolswaptest *PoolswaptestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolswaptest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolswaptest *PoolswaptestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolswaptest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolswaptest *PoolswaptestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolswaptest.Contract.contract.Transact(opts, method, params...)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Poolswaptest *PoolswaptestCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolswaptest.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Poolswaptest *PoolswaptestSession) Manager() (common.Address, error) {
	return _Poolswaptest.Contract.Manager(&_Poolswaptest.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Poolswaptest *PoolswaptestCallerSession) Manager() (common.Address, error) {
	return _Poolswaptest.Contract.Manager(&_Poolswaptest.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x2229d0b4.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, (bool,bool) testSettings, bytes hookData) payable returns(int256 delta)
func (_Poolswaptest *PoolswaptestTransactor) Swap(opts *bind.TransactOpts, key PoolKey, params IPoolManagerSwapParams, testSettings PoolSwapTestTestSettings, hookData []byte) (*types.Transaction, error) {
	return _Poolswaptest.contract.Transact(opts, "swap", key, params, testSettings, hookData)
}

// Swap is a paid mutator transaction binding the contract method 0x2229d0b4.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, (bool,bool) testSettings, bytes hookData) payable returns(int256 delta)
func (_Poolswaptest *PoolswaptestSession) Swap(key PoolKey, params IPoolManagerSwapParams, testSettings PoolSwapTestTestSettings, hookData []byte) (*types.Transaction, error) {
	return _Poolswaptest.Contract.Swap(&_Poolswaptest.TransactOpts, key, params, testSettings, hookData)
}

// Swap is a paid mutator transaction binding the contract method 0x2229d0b4.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, (bool,bool) testSettings, bytes hookData) payable returns(int256 delta)
func (_Poolswaptest *PoolswaptestTransactorSession) Swap(key PoolKey, params IPoolManagerSwapParams, testSettings PoolSwapTestTestSettings, hookData []byte) (*types.Transaction, error) {
	return _Poolswaptest.Contract.Swap(&_Poolswaptest.TransactOpts, key, params, testSettings, hookData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_Poolswaptest *PoolswaptestTransactor) UnlockCallback(opts *bind.TransactOpts, rawData []byte) (*types.Transaction, error) {
	return _Poolswaptest.contract.Transact(opts, "unlockCallback", rawData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_Poolswaptest *PoolswaptestSession) UnlockCallback(rawData []byte) (*types.Transaction, error) {
	return _Poolswaptest.Contract.UnlockCallback(&_Poolswaptest.TransactOpts, rawData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_Poolswaptest *PoolswaptestTransactorSession) UnlockCallback(rawData []byte) (*types.Transaction, error) {
	return _Poolswaptest.Contract.UnlockCallback(&_Poolswaptest.TransactOpts, rawData)
}

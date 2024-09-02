// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// HooksPermissions is an auto generated low-level Go binding around an user-defined struct.
type HooksPermissions struct {
	BeforeInitialize                bool
	AfterInitialize                 bool
	BeforeAddLiquidity              bool
	AfterAddLiquidity               bool
	BeforeRemoveLiquidity           bool
	AfterRemoveLiquidity            bool
	BeforeSwap                      bool
	AfterSwap                       bool
	BeforeDonate                    bool
	AfterDonate                     bool
	BeforeSwapReturnDelta           bool
	AfterSwapReturnDelta            bool
	AfterAddLiquidityReturnDelta    bool
	AfterRemoveLiquidityReturnDelta bool
}

// IPoolManagerModifyLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerModifyLiquidityParams struct {
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
}

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_poolManager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterAddLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterRemoveLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int128\",\"internalType\":\"int128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterSwapCount\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeAddLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeAddLiquidityCount\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveLiquidityCount\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BeforeSwapDelta\"},{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeSwapCount\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHookPermissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHooks.Permissions\",\"components\":[{\"name\":\"beforeInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockCallback\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"HookNotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LockFailure\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPoolManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]}],",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// AfterSwapCount is a free data retrieval call binding the contract method 0xdacdbe77.
//
// Solidity: function afterSwapCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCaller) AfterSwapCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "afterSwapCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AfterSwapCount is a free data retrieval call binding the contract method 0xdacdbe77.
//
// Solidity: function afterSwapCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterSession) AfterSwapCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.AfterSwapCount(&_Counter.CallOpts, arg0)
}

// AfterSwapCount is a free data retrieval call binding the contract method 0xdacdbe77.
//
// Solidity: function afterSwapCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCallerSession) AfterSwapCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.AfterSwapCount(&_Counter.CallOpts, arg0)
}

// BeforeAddLiquidityCount is a free data retrieval call binding the contract method 0xc6939957.
//
// Solidity: function beforeAddLiquidityCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCaller) BeforeAddLiquidityCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "beforeAddLiquidityCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeforeAddLiquidityCount is a free data retrieval call binding the contract method 0xc6939957.
//
// Solidity: function beforeAddLiquidityCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterSession) BeforeAddLiquidityCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.BeforeAddLiquidityCount(&_Counter.CallOpts, arg0)
}

// BeforeAddLiquidityCount is a free data retrieval call binding the contract method 0xc6939957.
//
// Solidity: function beforeAddLiquidityCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCallerSession) BeforeAddLiquidityCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.BeforeAddLiquidityCount(&_Counter.CallOpts, arg0)
}

// BeforeRemoveLiquidityCount is a free data retrieval call binding the contract method 0xdb6a9451.
//
// Solidity: function beforeRemoveLiquidityCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCaller) BeforeRemoveLiquidityCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "beforeRemoveLiquidityCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeforeRemoveLiquidityCount is a free data retrieval call binding the contract method 0xdb6a9451.
//
// Solidity: function beforeRemoveLiquidityCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterSession) BeforeRemoveLiquidityCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.BeforeRemoveLiquidityCount(&_Counter.CallOpts, arg0)
}

// BeforeRemoveLiquidityCount is a free data retrieval call binding the contract method 0xdb6a9451.
//
// Solidity: function beforeRemoveLiquidityCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCallerSession) BeforeRemoveLiquidityCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.BeforeRemoveLiquidityCount(&_Counter.CallOpts, arg0)
}

// BeforeSwapCount is a free data retrieval call binding the contract method 0x5c338d32.
//
// Solidity: function beforeSwapCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCaller) BeforeSwapCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "beforeSwapCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeforeSwapCount is a free data retrieval call binding the contract method 0x5c338d32.
//
// Solidity: function beforeSwapCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterSession) BeforeSwapCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.BeforeSwapCount(&_Counter.CallOpts, arg0)
}

// BeforeSwapCount is a free data retrieval call binding the contract method 0x5c338d32.
//
// Solidity: function beforeSwapCount(bytes32 ) view returns(uint256 count)
func (_Counter *CounterCallerSession) BeforeSwapCount(arg0 [32]byte) (*big.Int, error) {
	return _Counter.Contract.BeforeSwapCount(&_Counter.CallOpts, arg0)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Counter *CounterCaller) GetHookPermissions(opts *bind.CallOpts) (HooksPermissions, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getHookPermissions")

	if err != nil {
		return *new(HooksPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new(HooksPermissions)).(*HooksPermissions)

	return out0, err

}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Counter *CounterSession) GetHookPermissions() (HooksPermissions, error) {
	return _Counter.Contract.GetHookPermissions(&_Counter.CallOpts)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Counter *CounterCallerSession) GetHookPermissions() (HooksPermissions, error) {
	return _Counter.Contract.GetHookPermissions(&_Counter.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Counter *CounterCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Counter *CounterSession) PoolManager() (common.Address, error) {
	return _Counter.Contract.PoolManager(&_Counter.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Counter *CounterCallerSession) PoolManager() (common.Address, error) {
	return _Counter.Contract.PoolManager(&_Counter.CallOpts)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Counter *CounterTransactor) AfterAddLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "afterAddLiquidity", arg0, arg1, arg2, arg3, arg4)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Counter *CounterSession) AfterAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterAddLiquidity(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Counter *CounterTransactorSession) AfterAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterAddLiquidity(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Counter *CounterTransactor) AfterDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "afterDonate", arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Counter *CounterSession) AfterDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterDonate(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Counter *CounterTransactorSession) AfterDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterDonate(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Counter *CounterTransactor) AfterInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "afterInitialize", arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Counter *CounterSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterInitialize(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Counter *CounterTransactorSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterInitialize(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Counter *CounterTransactor) AfterRemoveLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "afterRemoveLiquidity", arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Counter *CounterSession) AfterRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterRemoveLiquidity(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Counter *CounterTransactorSession) AfterRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterRemoveLiquidity(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Counter *CounterTransactor) AfterSwap(opts *bind.TransactOpts, arg0 common.Address, key PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "afterSwap", arg0, key, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Counter *CounterSession) AfterSwap(arg0 common.Address, key PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterSwap(&_Counter.TransactOpts, arg0, key, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Counter *CounterTransactorSession) AfterSwap(arg0 common.Address, key PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.AfterSwap(&_Counter.TransactOpts, arg0, key, arg2, arg3, arg4)
}

// BeforeAddLiquidity is a paid mutator transaction binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Counter *CounterTransactor) BeforeAddLiquidity(opts *bind.TransactOpts, arg0 common.Address, key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "beforeAddLiquidity", arg0, key, arg2, arg3)
}

// BeforeAddLiquidity is a paid mutator transaction binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Counter *CounterSession) BeforeAddLiquidity(arg0 common.Address, key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeAddLiquidity(&_Counter.TransactOpts, arg0, key, arg2, arg3)
}

// BeforeAddLiquidity is a paid mutator transaction binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Counter *CounterTransactorSession) BeforeAddLiquidity(arg0 common.Address, key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeAddLiquidity(&_Counter.TransactOpts, arg0, key, arg2, arg3)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Counter *CounterTransactor) BeforeDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "beforeDonate", arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Counter *CounterSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeDonate(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Counter *CounterTransactorSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeDonate(&_Counter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Counter *CounterTransactor) BeforeInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "beforeInitialize", arg0, arg1, arg2, arg3)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Counter *CounterSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeInitialize(&_Counter.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Counter *CounterTransactorSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeInitialize(&_Counter.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Counter *CounterTransactor) BeforeRemoveLiquidity(opts *bind.TransactOpts, arg0 common.Address, key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "beforeRemoveLiquidity", arg0, key, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Counter *CounterSession) BeforeRemoveLiquidity(arg0 common.Address, key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeRemoveLiquidity(&_Counter.TransactOpts, arg0, key, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Counter *CounterTransactorSession) BeforeRemoveLiquidity(arg0 common.Address, key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeRemoveLiquidity(&_Counter.TransactOpts, arg0, key, arg2, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) , bytes ) returns(bytes4, int256, uint24)
func (_Counter *CounterTransactor) BeforeSwap(opts *bind.TransactOpts, arg0 common.Address, key PoolKey, arg2 IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "beforeSwap", arg0, key, arg2, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) , bytes ) returns(bytes4, int256, uint24)
func (_Counter *CounterSession) BeforeSwap(arg0 common.Address, key PoolKey, arg2 IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeSwap(&_Counter.TransactOpts, arg0, key, arg2, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) , bytes ) returns(bytes4, int256, uint24)
func (_Counter *CounterTransactorSession) BeforeSwap(arg0 common.Address, key PoolKey, arg2 IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Counter.Contract.BeforeSwap(&_Counter.TransactOpts, arg0, key, arg2, arg3)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Counter *CounterTransactor) UnlockCallback(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "unlockCallback", data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Counter *CounterSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Counter.Contract.UnlockCallback(&_Counter.TransactOpts, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Counter *CounterTransactorSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Counter.Contract.UnlockCallback(&_Counter.TransactOpts, data)
}

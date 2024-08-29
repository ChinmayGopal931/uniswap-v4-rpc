// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csmm

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

// CsmmMetaData contains all meta data concerning the Csmm contract.
var CsmmMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"poolManager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"amountEach\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterAddLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterRemoveLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int128\",\"internalType\":\"int128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeAddLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"beforeDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BeforeSwapDelta\"},{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getHookPermissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHooks.Permissions\",\"components\":[{\"name\":\"beforeInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockCallback\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"AddLiquidityThroughHook\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookNotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LockFailure\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPoolManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]}],",
}

// CsmmABI is the input ABI used to generate the binding from.
// Deprecated: Use CsmmMetaData.ABI instead.
var CsmmABI = CsmmMetaData.ABI

// Csmm is an auto generated Go binding around an Ethereum contract.
type Csmm struct {
	CsmmCaller     // Read-only binding to the contract
	CsmmTransactor // Write-only binding to the contract
	CsmmFilterer   // Log filterer for contract events
}

// CsmmCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsmmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsmmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsmmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsmmSession struct {
	Contract     *Csmm             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsmmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsmmCallerSession struct {
	Contract *CsmmCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CsmmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsmmTransactorSession struct {
	Contract     *CsmmTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsmmRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsmmRaw struct {
	Contract *Csmm // Generic contract binding to access the raw methods on
}

// CsmmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsmmCallerRaw struct {
	Contract *CsmmCaller // Generic read-only contract binding to access the raw methods on
}

// CsmmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsmmTransactorRaw struct {
	Contract *CsmmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsmm creates a new instance of Csmm, bound to a specific deployed contract.
func NewCsmm(address common.Address, backend bind.ContractBackend) (*Csmm, error) {
	contract, err := bindCsmm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csmm{CsmmCaller: CsmmCaller{contract: contract}, CsmmTransactor: CsmmTransactor{contract: contract}, CsmmFilterer: CsmmFilterer{contract: contract}}, nil
}

// NewCsmmCaller creates a new read-only instance of Csmm, bound to a specific deployed contract.
func NewCsmmCaller(address common.Address, caller bind.ContractCaller) (*CsmmCaller, error) {
	contract, err := bindCsmm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsmmCaller{contract: contract}, nil
}

// NewCsmmTransactor creates a new write-only instance of Csmm, bound to a specific deployed contract.
func NewCsmmTransactor(address common.Address, transactor bind.ContractTransactor) (*CsmmTransactor, error) {
	contract, err := bindCsmm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsmmTransactor{contract: contract}, nil
}

// NewCsmmFilterer creates a new log filterer instance of Csmm, bound to a specific deployed contract.
func NewCsmmFilterer(address common.Address, filterer bind.ContractFilterer) (*CsmmFilterer, error) {
	contract, err := bindCsmm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsmmFilterer{contract: contract}, nil
}

// bindCsmm binds a generic wrapper to an already deployed contract.
func bindCsmm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsmmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csmm *CsmmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csmm.Contract.CsmmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csmm *CsmmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmm.Contract.CsmmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csmm *CsmmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csmm.Contract.CsmmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csmm *CsmmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csmm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csmm *CsmmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csmm *CsmmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csmm.Contract.contract.Transact(opts, method, params...)
}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) pure returns(bytes4)
func (_Csmm *CsmmCaller) BeforeAddLiquidity(opts *bind.CallOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Csmm.contract.Call(opts, &out, "beforeAddLiquidity", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) pure returns(bytes4)
func (_Csmm *CsmmSession) BeforeAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _Csmm.Contract.BeforeAddLiquidity(&_Csmm.CallOpts, arg0, arg1, arg2, arg3)
}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) pure returns(bytes4)
func (_Csmm *CsmmCallerSession) BeforeAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _Csmm.Contract.BeforeAddLiquidity(&_Csmm.CallOpts, arg0, arg1, arg2, arg3)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Csmm *CsmmCaller) GetHookPermissions(opts *bind.CallOpts) (HooksPermissions, error) {
	var out []interface{}
	err := _Csmm.contract.Call(opts, &out, "getHookPermissions")

	if err != nil {
		return *new(HooksPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new(HooksPermissions)).(*HooksPermissions)

	return out0, err

}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Csmm *CsmmSession) GetHookPermissions() (HooksPermissions, error) {
	return _Csmm.Contract.GetHookPermissions(&_Csmm.CallOpts)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Csmm *CsmmCallerSession) GetHookPermissions() (HooksPermissions, error) {
	return _Csmm.Contract.GetHookPermissions(&_Csmm.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Csmm *CsmmCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmm.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Csmm *CsmmSession) PoolManager() (common.Address, error) {
	return _Csmm.Contract.PoolManager(&_Csmm.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Csmm *CsmmCallerSession) PoolManager() (common.Address, error) {
	return _Csmm.Contract.PoolManager(&_Csmm.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9e436f57.
//
// Solidity: function addLiquidity((address,address,uint24,int24,address) key, uint256 amountEach) returns()
func (_Csmm *CsmmTransactor) AddLiquidity(opts *bind.TransactOpts, key PoolKey, amountEach *big.Int) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "addLiquidity", key, amountEach)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9e436f57.
//
// Solidity: function addLiquidity((address,address,uint24,int24,address) key, uint256 amountEach) returns()
func (_Csmm *CsmmSession) AddLiquidity(key PoolKey, amountEach *big.Int) (*types.Transaction, error) {
	return _Csmm.Contract.AddLiquidity(&_Csmm.TransactOpts, key, amountEach)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9e436f57.
//
// Solidity: function addLiquidity((address,address,uint24,int24,address) key, uint256 amountEach) returns()
func (_Csmm *CsmmTransactorSession) AddLiquidity(key PoolKey, amountEach *big.Int) (*types.Transaction, error) {
	return _Csmm.Contract.AddLiquidity(&_Csmm.TransactOpts, key, amountEach)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Csmm *CsmmTransactor) AfterAddLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "afterAddLiquidity", arg0, arg1, arg2, arg3, arg4)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Csmm *CsmmSession) AfterAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterAddLiquidity(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Csmm *CsmmTransactorSession) AfterAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterAddLiquidity(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactor) AfterDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "afterDonate", arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Csmm *CsmmSession) AfterDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterDonate(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactorSession) AfterDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterDonate(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactor) AfterInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "afterInitialize", arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Csmm *CsmmSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterInitialize(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactorSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterInitialize(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Csmm *CsmmTransactor) AfterRemoveLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "afterRemoveLiquidity", arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Csmm *CsmmSession) AfterRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterRemoveLiquidity(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Csmm *CsmmTransactorSession) AfterRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterRemoveLiquidity(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Csmm *CsmmTransactor) AfterSwap(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "afterSwap", arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Csmm *CsmmSession) AfterSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterSwap(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Csmm *CsmmTransactorSession) AfterSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.AfterSwap(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactor) BeforeDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "beforeDonate", arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Csmm *CsmmSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeDonate(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactorSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeDonate(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactor) BeforeInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "beforeInitialize", arg0, arg1, arg2, arg3)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Csmm *CsmmSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeInitialize(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactorSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeInitialize(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactor) BeforeRemoveLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "beforeRemoveLiquidity", arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Csmm *CsmmSession) BeforeRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeRemoveLiquidity(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Csmm *CsmmTransactorSession) BeforeRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeRemoveLiquidity(&_Csmm.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes ) returns(bytes4, int256, uint24)
func (_Csmm *CsmmTransactor) BeforeSwap(opts *bind.TransactOpts, arg0 common.Address, key PoolKey, params IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "beforeSwap", arg0, key, params, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes ) returns(bytes4, int256, uint24)
func (_Csmm *CsmmSession) BeforeSwap(arg0 common.Address, key PoolKey, params IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeSwap(&_Csmm.TransactOpts, arg0, key, params, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes ) returns(bytes4, int256, uint24)
func (_Csmm *CsmmTransactorSession) BeforeSwap(arg0 common.Address, key PoolKey, params IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Csmm.Contract.BeforeSwap(&_Csmm.TransactOpts, arg0, key, params, arg3)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Csmm *CsmmTransactor) UnlockCallback(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Csmm.contract.Transact(opts, "unlockCallback", data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Csmm *CsmmSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Csmm.Contract.UnlockCallback(&_Csmm.TransactOpts, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Csmm *CsmmTransactorSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Csmm.Contract.UnlockCallback(&_Csmm.TransactOpts, data)
}

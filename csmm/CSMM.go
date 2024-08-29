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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"poolManager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"amountEach\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterAddLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterRemoveLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int128\",\"internalType\":\"int128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeAddLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"beforeDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BeforeSwapDelta\"},{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getHookPermissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHooks.Permissions\",\"components\":[{\"name\":\"beforeInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockCallback\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"AddLiquidityThroughHook\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookNotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LockFailure\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPoolManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]}],",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) pure returns(bytes4)
func (_Contracts *ContractsCaller) BeforeAddLiquidity(opts *bind.CallOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "beforeAddLiquidity", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) pure returns(bytes4)
func (_Contracts *ContractsSession) BeforeAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _Contracts.Contract.BeforeAddLiquidity(&_Contracts.CallOpts, arg0, arg1, arg2, arg3)
}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) pure returns(bytes4)
func (_Contracts *ContractsCallerSession) BeforeAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _Contracts.Contract.BeforeAddLiquidity(&_Contracts.CallOpts, arg0, arg1, arg2, arg3)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Contracts *ContractsCaller) GetHookPermissions(opts *bind.CallOpts) (HooksPermissions, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getHookPermissions")

	if err != nil {
		return *new(HooksPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new(HooksPermissions)).(*HooksPermissions)

	return out0, err

}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Contracts *ContractsSession) GetHookPermissions() (HooksPermissions, error) {
	return _Contracts.Contract.GetHookPermissions(&_Contracts.CallOpts)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Contracts *ContractsCallerSession) GetHookPermissions() (HooksPermissions, error) {
	return _Contracts.Contract.GetHookPermissions(&_Contracts.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Contracts *ContractsCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Contracts *ContractsSession) PoolManager() (common.Address, error) {
	return _Contracts.Contract.PoolManager(&_Contracts.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Contracts *ContractsCallerSession) PoolManager() (common.Address, error) {
	return _Contracts.Contract.PoolManager(&_Contracts.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9e436f57.
//
// Solidity: function addLiquidity((address,address,uint24,int24,address) key, uint256 amountEach) returns()
func (_Contracts *ContractsTransactor) AddLiquidity(opts *bind.TransactOpts, key PoolKey, amountEach *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addLiquidity", key, amountEach)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9e436f57.
//
// Solidity: function addLiquidity((address,address,uint24,int24,address) key, uint256 amountEach) returns()
func (_Contracts *ContractsSession) AddLiquidity(key PoolKey, amountEach *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddLiquidity(&_Contracts.TransactOpts, key, amountEach)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9e436f57.
//
// Solidity: function addLiquidity((address,address,uint24,int24,address) key, uint256 amountEach) returns()
func (_Contracts *ContractsTransactorSession) AddLiquidity(key PoolKey, amountEach *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddLiquidity(&_Contracts.TransactOpts, key, amountEach)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Contracts *ContractsTransactor) AfterAddLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "afterAddLiquidity", arg0, arg1, arg2, arg3, arg4)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Contracts *ContractsSession) AfterAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterAddLiquidity(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x5a2a8100.
//
// Solidity: function afterAddLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Contracts *ContractsTransactorSession) AfterAddLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterAddLiquidity(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) AfterDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "afterDonate", arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) AfterDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterDonate(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) AfterDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterDonate(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) AfterInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "afterInitialize", arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterInitialize(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0xa910f80f.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterInitialize(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Contracts *ContractsTransactor) AfterRemoveLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "afterRemoveLiquidity", arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Contracts *ContractsSession) AfterRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterRemoveLiquidity(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x8db2b652.
//
// Solidity: function afterRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , int256 , bytes ) returns(bytes4, int256)
func (_Contracts *ContractsTransactorSession) AfterRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterRemoveLiquidity(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Contracts *ContractsTransactor) AfterSwap(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "afterSwap", arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Contracts *ContractsSession) AfterSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterSwap(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) returns(bytes4, int128)
func (_Contracts *ContractsTransactorSession) AfterSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AfterSwap(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) BeforeDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "beforeDonate", arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeDonate(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeDonate(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) BeforeInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "beforeInitialize", arg0, arg1, arg2, arg3)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeInitialize(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeInitialize is a paid mutator transaction binding the contract method 0x3440d820.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeInitialize(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) BeforeRemoveLiquidity(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "beforeRemoveLiquidity", arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) BeforeRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeRemoveLiquidity(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeRemoveLiquidity is a paid mutator transaction binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address , (address,address,uint24,int24,address) , (int24,int24,int256,bytes32) , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) BeforeRemoveLiquidity(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeRemoveLiquidity(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes ) returns(bytes4, int256, uint24)
func (_Contracts *ContractsTransactor) BeforeSwap(opts *bind.TransactOpts, arg0 common.Address, key PoolKey, params IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "beforeSwap", arg0, key, params, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes ) returns(bytes4, int256, uint24)
func (_Contracts *ContractsSession) BeforeSwap(arg0 common.Address, key PoolKey, params IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeSwap(&_Contracts.TransactOpts, arg0, key, params, arg3)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address , (address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes ) returns(bytes4, int256, uint24)
func (_Contracts *ContractsTransactorSession) BeforeSwap(arg0 common.Address, key PoolKey, params IPoolManagerSwapParams, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.BeforeSwap(&_Contracts.TransactOpts, arg0, key, params, arg3)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Contracts *ContractsTransactor) UnlockCallback(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlockCallback", data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Contracts *ContractsSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UnlockCallback(&_Contracts.TransactOpts, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Contracts *ContractsTransactorSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UnlockCallback(&_Contracts.TransactOpts, data)
}

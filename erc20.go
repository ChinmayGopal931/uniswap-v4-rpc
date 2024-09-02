package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var erc20ABI abi.ABI

func init() {
	var err error
	erc20ABI, err = abi.JSON(strings.NewReader(erc20ABIJson))
	if err != nil {
		panic(err)
	}
}

type ERC20 struct {
	address  common.Address
	contract *bind.BoundContract
}

// func newERC20(address common.Address) (*ERC20, error) {
// 	contract := bind.NewBoundContract(address, erc20ABI, EthClient, EthClient, EthClient)
// 	return &ERC20{
// 		address:  address,
// 		contract: contract,
// 	}, nil
// }

func newERC20(tokenAddress common.Address) (*ERC20, error) {
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJson))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(tokenAddress, erc20ABI, ethClient, ethClient, ethClient)
	return &ERC20{contract: contract}, nil
}

func (e *ERC20) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := e.contract.Call(opts, &out, "balanceOf", account)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

func (e *ERC20) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return e.contract.Transact(opts, "transfer", recipient, amount)
}

const erc20ABIJson = `[
	{
		"constant": true,
		"inputs": [{"name": "_owner", "type": "address"}],
		"name": "balanceOf",
		"outputs": [{"name": "balance", "type": "uint256"}],
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [{"name": "_to", "type": "address"}, {"name": "_value", "type": "uint256"}],
		"name": "transfer",
		"outputs": [{"name": "success", "type": "bool"}],
		"type": "function"
	}
]`

var EthClient *ethclient.Client

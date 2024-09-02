package main

// import (
// 	"strings"

// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// var (
// 	ethClient     *ethclient.Client
// 	swapRouterABI abi.ABI
// )

// func initEthereumClient() error {
// 	var err error
// 	ethClient, err = ethclient.Dial(config.EthereumNodeURL)
// 	if err != nil {
// 		return err
// 	}

// 	// Set the EthClient in erc20.go
// 	EthClient = ethClient

// 	swapRouterABI, err = abi.JSON(strings.NewReader(config.SwapRouterABIJSON))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

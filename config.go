// package main

// import (
// 	"crypto/ecdsa"
// 	"log"
// 	"os"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// )

// var config struct {
// 	EthereumNodeURL   string
// 	SwapRouterAddress common.Address
// 	HooksAddress      common.Address
// 	PrivateKey        *ecdsa.PrivateKey
// 	SwapRouterABIJSON string
// }

// func init() {
// 	config.EthereumNodeURL = getEnv("ETHEREUM_NODE_URL", "http://localhost:8545")
// 	config.SwapRouterAddress = common.HexToAddress(getEnv("SWAP_ROUTER_ADDRESS", "0xa82fF9aFd8f496c3d6ac40E2a0F282E47488CFc9"))
// 	config.HooksAddress = common.HexToAddress(getEnv("HOOKS_ADDRESS", "0x620687Fd59b0e5733aFdD8C8ADDd931f92a64aC0"))

// 	privateKeyHex := getEnv("PRIVATE_KEY", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		log.Fatalf("Failed to parse private key: %v", err)
// 	}
// 	config.PrivateKey = privateKey

// 	config.SwapRouterABIJSON = `[{"type":"constructor","inputs":[{"name":"_manager","type":"address","internalType":"contract IPoolManager"}],"stateMutability":"nonpayable"},{"type":"function","name":"manager","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IPoolManager"}],"stateMutability":"view"},{"type":"function","name":"swap","inputs":[{"name":"key","type":"tuple","internalType":"struct PoolKey","components":[{"name":"currency0","type":"address","internalType":"Currency"},{"name":"currency1","type":"address","internalType":"Currency"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"tickSpacing","type":"int24","internalType":"int24"},{"name":"hooks","type":"address","internalType":"contract IHooks"}]},{"name":"params","type":"tuple","internalType":"struct IPoolManager.SwapParams","components":[{"name":"zeroForOne","type":"bool","internalType":"bool"},{"name":"amountSpecified","type":"int256","internalType":"int256"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]},{"name":"testSettings","type":"tuple","internalType":"struct PoolSwapTest.TestSettings","components":[{"name":"takeClaims","type":"bool","internalType":"bool"},{"name":"settleUsingBurn","type":"bool","internalType":"bool"}]},{"name":"hookData","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"delta","type":"int256","internalType":"BalanceDelta"}],"stateMutability":"payable"},{"type":"function","name":"unlockCallback","inputs":[{"name":"rawData","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"bytes","internalType":"bytes"}],"stateMutability":"nonpayable"},{"type":"error","name":"NoSwapOccurred","inputs":[]}]`
// }

// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

package main

import (
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var config struct {
	EthereumNodeURL                string
	SwapRouterAddress              common.Address
	PoolModifyLiquidityTestAddress common.Address
	HooksAddress                   common.Address
	PrivateKey                     *ecdsa.PrivateKey
	SwapRouterABIJSON              string
	PoolModifyLiquidityTestABIJSON string
}

func init() {
	config.EthereumNodeURL = getEnv("ETHEREUM_NODE_URL", "http://localhost:8545")
	config.SwapRouterAddress = common.HexToAddress(getEnv("SWAP_ROUTER_ADDRESS", "0x998abeb3E57409262aE5b751f60747921B33613E"))
	config.PoolModifyLiquidityTestAddress = common.HexToAddress(getEnv("POOL_MODIFY_LIQUIDITY_TEST_ADDRESS", "0x4631BCAbD6dF18D94796344963cB60d44a4136b6"))
	config.HooksAddress = common.HexToAddress(getEnv("HOOKS_ADDRESS", "0x3C2836994A1EFA4ef01f5cD3d867E46485944AC0"))

	privateKeyHex := getEnv("PRIVATE_KEY", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	config.PrivateKey = privateKey

	// config.SwapRouterABIJSON = `[{"type":"constructor","inputs":[{"name":"_manager","type":"address","internalType":"contract IPoolManager"}],"stateMutability":"nonpayable"},{"type":"function","name":"manager","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IPoolManager"}],"stateMutability":"view"},{"type":"function","name":"swap","inputs":[{"name":"key","type":"tuple","internalType":"struct PoolKey","components":[{"name":"currency0","type":"address","internalType":"Currency"},{"name":"currency1","type":"address","internalType":"Currency"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"tickSpacing","type":"int24","internalType":"int24"},{"name":"hooks","type":"address","internalType":"contract IHooks"}]},{"name":"params","type":"tuple","internalType":"struct IPoolManager.SwapParams","components":[{"name":"zeroForOne","type":"bool","internalType":"bool"},{"name":"amountSpecified","type":"int256","internalType":"int256"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]},{"name":"testSettings","type":"tuple","internalType":"struct PoolSwapTest.TestSettings","components":[{"name":"takeClaims","type":"bool","internalType":"bool"},{"name":"settleUsingBurn","type":"bool","internalType":"bool"}]},{"name":"hookData","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"delta","type":"int256","internalType":"BalanceDelta"}],"stateMutability":"payable"},{"type":"function","name":"unlockCallback","inputs":[{"name":"rawData","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"bytes","internalType":"bytes"}],"stateMutability":"nonpayable"},{"type":"error","name":"NoSwapOccurred","inputs":[]}]`

	config.SwapRouterABIJSON = `[{"type":"constructor","inputs":[{"name":"_manager","type":"address","internalType":"contract IPoolManager"}],"stateMutability":"nonpayable"},{"type":"function","name":"manager","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IPoolManager"}],"stateMutability":"view"},{"type":"function","name":"swap","inputs":[{"name":"key","type":"tuple","internalType":"struct PoolKey","components":[{"name":"currency0","type":"address","internalType":"Currency"},{"name":"currency1","type":"address","internalType":"Currency"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"tickSpacing","type":"int24","internalType":"int24"},{"name":"hooks","type":"address","internalType":"contract IHooks"}]},{"name":"params","type":"tuple","internalType":"struct IPoolManager.SwapParams","components":[{"name":"zeroForOne","type":"bool","internalType":"bool"},{"name":"amountSpecified","type":"int256","internalType":"int256"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]},{"name":"testSettings","type":"tuple","internalType":"struct PoolSwapTest.TestSettings","components":[{"name":"takeClaims","type":"bool","internalType":"bool"},{"name":"settleUsingBurn","type":"bool","internalType":"bool"}]},{"name":"hookData","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"delta","type":"int256","internalType":"BalanceDelta"}],"stateMutability":"payable"},{"type":"function","name":"unlockCallback","inputs":[{"name":"rawData","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"bytes","internalType":"bytes"}],"stateMutability":"nonpayable"},{"type":"error","name":"NoSwapOccurred","inputs":[]}]`

	config.PoolModifyLiquidityTestABIJSON = `[
		{
			"type": "function",
			"name": "modifyLiquidity",
			"inputs": [
			  {
				"name": "key",
				"type": "tuple",
				"internalType": "struct PoolKey",
				"components": [
				  {
					"name": "currency0",
					"type": "address",
					"internalType": "Currency"
				  },
				  {
					"name": "currency1",
					"type": "address",
					"internalType": "Currency"
				  },
				  { "name": "fee", "type": "uint24", "internalType": "uint24" },
				  { "name": "tickSpacing", "type": "int24", "internalType": "int24" },
				  {
					"name": "hooks",
					"type": "address",
					"internalType": "contract IHooks"
				  }
				]
			  },
			  {
				"name": "params",
				"type": "tuple",
				"internalType": "struct IPoolManager.ModifyLiquidityParams",
				"components": [
				  { "name": "tickLower", "type": "int24", "internalType": "int24" },
				  { "name": "tickUpper", "type": "int24", "internalType": "int24" },
				  {
					"name": "liquidityDelta",
					"type": "int256",
					"internalType": "int256"
				  },
				  { "name": "salt", "type": "bytes32", "internalType": "bytes32" }
				]
			  },
			  { "name": "hookData", "type": "bytes", "internalType": "bytes" }
			],
			"outputs": [
			  { "name": "delta", "type": "int256", "internalType": "BalanceDelta" }
			],
			"stateMutability": "payable"
		  }
	  ]`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

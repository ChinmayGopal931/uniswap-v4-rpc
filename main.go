// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/zap"
// )

// func main() {
// 	// Initialize logger
// 	logger, err := zap.NewProduction()
// 	if err != nil {
// 		log.Fatalf("Failed to initialize logger: %v", err)
// 	}
// 	defer logger.Sync()

// 	// Initialize Ethereum client
// 	if err := initEthereumClient(); err != nil {
// 		logger.Fatal("Failed to initialize Ethereum client", zap.Error(err))
// 	}

// 	// Set up Gin router
// 	router := gin.Default()

// 	// Add logger to Gin context
// 	router.Use(func(c *gin.Context) {
// 		c.Set("logger", logger)
// 		c.Next()
// 	})

// 	// Register routes
// 	router.POST("/swap", performSwap)

// 	// Start the server
// 	logger.Info("Starting server on :8080")
// 	if err := router.Run(":8080"); err != nil {
// 		logger.Fatal("Failed to start server", zap.Error(err))
// 	}
// }

package main

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ethClient                  *ethclient.Client
	chainID                    *big.Int
	swapRouterABI              abi.ABI
	poolModifyLiquidityTestABI abi.ABI
)

func initEthereumClient() error {
	var err error
	ethClient, err = ethclient.Dial(config.EthereumNodeURL)
	if err != nil {
		return err
	}

	// Set the EthClient in erc20.go
	EthClient = ethClient

	swapRouterABI, err = abi.JSON(strings.NewReader(config.SwapRouterABIJSON))
	if err != nil {
		return err
	}

	poolModifyLiquidityTestABI, err = abi.JSON(strings.NewReader(config.PoolModifyLiquidityTestABIJSON))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize Ethereum client and ABIs
	err = initEthereumClient()
	if err != nil {
		logger.Fatal("Failed to initialize Ethereum client", zap.Error(err))
	}

	// Get chain ID
	chainID, err = ethClient.ChainID(context.Background())
	if err != nil {
		logger.Fatal("Failed to get chain ID", zap.Error(err))
	}

	// Set up Gin router
	router := gin.Default()

	// Add logger to Gin context
	router.Use(func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	})

	// Register routes
	router.POST("/swap", performSwap)
	router.POST("/modify-liquidity", modifyLiquidity)

	// Start the server
	logger.Info("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

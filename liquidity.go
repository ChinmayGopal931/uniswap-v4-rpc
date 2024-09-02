package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	MinTick        int32 = -887272
	MaxTick        int32 = 887272
	MinTickSpacing int32 = 1
	MaxTickSpacing int32 = 32767 // max value for int16
)

type modifyLiquidityRequest struct {
	Currency0      string `json:"currency0"`
	Currency1      string `json:"currency1"`
	Fee            uint32 `json:"fee"`
	TickSpacing    int32  `json:"tickSpacing"`
	Hooks          string `json:"hooks"`
	LiquidityDelta string `json:"liquidityDelta"`
	Salt           string `json:"salt"`
	HookData       string `json:"hookData"`
}

func modifyLiquidity(c *gin.Context) {
	loggerInterface, exists := c.Get("logger")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	logger, ok := loggerInterface.(*zap.Logger)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid logger type"})
		return
	}

	var req modifyLiquidityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string inputs to appropriate types
	currency0 := common.HexToAddress(req.Currency0)
	currency1 := common.HexToAddress(req.Currency1)
	fee := new(big.Int).SetUint64(uint64(req.Fee))
	tickSpacing := new(big.Int).SetInt64(int64(req.TickSpacing))
	hooks := common.HexToAddress(req.Hooks)
	hookData := common.FromHex(req.HookData)

	poolKey := struct {
		Currency0   common.Address
		Currency1   common.Address
		Fee         *big.Int
		TickSpacing *big.Int
		Hooks       common.Address
	}{
		Currency0:   currency0,
		Currency1:   currency1,
		Fee:         fee,
		TickSpacing: tickSpacing,
		Hooks:       hooks,
	}

	liquidityDelta, ok := new(big.Int).SetString(req.LiquidityDelta, 10)
	if !ok {
		logger.Error("Invalid liquidityDelta")
		c.JSON(400, gin.H{"error": "Invalid liquidityDelta"})
		return
	}

	// Calculate min and max usable ticks
	minTick := (MinTick / req.TickSpacing) * req.TickSpacing
	maxTick := (MaxTick / req.TickSpacing) * req.TickSpacing

	params := struct {
		TickLower      *big.Int
		TickUpper      *big.Int
		LiquidityDelta *big.Int
		Salt           [32]byte
	}{
		TickLower:      new(big.Int).SetInt64(int64(minTick)),
		TickUpper:      new(big.Int).SetInt64(int64(maxTick)),
		LiquidityDelta: liquidityDelta,
		Salt:           common.HexToHash(req.Salt),
	}

	fmt.Printf("params: %+v\n", params)

	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, chainID)
	if err != nil {
		logger.Error("Failed to create transactor", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	poolModifyLiquidityTestABI, err := abi.JSON(strings.NewReader(config.PoolModifyLiquidityTestABIJSON))
	if err != nil {
		logger.Error("Failed to parse ABI", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	data, err := poolModifyLiquidityTestABI.Pack("modifyLiquidity", poolKey, params, hookData)
	if err != nil {
		logger.Error("Failed to pack data", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		logger.Error("Failed to get nonce", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("Failed to get gas price", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	tx := types.NewTransaction(nonce, config.PoolModifyLiquidityTestAddress, big.NewInt(0), 500000, gasPrice, data)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		logger.Error("Failed to sign transaction", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Error("Failed to send transaction", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info("Modify liquidity transaction sent", zap.String("txHash", signedTx.Hash().Hex()))
	c.JSON(200, gin.H{"txHash": signedTx.Hash().Hex()})
}

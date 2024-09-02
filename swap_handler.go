package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SwapRequest struct {
	Currency0         string `json:"currency0"`
	Currency1         string `json:"currency1"`
	Fee               string `json:"fee"`
	TickSpacing       string `json:"tickSpacing"`
	ZeroForOne        bool   `json:"zeroForOne"`
	AmountSpecified   string `json:"amountSpecified"`
	SqrtPriceLimitX96 string `json:"sqrtPriceLimitX96"`
	TakeClaims        bool   `json:"takeClaims"`
	SettleUsingBurn   bool   `json:"settleUsingBurn"`
	HookData          string `json:"hookData"`
}

func performSwap(c *gin.Context) {
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

	var req SwapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string inputs to appropriate types
	currency0 := common.HexToAddress(req.Currency0)
	currency1 := common.HexToAddress(req.Currency1)
	fee, _ := new(big.Int).SetString(req.Fee, 10)
	tickSpacing, _ := new(big.Int).SetString(req.TickSpacing, 10)
	amountSpecified, _ := new(big.Int).SetString(req.AmountSpecified, 10)
	sqrtPriceLimitX96, _ := new(big.Int).SetString(req.SqrtPriceLimitX96, 10)
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
		Hooks:       config.HooksAddress,
	}

	swapParams := struct {
		ZeroForOne        bool
		AmountSpecified   *big.Int
		SqrtPriceLimitX96 *big.Int
	}{
		ZeroForOne:        req.ZeroForOne,
		AmountSpecified:   amountSpecified,
		SqrtPriceLimitX96: sqrtPriceLimitX96,
	}

	testSettings := struct {
		TakeClaims      bool
		SettleUsingBurn bool
	}{
		TakeClaims:      req.TakeClaims,
		SettleUsingBurn: req.SettleUsingBurn,
	}

	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, big.NewInt(31337)) // Anvil chain ID
	if err != nil {
		logger.Error("Failed to create transactor", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
		return
	}

	// Check balances before swap
	balancesBefore, err := checkBalances(currency0, currency1, auth.From, logger)
	if err != nil {
		logger.Error("Failed to check balances before swap", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check initial balances"})
		return
	}

	logger.Info("Balances before swap",
		zap.String("currency0", balancesBefore.Currency0.String()),
		zap.String("currency1", balancesBefore.Currency1.String()),
	)

	data, err := swapRouterABI.Pack("swap", poolKey, swapParams, testSettings, hookData)
	if err != nil {
		logger.Error("Failed to pack data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to pack data"})
		return
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		logger.Error("Failed to get nonce", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nonce"})
		return
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("Failed to get gas price", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gas price"})
		return
	}

	tx := types.NewTransaction(nonce, config.SwapRouterAddress, big.NewInt(0), 500000, gasPrice, data)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		logger.Error("Failed to sign transaction", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign transaction"})
		return
	}

	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Error("Failed to send transaction", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send transaction"})
		return
	}

	logger.Info("Transaction sent", zap.String("txHash", signedTx.Hash().Hex()))

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), ethClient, signedTx)
	if err != nil {
		logger.Error("Failed to wait for transaction to be mined", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to wait for transaction to be mined"})
		return
	}

	if receipt.Status == 0 {
		logger.Error("Transaction failed", zap.String("txHash", signedTx.Hash().Hex()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}

	// Check balances after swap
	balancesAfter, err := checkBalances(currency0, currency1, auth.From, logger)
	if err != nil {
		logger.Error("Failed to check balances after swap", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check final balances"})
		return
	}

	// Calculate balance changes
	delta0 := new(big.Int).Sub(balancesAfter.Currency0, balancesBefore.Currency0)
	delta1 := new(big.Int).Sub(balancesAfter.Currency1, balancesBefore.Currency1)

	logger.Info("Swap completed",
		zap.String("txHash", signedTx.Hash().Hex()),
		zap.String("currency0_before", balancesBefore.Currency0.String()),
		zap.String("currency1_before", balancesBefore.Currency1.String()),
		zap.String("currency0_after", balancesAfter.Currency0.String()),
		zap.String("currency1_after", balancesAfter.Currency1.String()),
		zap.String("currency0_delta", delta0.String()),
		zap.String("currency1_delta", delta1.String()),
	)

	c.JSON(http.StatusOK, gin.H{
		"txHash":         signedTx.Hash().Hex(),
		"balancesBefore": gin.H{"currency0": balancesBefore.Currency0.String(), "currency1": balancesBefore.Currency1.String()},
		"balancesAfter":  gin.H{"currency0": balancesAfter.Currency0.String(), "currency1": balancesAfter.Currency1.String()},
		"deltaBalances":  gin.H{"currency0": delta0.String(), "currency1": delta1.String()},
	})
}

type Balances struct {
	Currency0 *big.Int
	Currency1 *big.Int
}

func checkBalances(currency0, currency1, owner common.Address, logger *zap.Logger) (*Balances, error) {
	balance0, err := getBalance(currency0, owner)
	if err != nil {
		logger.Error("Failed to get balance of currency0", zap.Error(err))
		return nil, fmt.Errorf("failed to get balance of currency0: %v", err)
	}

	balance1, err := getBalance(currency1, owner)
	if err != nil {
		logger.Error("Failed to get balance of currency1", zap.Error(err))
		return nil, fmt.Errorf("failed to get balance of currency1: %v", err)
	}

	return &Balances{
		Currency0: balance0,
		Currency1: balance1,
	}, nil
}

// Helper function to get balance of a token for an address
func getBalance(tokenAddress, ownerAddress common.Address) (*big.Int, error) {
	token, err := newERC20(tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token contract: %v", err)
	}

	balance, err := token.BalanceOf(&bind.CallOpts{}, ownerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return balance, nil
}

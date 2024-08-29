// rpc/server.go
package rpc

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"uniswap-v4-rpc/contracts/csmm"
	"uniswap-v4-rpc/contracts/erc20"
)

type Server struct {
	ethClient  *ethclient.Client
	csmm       *csmm.Csmm
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
}

func NewServer(ethClient *ethclient.Client, csmmAddress common.Address, privateKeyHex string) (*Server, error) {
	csmmContract, err := csmm.NewCsmm(csmmAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate CSMM contract: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	return &Server{
		ethClient:  ethClient,
		csmm:       csmmContract,
		privateKey: privateKey,
		chainID:    chainID,
	}, nil
}

type AddLiquidityRequest struct {
	Key struct {
		Currency0    common.Address `json:"currency0"`
		Currency1    common.Address `json:"currency1"`
		Fee          string         `json:"fee"`
		TickSpacing  int32          `json:"tickSpacing"`
		HooksAddress common.Address `json:"hooksAddress"`
	} `json:"key"`
	AmountEach string `json:"amountEach"`
}

type AddLiquidityResponse struct {
	TransactionHash string `json:"transactionHash"`
}

//
func (s *Server) AddLiquidity(params json.RawMessage) (string, error) {
	var p struct {
		Key struct {
			Currency0   common.Address `json:"currency0"`
			Currency1   common.Address `json:"currency1"`
			Fee         string         `json:"fee"`
			TickSpacing int32          `json:"tickSpacing"`
			Hooks       common.Address `json:"hooks"`
		} `json:"key"`
		AmountEach string `json:"amountEach"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return "", fmt.Errorf("failed to unmarshal params: %v", err)
	}

	amountEach, ok := new(big.Int).SetString(p.AmountEach, 10)
	if !ok {
		return "", fmt.Errorf("invalid amountEach")
	}

	fee, ok := new(big.Int).SetString(p.Key.Fee, 10)
	if !ok {
		return "", fmt.Errorf("invalid fee")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	// Set gas limit and gas price
	auth.GasLimit = uint64(3000000) // Adjust as needed
	gasPrice, err := s.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice

	key := csmm.PoolKey{
		Currency0:   p.Key.Currency0,
		Currency1:   p.Key.Currency1,
		Fee:         fee,
		TickSpacing: big.NewInt(int64(p.Key.TickSpacing)),
		Hooks:       p.Key.Hooks,
	}

	tx, err := s.csmm.AddLiquidity(auth, key, amountEach)
	if err != nil {
		return "", fmt.Errorf("failed to add liquidity: %v", err)
	}

	return tx.Hash().Hex(), nil
}




func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		JsonRpc string          `json:"jsonrpc"`
		Method  string          `json:"method"`
		Params  json.RawMessage `json:"params"`
		Id      int             `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result interface{}
	var err error

	switch req.Method {
	case "addLiquidity":
		result, err = s.AddLiquidity(req.Params)
	case "approveToken":
		result, err = s.ApproveToken(req.Params)
	// Add other cases as needed
	default:
		err = fmt.Errorf("method not found")
	}

	response := struct {
		JsonRpc string      `json:"jsonrpc"`
		Result  interface{} `json:"result,omitempty"`
		Error   string      `json:"error,omitempty"`
		Id      int         `json:"id"`
	}{
		JsonRpc: "2.0",
		Id:      req.Id,
	}

	if err != nil {
		response.Error = err.Error()
	} else {
		response.Result = result
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) ApproveToken(params json.RawMessage) (string, error) {
	var p struct {
		Token   common.Address `json:"token"`
		Spender common.Address `json:"spender"`
		Amount  *big.Int       `json:"amount"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	token, err := erc20.NewErc20(p.Token, s.ethClient)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate token contract: %v", err)
	}

	tx, err := token.Approve(auth, p.Spender, p.Amount)
	if err != nil {
		return "", fmt.Errorf("failed to approve token: %v", err)
	}

	return tx.Hash().Hex(), nil
}
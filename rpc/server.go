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

	"uniswap-v4-rpc/contracts"
)


type Server struct {
	ethClient  *ethclient.Client
	contracts  *contracts.Contracts
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
}

func NewServer(ethClient *ethclient.Client, contractAddress common.Address, privateKeyHex string) (*Server, error) {
	contractsInstance, err := contracts.NewContracts(contractAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Contracts: %v", err)
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
		contracts:  contractsInstance,
		privateKey: privateKey,
		chainID:    chainID,
	}, nil
}

type AddLiquidityRequest struct {
	Key struct {
		Currency0    common.Address `json:"currency0"`
		Currency1    common.Address `json:"currency1"`
		Fee          *big.Int       `json:"fee"`
		TickSpacing  int32          `json:"tickSpacing"`
		HooksAddress common.Address `json:"hooksAddress"`
	} `json:"key"`
	AmountEach string `json:"amountEach"`
}

type AddLiquidityResponse struct {
	TransactionHash string `json:"transactionHash"`
}

//
func (s *Server) AddLiquidity(r *http.Request, req *AddLiquidityRequest, res *AddLiquidityResponse) error {
	amountEach, ok := new(big.Int).SetString(req.AmountEach, 10)
	if !ok {
		return fmt.Errorf("invalid amountEach")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	key := contracts.PoolKey{
		Currency0:   req.Key.Currency0,
		Currency1:   req.Key.Currency1,
		Fee:         req.Key.Fee,
		TickSpacing: big.NewInt(int64(req.Key.TickSpacing)),
		Hooks:       req.Key.HooksAddress,
	}

	tx, err := s.contracts.AddLiquidity(auth, key, amountEach)
	if err != nil {
		return fmt.Errorf("failed to add liquidity: %v", err)
	}

	res.TransactionHash = tx.Hash().Hex()
	return nil
}



func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

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
		var params AddLiquidityRequest
		if err := json.Unmarshal(req.Params, &params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var res AddLiquidityResponse
		err = s.AddLiquidity(r, &params, &res)
		result = res
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
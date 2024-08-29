package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"uniswap-v4-rpc/rpc"
	"uniswap-v4-rpc/contracts/erc20"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Ethereum client
	ethClient, err := ethclient.Dial("http://localhost:8545") // Replace with your Ethereum node URL
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Fetch and log the chain ID
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}
	log.Printf("Connected to chain ID: %d", chainID)

	// Address of your deployed Contracts contract
	//contractAddress := common.HexToAddress("0x4F30fc8417172821c004a571Cb70085EA3c3C888") // Replace with your contract address

	// Get private key from environment variable
	// privateKeyHex := os.Getenv("ETHEREUM_PRIVATE_KEY")
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	log.Printf("Connected privateKeyHex: %d", privateKeyHex)

	if privateKeyHex == "" {
		log.Fatal("ETHEREUM_PRIVATE_KEY environment variable is not set")
	}

	// Convert private key from hex to ECDSA
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to get public key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("Connected Address: %s", address.Hex())

	// Check balance
	balance, err := ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	log.Printf("Account balance: %s wei", balance.String())

	token0Address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	token1Address := common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
	csmmAddress := common.HexToAddress("0x4F30fc8417172821c004a571Cb70085EA3c3C888")

	err = increaseAllowance(ethClient, privateKey, token0Address, csmmAddress, big.NewInt(1e18)) // Approve 1 token
	if err != nil {
		log.Printf("Failed to increase allowance for Token0: %v", err)
	}

	err = increaseAllowance(ethClient, privateKey, token1Address, csmmAddress, big.NewInt(1e18)) // Approve 1 token
	if err != nil {
		log.Printf("Failed to increase allowance for Token1: %v", err)
	}

	
	// Check allowances
	allowance0, err := checkAllowance(ethClient, token0Address, address, csmmAddress)
	if err != nil {
		log.Printf("Failed to check allowance for Token0: %v", err)
	} else {
		log.Printf("Allowance for Token0: %s", allowance0.String())
	}

	allowance1, err := checkAllowance(ethClient, token1Address, address, csmmAddress)
	if err != nil {
		log.Printf("Failed to check allowance for Token1: %v", err)
	} else {
		log.Printf("Allowance for Token1: %s", allowance1.String())
	}

	// Create RPC server
	rpcServer, err := rpc.NewServer(ethClient, csmmAddress, privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to create RPC server: %v", err)
	}

	// Start HTTP server
	http.HandleFunc("/", rpcServer.ServeHTTP)
	log.Println("Starting JSON-RPC server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkAllowance(client *ethclient.Client, tokenAddress, ownerAddress, spenderAddress common.Address) (*big.Int, error) {
	token, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token contract: %v", err)
	}

	allowance, err := token.Allowance(&bind.CallOpts{}, ownerAddress, spenderAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check allowance: %v", err)
	}

	return allowance, nil
}

func increaseAllowance(client *ethclient.Client, privateKey *ecdsa.PrivateKey, tokenAddress, spenderAddress common.Address, amount *big.Int) error {
	token, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate token contract: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	tx, err := token.Approve(auth, spenderAddress, amount)
	if err != nil {
		return fmt.Errorf("failed to approve token: %v", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	log.Printf("Increased allowance for token %s, transaction hash: %s", tokenAddress.Hex(), tx.Hash().Hex())
	return nil
}
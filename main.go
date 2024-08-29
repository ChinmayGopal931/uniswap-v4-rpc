package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"uniswap-v4-rpc/rpc"
)

func main() {
	// Connect to Ethereum client
	ethClient, err := ethclient.Dial("http://localhost:8545") // Replace with your Ethereum node URL
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Address of your deployed CSMM contract
	csmmAddress := common.HexToAddress("0x1234...") // Replace with your contract address

	// Get private key from environment variable
	privateKey := os.Getenv("ETHEREUM_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("ETHEREUM_PRIVATE_KEY environment variable is not set")
	}

	// Create RPC server
	rpcServer, err := rpc.NewServer(ethClient, csmmAddress, privateKey)
	if err != nil {
		log.Fatalf("Failed to create RPC server: %v", err)
	}

	// Start HTTP server
	http.HandleFunc("/", rpcServer.ServeHTTP)
	log.Println("Starting JSON-RPC server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
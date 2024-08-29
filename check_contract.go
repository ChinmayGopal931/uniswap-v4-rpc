package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress("0x4F30fc8417172821c004a571Cb70085EA3c3C888")
    bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
    if err != nil {
        log.Fatal(err)
    }

    if len(bytecode) > 0 {
        fmt.Printf("Contract is deployed at %s\n", address.Hex())
    } else {
        fmt.Printf("No contract found at %s\n", address.Hex())
    }
}
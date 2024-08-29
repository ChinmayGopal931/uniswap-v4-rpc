// keystore/keystore.go
package keystore

import (
	"crypto/ecdsa"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Keystore struct {
	keys map[common.Address]*ecdsa.PrivateKey
	mu   sync.RWMutex
}

func NewKeystore() *Keystore {
	return &Keystore{
		keys: make(map[common.Address]*ecdsa.PrivateKey),
	}
}

func (ks *Keystore) NewAccount() (common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	ks.mu.Lock()
	ks.keys[address] = privateKey
	ks.mu.Unlock()

	return address, nil
}

func (ks *Keystore) GetPrivateKey(address common.Address) (*ecdsa.PrivateKey, error) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	privateKey, ok := ks.keys[address]
	if !ok {
		return nil, errors.New("address not found in keystore")
	}

	return privateKey, nil
}
package bitcoin

import (
	"fmt"
	"goApiStartetProject/internal/util/wallet"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

func BTCAddress(seedphrase string) (*btcutil.AddressPubKeyHash, error) {
	// 1. Generate a random 128-bit seed
	seed, _ := wallet.HashSeedPhrase(seedphrase)
	// 2. Create a master key from the seed
	masterKey, err := hdkeychain.NewMaster(seed, nil)
	if err != nil {
		return &btcutil.AddressPubKeyHash{}, fmt.Errorf("Failed to create master key: %v", err)
	}

	// 3. Derive a child key from the master key
	childKey, err := masterKey.Child(0) // Derive the first child key
	if err != nil {
		return &btcutil.AddressPubKeyHash{}, fmt.Errorf("Failed to derive child key: %v", err)
	}

	// 4. Get the address and private key corresponding to the child key
	address, err := childKey.Address(&chaincfg.SimNetParams)
	if err != nil {
		return &btcutil.AddressPubKeyHash{}, fmt.Errorf("Failed to get address: %v", err)
	}
	privateKey, err := childKey.ECPrivKey()
	if err != nil {
		return &btcutil.AddressPubKeyHash{}, fmt.Errorf("Failed to get private key: %v", err)
	}

	// 5. Print the address and private key
	fmt.Println("Address:", address)
	fmt.Println("Private Key:", privateKey)

	return address, nil
}

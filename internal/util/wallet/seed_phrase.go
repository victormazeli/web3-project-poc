package wallet

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

func SeedPhrase() string {
	// Generate a random 12-word mnemonic
	entropy, _ := bip39.NewEntropy(128) // 128 bits produces 12 words
	mnemonic, _ := bip39.NewMnemonic(entropy)

	fmt.Println("Generated 12-word mnemonic:")
	fmt.Println(mnemonic)

	return mnemonic
}

func HashSeedPhrase(mnemonic string) ([]byte, error) {
	// To get the seed from the mnemonic
	hashedSeed := bip39.NewSeed(mnemonic, "") // empty passphrase
	fmt.Printf("Seed: %x\n", hashedSeed)

	return hashedSeed, nil
}

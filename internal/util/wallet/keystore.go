package wallet

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewKeystoreAccount generates a new Ethereum account using a keystore.
//
// It takes a password as a parameter and returns the address of the
// newly generated account and any error that occurred during the process.
func NewKeystoreAccount(password string, client *ethclient.Client) (*common.Address, error) {
	// Set the directory for the keystore.
	keystoreDir := "./wallets"
	
	// Set the scrypt parameters for the keystore.
	scryptN, scryptP := keystore.StandardScryptN, keystore.StandardScryptP
	
	// Initialize the keystore using the directory and scrypt parameters.
	ks := keystore.NewKeyStore(keystoreDir, scryptN, scryptP)
	

	// Generate a new account using the provided password.
	account, err := ks.NewAccount(password)
	
	// If an error occurred, return it. Otherwise, return the address of the new account.
	if err != nil {
		return nil, err
	}
	
	return &account.Address, nil
}

func ImportKeystore(filePath, keystoreDir, password string) (string, error) {
    keystore := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
    jsonBytes, err := os.ReadFile(filePath)
    if err != nil {
        return "", err
    }

    account, err := keystore.Import(jsonBytes, password, password)
    if err != nil {
        return "", err
    }

    if err := os.Remove(filePath); err != nil {
        return "", err
    }

    return account.Address.Hex(), nil
}
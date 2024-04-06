package wallet

import (
	"fmt"
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

func ImportKeystore(filePath, keystoreDir, password string) (common.Address, error) {
	
	ks := keystore.NewKeyStore("/tmps", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(keystoreDir+filePath)
    if err != nil {
        return common.Address{}, err
    }

    account, err := ks.Import(jsonBytes, password, "passwored")
	fmt.Println(err)
    if err != nil {
        return common.Address{}, err
    }

    if err := os.Remove(keystoreDir+filePath); err != nil {
        return common.Address{}, err
    }

	fmt.Println(account)
    return account.Address, nil
}
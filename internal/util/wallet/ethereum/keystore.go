package wallet

import (
	"crypto/ecdsa"
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

func ImportAccountFromKeystore(passphrase string) (common.Address, error) {
    // Directory for the keystore.
	keystoreDir := "./wallets"
    
    // Open the keystore directory
    ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

    // Get list of accounts
    accounts := ks.Accounts()

    account := common.Address{}

    if len(accounts) == 0 {
		return common.Address{}, fmt.Errorf("no accounts found in keystore")
	}

    // Iterate through accounts
    for _, acc := range accounts {
        // Unlock account using passphrase
        err := ks.Unlock(acc, passphrase)
        if err == nil {
            account = acc.Address
        } else if err == keystore.ErrNoMatch {
            return common.Address{}, fmt.Errorf("Account not found: %s", acc.Address.Hex())
        } else if err == keystore.ErrDecrypt {
            return common.Address{}, fmt.Errorf("Incorrect passphrase for account: %s", acc.Address.Hex())
        } else {
            return common.Address{}, fmt.Errorf("Error unlocking account: %v", err)
        }
    }

    return account, nil
}

func GetPrivateKeyFromKeystore(keystorePath string, password string, targetAddress common.Address) (*ecdsa.PrivateKey, error) {
    // Read the keystore file
    keyjson, err := os.ReadFile(keystorePath)
    if err != nil {
        return &ecdsa.PrivateKey{}, err
    }

    // Decrypt the keystore file
    key, err := keystore.DecryptKey(keyjson, password)
    if err != nil {
        return &ecdsa.PrivateKey{}, err
    }

    // Derive the public address from the private key
    address := key.Address

    // Compare with the target address
    if address != targetAddress {
        return &ecdsa.PrivateKey{}, fmt.Errorf("keystore does not correspond to the provided address")
    }

    // Get the private key
    privateKey := key.PrivateKey
	fmt.Println(key)


    // Convert private key to hex string
    // privateKeyHex := fmt.Sprintf("%x", privateKey.D.Bytes())

    return privateKey, nil
}
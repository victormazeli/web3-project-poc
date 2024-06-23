package validator

import (
	"context"
	// "fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func GenerateAddress() common.Address {
// 	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
// 	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B
// 	// fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c76
// 	fmt.Println(address.Bytes())

// 	return address
// }

func CheckIsValidAddress(ctx context.Context, client *ethclient.Client, address string) (bool, error){
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	isValidAddr := re.MatchString(address)

	if !isValidAddr {
		return false, nil
	}

	return CheckIsNotContractAddress(ctx, address, client)

}

func CheckIsNotContractAddress(ctx context.Context, addr string, client *ethclient.Client) (bool, error) {
	// a random user account address
    address := common.HexToAddress(addr)
    bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
    if err != nil {
        return false, err
    }

    isContract := len(bytecode) > 0

    // fmt.Printf("is contract: %v\n", isContract) // is contract: false

	if isContract {
		return false, nil
	}

	return true, nil
}

package wallet

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransactionHistory implements TransactionServiceInterface.
func TransactionHistory(ctx context.Context, block *types.Block) types.Transactions {

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
	}

	return block.Transactions()
}

func GetEthWalletBalance(ctx context.Context, client *ethclient.Client, address common.Address) (*big.Float, error){
	// Get balance of the Ethereum address
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(balance)


	// Convert balance from wei to ether
	etherBalance := new(big.Float).SetInt(balance)
	etherBalance = etherBalance.Quo(etherBalance, big.NewFloat(1e18))


	return etherBalance, nil
}

func FetchSenderAddress(client *ethclient.Client, tx *types.Transaction) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if msg, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
		fmt.Println(msg.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
	}
}

func FetchTxReciept(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, err
	}

	fmt.Println(receipt.Status) // 1
	fmt.Println(receipt.Logs)

	return receipt, nil
}



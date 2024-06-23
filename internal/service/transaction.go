package service

import (
	"context"
	"fmt"
	"goApiStartetProject/internal/domain"
	"goApiStartetProject/internal/storages/postgres/repository"
	"goApiStartetProject/internal/util/validator/ethereum"
	"goApiStartetProject/internal/util/wallet/ethereum"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
)

type TransactionServiceInterface interface {

	TransferCoin(ctx context.Context, client *ethclient.Client, req domain.TransferCoinRequestPayload) (common.Hash, error)

	TransactionHistory(ctx context.Context, block *types.Block) types.Transactions

	GetBlockInfo(ctx context.Context, client *ethclient.Client) (*types.Block, error)
}

type TransactionService struct {
	repo *repository.Repository
}

// TransferCoin implements TransactionServiceInterface.
func (t *TransactionService) TransferCoin(ctx context.Context, client *ethclient.Client, txReq domain.TransferCoinRequestPayload) (common.Hash, error) {
	// tx := types.NewTransaction(txReq.Nonce, txReq.ToAddress, txReq.Amount, txReq.GasLimit, txReq.GasPrice, nil)
	// Set the directory for the keystore.
	// keystoreDir := ""

	filePath := "./wallets/UTC--2024-03-22T11-22-29.712456800Z--6713fe6cb976e6787ac0dbd2c26efcbb2e41d6e0"
	isValid, err := validator.CheckIsValidAddress(ctx, client, txReq.FromAddress.String())
	if !isValid {
		return common.Hash{}, fmt.Errorf("sender address %v is not valid: err %v", txReq.FromAddress, err)
	}
	isValid, err = validator.CheckIsValidAddress(ctx, client, txReq.ToAddress.String())
	if !isValid {
		return common.Hash{}, fmt.Errorf("recipient address %v is not valid: err %v", txReq.ToAddress, err)
	}
	balance, err := wallet.GetEthWalletBalance(ctx, client, txReq.FromAddress)
	if err != nil{
		return common.Hash{}, fmt.Errorf("error getting wallet balance")
	}
	fmt.Println(balance)

	intAmount := big.NewInt(1234567890)

	floatAmount := new(big.Float).SetInt(intAmount)
	isValidBalance := validator.CanTransferAmount(floatAmount, balance)
	if !isValidBalance {
		return common.Hash{}, fmt.Errorf("insufficient wallet balance")
	}

	privateKey, err := wallet.GetPrivateKeyFromKeystore(filePath, txReq.Password, txReq.FromAddress)
	if err != nil{
		fmt.Println(err)
		return common.Hash{}, err
	}

	// txReq.FromAddress = publicAddress
	// fmt.Println(publicAddress)
	// privateKey, err := crypto.HexToECDSA(publicAddress.Hex())
    // if err != nil {
    //     log.Fatal(err)
    // }

    // publicKey := privateKey.Public()
    // publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    // if !ok {
    //     log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    // }

    // fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	txReq.SetTransferCoinReqPayload(ctx, client)

	// toAddress := common.HexToAddress(publicAddress.Hex())
    var data []byte// Convert big.Float to big.Int
	
    tx := types.NewTransaction(txReq.Nonce, txReq.ToAddress, txReq.Amount, txReq.GasLimit, txReq.GasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
	log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
		return common.Hash{}, err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	return signedTx.Hash(), nil
}

// TransactionHistory implements TransactionServiceInterface.
func (t *TransactionService) TransactionHistory(ctx context.Context, block *types.Block) types.Transactions {
	panic("unimplemented")
}

// GetBlockInfo implements TransactionServiceInterface.
func (t *TransactionService) GetBlockInfo(ctx context.Context, client *ethclient.Client) (*types.Block, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		return nil, err
	}

	fmt.Println(header.Number.String())

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

	return block, nil
}

func NewTransactionService(db *sqlx.DB) TransactionServiceInterface {
	return &TransactionService{
		repo: repository.NewRepository(db),
	}
}

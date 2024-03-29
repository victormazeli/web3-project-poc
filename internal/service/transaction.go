package service

import (
	"context"
	"fmt"
	"goApiStartetProject/core/storages/postgres/repository"
	"goApiStartetProject/internal/domain"
	"goApiStartetProject/internal/util/wallet"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
)

type TransactionServiceInterface interface {
	SellCoin(ctx context.Context, id uint64) error

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
	keystoreDir := "./wallets"

	filePath := "UTC--2024-03-16T02-37-58.426400500Z--0b04d804275162b0a200d303271b1f2d03464ac3"
	
	publicAddress, err := wallet.ImportKeystore(filePath, keystoreDir, txReq.Password)
	if err != nil{
		return common.Hash{}, err
	}

	privateKey, err := crypto.HexToECDSA(publicAddress)
    if err != nil {
        log.Fatal(err)
    }

    // publicKey := privateKey.Public()
    // publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    // if !ok {
    //     log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    // }

    // fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	txReq.SetTransferCoinReqPayload(ctx, client)

	toAddress := common.HexToAddress(publicAddress)
    var data []byte
    tx := types.NewTransaction(txReq.Nonce, toAddress, txReq.Amount, txReq.GasLimit, txReq.GasPrice, data)

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

// SellCoin implements TransactionServiceInterface.
func (t *TransactionService) SellCoin(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

func NewTransactionService(db *sqlx.DB) TransactionServiceInterface {
	return &TransactionService{
		repo: repository.NewRepository(db),
	}
}

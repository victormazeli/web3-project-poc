package service

import (
	"context"
	"goApiStartetProject/internal/storages/postgres/repository"
	"goApiStartetProject/internal/util/wallet"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// "github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"fmt"
)

type AccountServiceInterface interface {
	GenerateAccountAddress(ctx context.Context, ethClient *ethclient.Client) (common.Address, error)
}

type AccountService struct {
	repo *repository.Repository
}

func NewAccountService(db *sqlx.DB) AccountServiceInterface {
	return &AccountService{
		repo: repository.NewRepository(db),
	}
}

func (w *AccountService) GenerateAccountAddress(ctx context.Context, ethClient *ethclient.Client) (common.Address, error) {
	
	address := wallet.GenerateAddress()

	balance, err := ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
	log.Fatal(err)
	}
	fmt.Println(balance)

	fmt.Println(address)

	return address, nil
}

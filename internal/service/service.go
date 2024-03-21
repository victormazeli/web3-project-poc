package service

import (
	"fmt"
	"goApiStartetProject/core/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	Env                *config.Env
	EthClient          *ethclient.Client
	UserService        UserServiceInterface
	WalletService     WalletServiceInterface
	TransactionService TransactionServiceInterface
}

func NewService(env *config.Env, ethClient *ethclient.Client, db *sqlx.DB) *Service {
	fmt.Printf("true %v", ethClient)
	return &Service{
		Env:         env,
		EthClient:   ethClient,
		UserService: NewUserService(db),
		WalletService: NewWalletService(db),
	}

}

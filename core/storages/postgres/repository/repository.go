package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	UserRepo        UserRepositoryInterface
	WalletRepo      WalletRepositoryInterface
	AccountRepo     AccountRepositoryInterface
	TransactionRepo TransactionRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:        NewUserRepository(db),
		WalletRepo:      NewWalletRepository(db),
		TransactionRepo: NewTransactionRepository(db),
		AccountRepo:     NewAccountRepository(db),
	}
}

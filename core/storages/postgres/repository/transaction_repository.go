package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	ID          uuid.UUID `db:"id"`
	Hash        string    `db:"hash"`
	Status      string    `db:"status"`
	CoinType    string    `db:"coin_type"`
	Amount      string    `db:"amount"`
	ToAddress   string    `db:"to_address"`
	FromAddress string    `db:"from_address"`
	GasPrice    string    `db:"gas_price"`
	GasLimit    string    `db:"gas_limit"`
	Nounce      string    `db:"nounce"`
	Data        string    `db:"data"`
	WalletID   uuid.UUID `db:"wallet_id"`
}

type TransactionRepositoryInterface interface {
	CreateTransaction(ctx context.Context, Transaction *Transaction) (uuid.UUID, error)
	GetTransactionByID(ctx context.Context, id int) (*Transaction, error)
	UpdateTransaction(ctx context.Context, Transaction *Transaction) error
	DeleteTransaction(ctx context.Context, id int) error
}

type TransactionRepository struct {
	db *sqlx.DB
}

// CreateTransaction implements TransactionRepositoryInterface.
func (t *TransactionRepository) CreateTransaction(ctx context.Context, Transaction *Transaction) (uuid.UUID, error) {
	panic("unimplemented")
}

// DeleteTransaction implements TransactionRepositoryInterface.
func (t *TransactionRepository) DeleteTransaction(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetTransactionByID implements TransactionRepositoryInterface.
func (t *TransactionRepository) GetTransactionByID(ctx context.Context, id int) (*Transaction, error) {
	panic("unimplemented")
}

// UpdateTransaction implements TransactionRepositoryInterface.
func (t *TransactionRepository) UpdateTransaction(ctx context.Context, Transaction *Transaction) error {
	panic("unimplemented")
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepositoryInterface {
	return &TransactionRepository{db: db}
}

package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Account struct {
	ID         uuid.UUID   `db:"id"`
	Balance    string      `db:"balance"`
	PrivateKey string      `db:"private_key"`
	PublicKey  string      `db:"public_key"`
	WalletID   uuid.UUID   `db:"wallet_id"`
	CreatedAt  time.Time   `db:"created_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
}


type AccountRepositoryInterface interface {
	CreateAccount(ctx context.Context, Account *Account) (uuid.UUID, error)
	GetAccountByID(ctx context.Context, id int) (*Account, error)
	UpdateAccount(ctx context.Context, Account *Account) error
	DeleteAccount(ctx context.Context, id int) error
}

type AccountRepository struct {
	db *sqlx.DB
}

// CreateAccount implements AccountRepositoryInterface.
func (a *AccountRepository) CreateAccount(ctx context.Context, Account *Account) (uuid.UUID, error) {
	panic("unimplemented")
}

// DeleteAccount implements AccountRepositoryInterface.
func (a *AccountRepository) DeleteAccount(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetAccountByID implements AccountRepositoryInterface.
func (a *AccountRepository) GetAccountByID(ctx context.Context, id int) (*Account, error) {
	panic("unimplemented")
}

// UpdateAccount implements AccountRepositoryInterface.
func (a *AccountRepository) UpdateAccount(ctx context.Context, Account *Account) error {
	panic("unimplemented")
}

func NewAccountRepository(db *sqlx.DB) AccountRepositoryInterface {
	return &AccountRepository{db: db}
}

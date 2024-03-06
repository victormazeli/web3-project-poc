package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type WalletRepositoryInterface interface {
	CreateWallet(ctx context.Context, wallet *Wallet) (int, error)
	GetWalletByID(ctx context.Context, id int) (*Wallet, error)
	UpdateWallet(ctx context.Context, wallet *Wallet) error
	DeleteWallet(ctx context.Context, id int) error
}

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) WalletRepositoryInterface {
	return &WalletRepository{db: db}
}

// CreateWallet implements WalletRepositoryInterface.
func (w *WalletRepository) CreateWallet(ctx context.Context, wallet *Wallet) (int, error) {
	panic("unimplemented")
}

// DeleteWallet implements WalletRepositoryInterface.
func (w *WalletRepository) DeleteWallet(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetWalletByID implements WalletRepositoryInterface.
func (w *WalletRepository) GetWalletByID(ctx context.Context, id int) (*Wallet, error) {
	panic("unimplemented")
}

// UpdateWallet implements WalletRepositoryInterface.
func (w *WalletRepository) UpdateWallet(ctx context.Context, wallet *Wallet) error {
	panic("unimplemented")
}



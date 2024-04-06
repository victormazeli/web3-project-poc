package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Balance    string    `db:"balance"`
	PrivateKey string    `db:"private_key"`
	PublicKey  string    `db:"public_key"`
	Type 		string		`db:"wallet_type"`
	Coin       CoinDetails `db:"coin_details"`
	UserID     uuid.UUID `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}



type CoinDetails struct {
	ChainName string `db:"name"`
	Currency  string `db:"currency"`
	Logo      string `db:"logo"`
}

type WalletRepositoryInterface interface {
	CreateWallet(ctx context.Context, wallet *Wallet) (uuid.UUID, error)
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
func (w *WalletRepository) CreateWallet(ctx context.Context, wallet *Wallet) (uuid.UUID, error) {
	query := `INSERT INTO users (id, name, type, balance, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	row := w.db.QueryRowContext(
		ctx,
		query,
		wallet.ID,
		wallet.Name,
		wallet.Balance,
		time.Now(),
		time.Now())
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to save user details: %w", err)
	}
	return id, nil
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

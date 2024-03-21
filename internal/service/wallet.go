package service

import (
	"context"
	// "crypto/ecdsa"
	// "errors"
	// "fmt"
	"goApiStartetProject/core/storages/postgres/repository"
	"goApiStartetProject/internal/domain"
	walletPkg "goApiStartetProject/internal/util/wallet"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"

	// "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// walletService is an interface for interacting with wallet-related business logic
type WalletServiceInterface interface {
	// Register registers a Create wallet
	GenerateWalletAddress(ctx context.Context, ethClient *ethclient.Client, wallet domain.CreateWalletRequestPayload) (string, error)
	// Getwallet returns a wallet by id
	Getwallet(ctx context.Context, id uint64) (*repository.Wallet, error)
	// Listwallets returns a list of wallets with pagination - admin privileges
	Listwallets(ctx context.Context, skip, limit uint64) ([]repository.Wallet, error)
	// Updatewallet updates a wallet
	Updatewallet(ctx context.Context, wallet repository.Wallet) (*repository.Wallet, error)
	// Deletewallet deletes a wallet
	Deletewallet(ctx context.Context, id uint64) error

	// Transaction related functions
	Deposit(ctx context.Context, id uuid.UUID) *repository.Wallet

	Withdraw(ctx context.Context, id uuid.UUID) *repository.Wallet

	Buy(ctx context.Context, amount string)

	Sell(ctx context.Context, amount string)

	Swap(ctx context.Context)
}

/**
 * walletService implements port.twalletService interface
 * and provides an access to the wallet repository
 * and cache service
 */
type WalletService struct {
	repo *repository.Repository
}


// NewWalletService creates a new instance of WalletService.
//
// Parameters:
// - db: The SQL database connection.
//
// Returns:
// - WalletServiceInterface: The newly created WalletService.
func NewWalletService(db *sqlx.DB) WalletServiceInterface {
	// Create a new instance of WalletService and initialize
	// the repository with the provided database connection.
	return &WalletService{
		repo: repository.NewRepository(db),
	}
}

// Buy implements WalletServiceInterface.
func (w *WalletService) Buy(ctx context.Context, amount string) {
	panic("unimplemented")
}

// Deposit implements WalletServiceInterface.
func (w *WalletService) Deposit(ctx context.Context, id uuid.UUID) *repository.Wallet {
	panic("unimplemented")
}

// Sell implements WalletServiceInterface.
func (w *WalletService) Sell(ctx context.Context, amount string) {
	panic("unimplemented")
}

// Swap implements WalletServiceInterface.
func (w *WalletService) Swap(ctx context.Context) {
	panic("unimplemented")
}

// Withdraw implements WalletServiceInterface.
func (w *WalletService) Withdraw(ctx context.Context, id uuid.UUID) *repository.Wallet {
	panic("unimplemented")
}

// Deletewallet implements WalletServiceInterface.
func (w *WalletService) Deletewallet(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// Getwallet implements WalletServiceInterface.
func (w *WalletService) Getwallet(ctx context.Context, id uint64) (*repository.Wallet, error) {
	panic("unimplemented")
}

// Listwallets implements WalletServiceInterface.
func (w *WalletService) Listwallets(ctx context.Context, skip uint64, limit uint64) ([]repository.Wallet, error) {
	panic("unimplemented")
}

// GenerateWallet implements WalletServiceInterface.
func (w *WalletService) GenerateWalletAddress(ctx context.Context, ethClient *ethclient.Client, wallet domain.CreateWalletRequestPayload) (string, error) {
	address, _ := walletPkg.NewKeystoreAccount(wallet.UserID.String(), ethClient)
	return address.Hex(), nil
}

// Updatewallet implements WalletServiceInterface.
func (w *WalletService) Updatewallet(ctx context.Context, wallet repository.Wallet) (*repository.Wallet, error) {
	panic("unimplemented")
}

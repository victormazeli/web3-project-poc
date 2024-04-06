package domain

import (
	"context"
	"fmt"
	"goApiStartetProject/internal/util/validator"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

type TransferCoinRequestPayload struct {
	Password	string		`json:"password" binding:"required"`
	Hash        string    `json:"hash"`
	Status      string    `json:"status"`
	CoinType    string    `json:"coin_type" binding:"required"`
	Amount      *big.Int    `json:"amount" binding:"required"`
	ToAddress   common.Address   `json:"to_address" binding:"required"`
	FromAddress common.Address    `json:"from_address"`
	GasPrice    *big.Int   `json:"gas_price"`
	GasLimit    uint64    `json:"gas_limit"`
	Nonce      uint64    `json:"nounce"`
	Data        string    `json:"data"`
	WalletID   uuid.UUID `json:"wallet_id"`
}

func (tx TransferCoinRequestPayload) GetNonce(ctx context.Context, client *ethclient.Client, fromAddress common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return common.Big1.Uint64(), err
	}

	return nonce, nil
}

func (tx TransferCoinRequestPayload) Validate(v *validator.Validator) bool {
	v.Check(tx.Password != "", "password", "password is required")
	v.Check(tx.ToAddress.String() != "", "to_address", "destination address is required")
	v.Check(tx.CoinType != "", "coin_type", "coin type is required")

	return v.Valid()
}

func (tx TransferCoinRequestPayload) SetTransferCoinReqPayload(ctx context.Context, client *ethclient.Client)  {
	fmt.Println(client)
	nonce, _ := tx.GetNonce(ctx, client, tx.FromAddress)


	gasLimit := uint64(21000) // in units

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx.GasLimit = gasLimit
	tx.GasPrice = gasPrice
	tx.Nonce    = nonce
}
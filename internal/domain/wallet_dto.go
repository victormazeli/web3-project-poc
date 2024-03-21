package domain

import "github.com/google/uuid"

const (
	MasterWallet = "Master"
	Regular	= "Regular"
)


type CreateWalletRequestPayload struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

type CreateWalletResponsePayload struct {
	WalletID uuid.UUID `json:"id"`
}

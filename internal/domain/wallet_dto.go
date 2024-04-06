package domain

import (
	"goApiStartetProject/internal/util/validator"

	"github.com/google/uuid"
)

const (
	MasterWallet = "Master"
	Regular	= "Regular"
)


type CreateWalletRequestPayload struct {
	Passphrase		string		`json:"passphrase"`
}

type CreateWalletResponsePayload struct {
	WalletID uuid.UUID `json:"id"`
}

func (w *CreateWalletRequestPayload) Validate(v *validator.Validator) bool {
	v.Check(w.Passphrase != "", "passphrase", "passphrase is required")

	return v.Valid()
}

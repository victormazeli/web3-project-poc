package validator

import "math/big"

func CanTransferAmount(trfAmount *big.Float, walletAmount *big.Float) (bool) {
	result := trfAmount.Cmp(walletAmount)
	switch result {
	case -1:
		return false
	case 0:
		return false
	default:
		return true
	}
}
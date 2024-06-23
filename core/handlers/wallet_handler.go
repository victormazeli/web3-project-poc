package handlers

import (
	"fmt"
	"goApiStartetProject/internal/domain"
	"goApiStartetProject/internal/service"
	"goApiStartetProject/internal/util/ApiResponse"
	"goApiStartetProject/internal/util/validator"

	"github.com/gin-gonic/gin"
)

type WalletHandlerInterface interface {
	CreateWallet(c *gin.Context)
	ImportWallet(c *gin.Context)
	// GetWallet(c *gin.Context)
	UpdateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
}

type WalletHandler struct {
	Handler *Handler
}

func NewWalletHandler(svc *service.Service) WalletHandlerInterface {
	return &WalletHandler{
		Handler: NewHandler(svc),
	}
}

func (w *WalletHandler) CreateWallet(c *gin.Context) {
	// Handle user creation logic
	var newWallet domain.WalletRequestPayload

	// fmt.Println(newWallet)
	if err := c.ShouldBindJSON(&newWallet); err != nil {
		ApiResponse.SendBadRequest(c, err.Error())
		return
	}
	v := validator.New()
	if !newWallet.Validate(v) {
		ApiResponse.SendValidationError(c, validator.NewValidationError("validation failed", v.Errors))
		return
	}

	walletRespPayload, err := w.Handler.Service.WalletService.GenerateWalletAddresses(c, w.Handler.Service.EthClient, newWallet)
	if err != nil {
		ApiResponse.SendInternalServerError(c, fmt.Sprintf("%s", err.Error()))
		return
	}

	ApiResponse.SendCreated(c, "Wallet Created", walletRespPayload)
}


// DeleteWallet implements WalletHandlerInterface.
func (w *WalletHandler) DeleteWallet(c *gin.Context) {
	panic("unimplemented")
}

// GetWallet implements WalletHandlerInterface.
func (w *WalletHandler) ImportWallet(c *gin.Context) {
	// Handle user creation logic
	var existingWallet domain.WalletRequestPayload

	// fmt.Println(newWallet)
	if err := c.ShouldBindJSON(&existingWallet); err != nil {
		ApiResponse.SendBadRequest(c, err.Error())
		return
	}
	v := validator.New()
	if !existingWallet.Validate(v) {
		ApiResponse.SendValidationError(c, validator.NewValidationError("validation failed", v.Errors))
		return
	}

	walletRespPayload, err := w.Handler.Service.WalletService.GenerateWalletAddresses(c, w.Handler.Service.EthClient, existingWallet)
	if err != nil {
		ApiResponse.SendInternalServerError(c, fmt.Sprintf("%s", err.Error()))
		return
	}

	ApiResponse.SendCreated(c, "Wallet Created", walletRespPayload)
}

// UpdateWallet implements WalletHandlerInterface.
func (w *WalletHandler) UpdateWallet(c *gin.Context) {
	panic("unimplemented")
}

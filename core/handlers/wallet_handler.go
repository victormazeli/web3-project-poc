package handlers

import (
	"fmt"
	"goApiStartetProject/internal/domain"
	"goApiStartetProject/internal/service"
	"goApiStartetProject/internal/util/ApiResponse"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type WalletHandlerInterface interface {
	CreateWallet(c *gin.Context)
	GetWallet(c *gin.Context)
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

var customWalletErrorMessages = map[string]map[string]string{
	"user_id": {
		"required": "UserID is required.",
	},
}

func (w *WalletHandler) CreateWallet(c *gin.Context) {
	// Handle user creation logic
	var newWallet domain.CreateWalletRequestPayload

	// fmt.Println(newWallet)
	if err := c.ShouldBindJSON(&newWallet); err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMsg string
		for _, e := range errs {
			if customMessages, ok := customErrorMessages[e.StructField()]; ok {
				if customMessage, ok := customMessages[e.Tag()]; ok {
					errorMsg += customMessage + " "
				}
			}
		}
		ApiResponse.SendBadRequest(c, errorMsg)
		return
	}

	walletRespPayload, err := w.Handler.Service.WalletService.GenerateWalletAddress(c, w.Handler.Service.EthClient, newWallet)
	if err != nil {
		ApiResponse.SendInternalServerError(c, fmt.Sprintf("%s", err.Error()))
	}

	ApiResponse.SendCreated(c, "Wallet Created", walletRespPayload)
}


// DeleteWallet implements WalletHandlerInterface.
func (w *WalletHandler) DeleteWallet(c *gin.Context) {
	panic("unimplemented")
}

// GetWallet implements WalletHandlerInterface.
func (w *WalletHandler) GetWallet(c *gin.Context) {
	panic("unimplemented")
}

// UpdateWallet implements WalletHandlerInterface.
func (w *WalletHandler) UpdateWallet(c *gin.Context) {
	panic("unimplemented")
}

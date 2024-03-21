package handlers

import (
	"goApiStartetProject/internal/domain"
	"goApiStartetProject/internal/service"
	"goApiStartetProject/internal/util/ApiResponse"
	"goApiStartetProject/internal/util/validator"

	"github.com/gin-gonic/gin"
)

type TransactionHandlerInterface interface {
	HandleNewTransaction(c *gin.Context)
	GetTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
	DeleteTransaction(c *gin.Context)
}

type TransactionHandler struct {
	Handler *Handler
}

// DeleteTransaction implements TransactionHandlerInterface.
func (w *TransactionHandler) DeleteTransaction(c *gin.Context) {
	panic("unimplemented")
}

// GetTransaction implements TransactionHandlerInterface.
func (w *TransactionHandler) GetTransaction(c *gin.Context) {
	panic("unimplemented")
}

// UpdateTransaction implements TransactionHandlerInterface.
func (w *TransactionHandler) UpdateTransaction(c *gin.Context) {
	panic("unimplemented")
}

func NewTransactionHandler(svc *service.Service) TransactionHandlerInterface {
	return &TransactionHandler{
		Handler: NewHandler(svc),
	}
}

func (w *TransactionHandler) HandleNewTransaction(c *gin.Context) {
	// Handle user creation logic
	var newTransaction domain.TransferCoinRequestPayload

	// fmt.Println(newTransaction)
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		ApiResponse.SendBadRequest(c, err.Error())
		return
	}

	v := validator.New()
	if !newTransaction.Validate(v) {
		ApiResponse.SendValidationError(c, validator.NewValidationError("validation failed", v.Errors))
		return
	}

	hash, err := w.Handler.Service.TransactionService.TransferCoin(c, w.Handler.Service.EthClient, newTransaction)
	if err != nil {
		ApiResponse.SendInternalServerError(c, err.Error())
	}

	txReciept, err := w.Handler.Service.EthClient.TransactionReceipt(c, hash)
	if err != nil {
		ApiResponse.SendInternalServerError(c, err.Error())
	}

	ApiResponse.SendCreated(c, "Transaction Created", txReciept)
}

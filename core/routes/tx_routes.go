package routes

import (
	"goApiStartetProject/core/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func TxRoute(handler *handlers.Handler, db *sqlx.DB, r *gin.RouterGroup) {

	txHandler := handlers.TransactionHandler{
		Handler: handler,
	}

	r.POST("/new-transaction", txHandler.HandleNewTransaction)
	// r.POST("/register", userHandler.CreateUser)
}

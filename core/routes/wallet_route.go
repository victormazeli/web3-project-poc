package routes


import (
	"goApiStartetProject/core/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func WalletRoute(handler *handlers.Handler, db *sqlx.DB, r *gin.RouterGroup) {

	walletHandler := handlers.WalletHandler{
		Handler: handler,
	}

	// r.GET("/:wallet-id", walletHandler.GetWallet)
	r.POST("/create-wallet", walletHandler.CreateWallet)
}

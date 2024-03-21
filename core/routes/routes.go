package routes

import (
	"goApiStartetProject/core/config"
	"goApiStartetProject/core/handlers"
	"goApiStartetProject/internal/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoute(env *config.Env, db *sqlx.DB, ethClient *ethclient.Client, rg *gin.RouterGroup) {
	srv := service.NewService(env, ethClient, db)
	handler := handlers.NewHandler(srv)
	UserRoute(handler, db, rg.Group("/users"))
	WalletRoute(handler, db, rg.Group("/wallets"))
	TxRoute(handler, db, rg.Group("/transactions"))
}

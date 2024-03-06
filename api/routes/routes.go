package routes

import (
	"goApiStartetProject/api/handlers"
	"goApiStartetProject/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoute(env *config.Env, db *sqlx.DB, ethClient *ethclient.Client, rg *gin.RouterGroup) {

	handler := handlers.NewHandler(ethClient, env)
	UserRoute(handler, db, rg)

}

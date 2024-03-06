package routes

import (
	"goApiStartetProject/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoute(env *config.Env, db *sqlx.DB, ethClient *ethclient.Client, rg *gin.RouterGroup) {
	UserRoute(env, db, ethClient, rg)

}

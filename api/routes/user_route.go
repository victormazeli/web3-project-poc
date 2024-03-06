package routes

import (
	"goApiStartetProject/api/handlers"
	"goApiStartetProject/config"
	"goApiStartetProject/db/repository"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRoute(env *config.Env, db *sqlx.DB, ethClient *ethclient.Client, r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)

	userHandler := handlers.UserHandler{
		Repo: userRepo,
		Handler: handlers.Handler{
			Env: env,
			EthClient: ethClient,
		},
	}

	r.GET("/:id", userHandler.GetUser)
	r.POST("/create_user", userHandler.CreateUser)
}

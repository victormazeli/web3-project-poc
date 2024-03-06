package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/jmoiron/sqlx"
// 	"goApiStartetProject/api/handlers"
// 	"goApiStartetProject/config"
// 	"goApiStartetProject/db/repository"
// )

// func WalletRoute(env *config.Env, db *sqlx.DB, r *gin.RouterGroup) {
// 	WalletRepo := repository.NewWalletRepository(db)

// 	WalletHandler := handlers.WalletHandler{
// 		Env:  env,
// 		Repo: WalletRepo,
// 	}

// 	r.GET("/:id", WalletHandler.GetWallet)
// 	r.POST("/create_wallet", WalletHandler.CreateWallet)
// }

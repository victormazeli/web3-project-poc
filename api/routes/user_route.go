package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"goApiStartetProject/api/handlers"
	"goApiStartetProject/config"
	"goApiStartetProject/db/repository"
)

func UserRoute(env *config.Env, db *sqlx.DB, r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)

	userHandler := handlers.UserHandler{
		Env:  env,
		Repo: userRepo,
	}

	r.GET("/:id", userHandler.GetUser)
	r.POST("/create_user", userHandler.CreateUser)
}

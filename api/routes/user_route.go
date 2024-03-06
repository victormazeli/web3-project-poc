package routes

import (
	"goApiStartetProject/api/handlers"
	"goApiStartetProject/db/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRoute(handler *handlers.Handler, db *sqlx.DB, r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)

	userHandler := handlers.UserHandler{
		Repo: userRepo,
		Handler: handler,
	}

	r.GET("/:id", userHandler.GetUser)
	r.POST("/create_user", userHandler.CreateUser)
}

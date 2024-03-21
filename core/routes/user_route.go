package routes

import (
	"goApiStartetProject/core/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRoute(handler *handlers.Handler, db *sqlx.DB, r *gin.RouterGroup) {

	userHandler := handlers.UserHandler{
		Handler: handler,
	}

	r.GET("/:id", userHandler.GetUser)
	r.POST("/register", userHandler.CreateUser)
}

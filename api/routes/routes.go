package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"goApiStartetProject/config"
)

func SetupRoute(env *config.Env, db *sqlx.DB, rg *gin.RouterGroup) {
	UserRoute(env, db, rg)

}

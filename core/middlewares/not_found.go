package middlewares

import (
	"github.com/gin-gonic/gin"
	"goApiStartetProject/internal/util/ApiResponse"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiResponse.SendNotFound(c, "resource not found")
		c.Next()
	}
}

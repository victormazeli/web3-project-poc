package handlers

import "github.com/gin-gonic/gin"

type AccountHandlerInterface interface {
	GetAccount(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type AccountHandler struct {
	Handler *Handler
}

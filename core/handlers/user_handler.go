package handlers

import (
	"fmt"
	"goApiStartetProject/internal/domain"
	"goApiStartetProject/internal/service"
	"goApiStartetProject/internal/util/ApiResponse"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandler struct {
	Handler *Handler
}

var customErrorMessages = map[string]map[string]string{
	"Email": {
		"required": "Email is required.",
		"email":    "Email is invalid.",
	},
	"Password": {
		"required": "Password is required.",
		"min":      "Password must be at least 8 characters long.",
	},
	"Country": {
		"required": "Username is required.",
	},
}

func NewUserHandler(svc *service.Service) UserHandlerInterface {
	return &UserHandler{
		Handler: NewHandler(svc),
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	// Handle user creation logic
	var newUser domain.NewUserRequestPayload

	// fmt.Println(newUser)
	if err := c.ShouldBindJSON(&newUser); err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMsg string
		for _, e := range errs {
			// fmt.Println(e.)
			if customMessages, ok := customErrorMessages[e.StructField()]; ok {
				if customMessage, ok := customMessages[e.Tag()]; ok {
					errorMsg += customMessage + " "
				}
			}
		}
		ApiResponse.SendBadRequest(c, errorMsg)
		return
	}

	// userRespPayload, err := u.Handler.Service.UserService.Register(c, newUser)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	ApiResponse.SendInternalServerError(c, err.Error())
	// 	return
	// }

	id := uuid.New()

	fmt.Println(id)
	ApiResponse.SendCreated(c, "User created", id)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	// Fetch user from database by userID
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "name": "John Doe"})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	// Update user in the database by userID
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	// Delete user from the database by userID
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goApiStartetProject/db/repository"
	"log"
	"net/http"
	"strconv"
)

type NewUser struct {
	Username string `json:"username" mod:"trim" binding:"required"`
	Email    string `json:"email" mod:"trim" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserHandler struct {
	Repo repository.UserRepositoryInterface
	Handler Handler
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
	"Username": {
		"required": "Username is required.",
	},
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	// Handle user creation logic
	var newUser NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMsg string
		for _, e := range errs {
			if customMessages, ok := customErrorMessages[e.StructField()]; ok {
				if customMessage, ok := customMessages[e.Tag()]; ok {
					errorMsg += customMessage + " "
				}
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}

	userModel := repository.User{
		Email:    newUser.Email,
		Password: newUser.Password,
		Username: newUser.Username,
	}

	createUser, err := u.Repo.CreateUser(c, &userModel)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unable to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "data": createUser})
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

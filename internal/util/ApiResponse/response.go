package ApiResponse

import (
	"goApiStartetProject/internal/util/validator"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response represents a standard JSON response structure.
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}


type Data struct {
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Status    string      `json:"status,omitempty"`
}

// SendSuccess sends a JSON response with a status code of 200 (OK).
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// SendCreated sends a JSON response with a status code of 201 (Created).
func SendCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

// SendError sends a JSON response with a specified status code and error message.
func SendError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
	})
}

// SendBadRequest sends a JSON response with a status code of 400 (Bad Request).
func SendBadRequest(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, message)
}

// SendUnauthorized sends a JSON response with a status code of 401 (Unauthorized).
func SendUnauthorized(c *gin.Context, message string) {
	SendError(c, http.StatusUnauthorized, message)
}

// SendForbidden sends a JSON response with a status code of 403 (Forbidden).
func SendForbidden(c *gin.Context, message string) {
	SendError(c, http.StatusForbidden, message)
}

// SendNotFound sends a JSON response with a status code of 404 (Not Found).
func SendNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, message)
}

// SendInternalServerError sends a JSON response with a status code of 500 (Internal Server Error).
func SendInternalServerError(c *gin.Context, message string) {
	SendError(c, http.StatusInternalServerError, message)
}

// SendMethodNotAllowedError sends a JSON response with a status code of 405 (Method Not Allowed).
func SendMethodNotAllowedError(c *gin.Context, message string) {
	SendError(c, http.StatusMethodNotAllowed, message)
}

func SendValidationError(ctx *gin.Context, errors validator.ValidationError) {
	responseData := Data{
		Errors:    errors.Fields,
		Message:   errors.Message,
		Status:    http.StatusText(http.StatusUnprocessableEntity),
		Timestamp: time.Now().Format(time.RFC850),
	}

	ctx.JSON(http.StatusUnprocessableEntity, responseData)

}

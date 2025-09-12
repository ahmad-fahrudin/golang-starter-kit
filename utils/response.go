package utils

import (
	"net/http"

	"golang-starter-kit/src/models"

	"github.com/gin-gonic/gin"
)

// RespondError sends a standardized error response
func RespondError(c *gin.Context, status int, code string, message string) {
	c.JSON(status, models.ErrorResponse{
		Error:   code,
		Message: message,
	})
}

// Convenience wrappers for common error statuses
func BadRequest(c *gin.Context, code, message string) {
	RespondError(c, http.StatusBadRequest, code, message)
}

func InvalidRequest(c *gin.Context, err error) {
	BadRequest(c, "invalid_request", err.Error())
}

func ValidationError(c *gin.Context, err error) {
	BadRequest(c, "validation_error", err.Error())
}

func Unauthorized(c *gin.Context, message string) {
	RespondError(c, http.StatusUnauthorized, "unauthorized", message)
}

func NotFound(c *gin.Context, code, message string) {
	RespondError(c, http.StatusNotFound, code, message)
}

func Conflict(c *gin.Context, code, message string) {
	RespondError(c, http.StatusConflict, code, message)
}

func InternalServerError(c *gin.Context, code, message string) {
	RespondError(c, http.StatusInternalServerError, code, message)
}

// Success responses
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func SuccessMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"message": message, "data": data})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"message": message, "data": data})
}

// Generic message response
func Message(c *gin.Context, status int, message string) {
	c.JSON(status, models.MessageResponse{Message: message})
}

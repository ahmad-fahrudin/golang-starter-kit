package middleware

import (
	"net/http"
	"strings"

	"golang-starter-kit/internal/models"
	"golang-starter-kit/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates a middleware function for JWT authentication
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "unauthorized",
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		var tokenString string

		// Check if the header has the Bearer prefix
		if strings.HasPrefix(authHeader, "Bearer ") {
			// Format: "Bearer <token>"
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// Format: "<token>" (langsung token tanpa Bearer)
			tokenString = authHeader
		}

		// Validate that we have a token
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "unauthorized",
				Message: "Token is required",
			})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "unauthorized",
				Message: "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}

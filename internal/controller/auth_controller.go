package controller

import (
	"net/http"

	"golang-starter-kit/internal/models"
	"golang-starter-kit/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
	userService service.UserService
	validator   *validator.Validate
}

// NewAuthController creates a new auth controller
func NewAuthController(userService service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
		validator:   validator.New(),
	}
}

// Login handles POST /auth/login
// @Summary      User Login
// @Description  Authenticate user with email and password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.LoginRequest true "Login credentials"
// @Router       /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	// Validate request
	if err := ac.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	loginResponse, err := ac.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error:   "login_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":    loginResponse,
	})
}

// Register handles POST /auth/register
// @Summary      User Registration
// @Description  Register a new user account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.UserCreateRequest true "User registration data"
// @Router       /auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	// Validate request
	if err := ac.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	user, err := ac.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error:   "registration_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"data":    user,
	})
}

// Logout handles POST /auth/logout
// @Summary      User Logout
// @Description  Logout user (client-side token removal)
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Router       /auth/logout [post]
func (ac *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, models.MessageResponse{
		Message: "Logout successful",
	})
}

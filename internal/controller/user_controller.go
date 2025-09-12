package controller

import (
	"net/http"

	"golang-starter-kit/internal/models"
	"golang-starter-kit/internal/service"
	"golang-starter-kit/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserController handles user-related HTTP requests
type UserController struct {
	userService service.UserService
	validator   *validator.Validate
}

// NewUserController creates a new user controller
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
		validator:   validator.New(),
	}
}

// GetUsersWithPagination handles POST /users/pagination
// @Summary      Get Users with Pagination
// @Description  Retrieve paginated list of users with optional filters
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body models.UserListRequest true "Pagination and filter parameters"
// @Success      200 {object} models.UsersListResponse
// @Router       /users/pagination [post]
func (uc *UserController) GetUsersWithPagination(c *gin.Context) {
	var req models.UserListRequest

	// Bind JSON body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidRequest(c, err)
		return
	}

	// Validate request
	if err := uc.validator.Struct(req); err != nil {
		utils.ValidationError(c, err)
		return
	}

	// Set default values if not provided
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	response, err := uc.userService.GetAllUsersWithFilter(req)
	if err != nil {
		utils.InternalServerError(c, "search_failed", err.Error())
		return
	}

	utils.Success(c, response)
}

// CreateUser handles POST /users
// @Summary      Create User
// @Description  Create a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body models.UserCreateRequest true "User data"
// @Success      201 {object} models.UserResponse
// @Router       /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidRequest(c, err)
		return
	}

	// Validate request
	if err := uc.validator.Struct(req); err != nil {
		utils.ValidationError(c, err)
		return
	}

	user, err := uc.userService.CreateUser(req)
	if err != nil {
		utils.Conflict(c, "creation_failed", err.Error())
		return
	}

	utils.Created(c, "User created successfully", user)
}

// GetUser handles GET /users/:id
// @Summary      Get User by ID
// @Description  Retrieve a user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} models.UserResponse
// @Router       /users/{id} [get]
func (uc *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := utils.StringToUint(idParam)
	if err != nil {
		utils.BadRequest(c, "invalid_id", "Invalid user ID")
		return
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		utils.NotFound(c, "user_not_found", err.Error())
		return
	}

	utils.Success(c, user)
}

// UpdateUser handles PUT /users/:id
// @Summary      Update User
// @Description  Update an existing user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        request body models.UserUpdateRequest true "User update data"
// @Success      200 {object} models.UserResponse
// @Router       /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := utils.StringToUint(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid user ID",
		})
		return
	}

	var req models.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidRequest(c, err)
		return
	}

	// Validate request
	if err := uc.validator.Struct(req); err != nil {
		utils.ValidationError(c, err)
		return
	}

	user, err := uc.userService.UpdateUser(id, req)
	if err != nil {
		utils.BadRequest(c, "update_failed", err.Error())
		return
	}

	utils.SuccessMessage(c, "User updated successfully", user)
}

// DeleteUser handles DELETE /users/:id
// @Summary      Delete User
// @Description  Delete a user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} models.MessageResponse
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := utils.StringToUint(idParam)
	if err != nil {
		utils.BadRequest(c, "invalid_id", "Invalid user ID")
		return
	}

	if err := uc.userService.DeleteUser(id); err != nil {
		utils.NotFound(c, "delete_failed", err.Error())
		return
	}

	utils.Message(c, http.StatusOK, "User deleted successfully")
}

// GetProfile handles GET /profile (protected route)
// @Summary      Get User Profile
// @Description  Get the authenticated user's profile
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.UserResponse
// @Router       /profile [get]
func (uc *UserController) GetProfile(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		utils.Unauthorized(c, "User not authenticated")
		return
	}

	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		utils.NotFound(c, "user_not_found", err.Error())
		return
	}

	utils.Success(c, user)
}

// UpdateProfile handles PUT /profile (protected route)
// @Summary      Update User Profile
// @Description  Update the authenticated user's profile
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body models.UserUpdateRequest true "Profile update data"
// @Success      200 {object} models.UserResponse
// @Router       /profile [put]
func (uc *UserController) UpdateProfile(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		utils.Unauthorized(c, "User not authenticated")
		return
	}

	var req models.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidRequest(c, err)
		return
	}

	// Validate request
	if err := uc.validator.Struct(req); err != nil {
		utils.ValidationError(c, err)
		return
	}

	user, err := uc.userService.UpdateUser(userID, req)
	if err != nil {
		utils.BadRequest(c, "update_failed", err.Error())
		return
	}

	utils.SuccessMessage(c, "Profile updated successfully", user)
}

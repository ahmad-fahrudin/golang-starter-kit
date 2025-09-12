package service

import (
	"errors"
	"golang-starter-kit/src/models"
	"golang-starter-kit/src/repository"
	"golang-starter-kit/utils"

	"gorm.io/gorm"
)

// UserService interface defines user service methods
type UserService interface {
	CreateUser(req models.UserCreateRequest) (*models.UserResponse, error)
	GetUserByID(id uint) (*models.UserResponse, error)
	UpdateUser(id uint, req models.UserUpdateRequest) (*models.UserResponse, error)
	DeleteUser(id uint) error
	GetAllUsersWithFilter(req models.UserListRequest) (*models.UsersListResponse, error)
	Login(req models.LoginRequest) (*models.LoginResponse, error)
}

// userService implements UserService interface
type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository, jwtSecret string) UserService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// CreateUser creates a new user
func (s *userService) CreateUser(req models.UserCreateRequest) (*models.UserResponse, error) {
	// Check if user with email already exists
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// GetUserByID gets a user by ID
func (s *userService) GetUserByID(id uint) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// UpdateUser updates a user
func (s *userService) UpdateUser(id uint, req models.UserUpdateRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Check if email is already taken by another user
		existingUser, err := s.userRepo.GetByEmail(req.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email is already taken")
		}
		user.Email = req.Email
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(id uint) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.userRepo.Delete(user.ID)
}

// GetAllUsersWithFilter gets all users with filters and pagination
func (s *userService) GetAllUsersWithFilter(req models.UserListRequest) (*models.UsersListResponse, error) {
	// Set default values
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100 // Max limit to prevent abuse
	}

	users, total, err := s.userRepo.GetAllWithFilter(req)
	if err != nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	// Calculate total pages
	totalPages := int(total) / req.Limit
	if int(total)%req.Limit > 0 {
		totalPages++
	}

	response := &models.UsersListResponse{
		Data: responses,
		Pagination: models.Pagination{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}

	return response, nil
}

// Login authenticates a user and returns a JWT token
func (s *userService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	response := &models.LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	}

	return response, nil
}

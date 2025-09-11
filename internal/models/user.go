package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password  string         `json:"-" gorm:"not null" validate:"required,min=6"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// UserCreateRequest represents the request payload for creating a user
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john@example.com"`
	Password string `json:"password" validate:"required,min=6" example:"password123"`
}

// UserUpdateRequest represents the request payload for updating a user
type UserUpdateRequest struct {
	Name  string `json:"name" validate:"omitempty,min=2,max=100" example:"Jane Doe"`
	Email string `json:"email" validate:"omitempty,email" example:"jane@example.com"`
}

// UserResponse represents the response payload for user data (without password)
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// UsersListResponse represents the response payload for users list with pagination
type UsersListResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination Pagination     `json:"pagination"`
}

// Pagination represents pagination information
type Pagination struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
	Total int `json:"total" example:"100"`
}

// ToResponse converts User model to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

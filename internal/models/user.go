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
	Data       []UserResponse `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

// UsersData represents the data payload for users list
type UsersData struct {
	Users []UserResponse `json:"users"`
}

// Pagination represents pagination information
type Pagination struct {
	Page       int   `json:"page" example:"1"`
	Limit      int   `json:"limit" example:"10"`
	Total      int64 `json:"total" example:"100"`
	TotalPages int   `json:"total_pages" example:"10"`
}

// UserFilter represents filter options for user listing
type UserFilter struct {
	Name      string `form:"name" json:"name"`
	Email     string `form:"email" json:"email"`
	SortBy    string `form:"sort_by" json:"sort_by" example:"name,email,created_at"`
	SortOrder string `form:"sort_order" json:"sort_order" example:"asc,desc"`
}

// UserListRequest represents the request payload for listing users with filters
type UserListRequest struct {
	Page   int        `form:"page" json:"page" example:"1"`
	Limit  int        `form:"limit" json:"limit" example:"10"`
	Filter UserFilter `form:"filter" json:"filter"`
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

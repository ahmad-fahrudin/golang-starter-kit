package models

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// LoginResponse represents the response payload for successful login
type LoginResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponse `json:"user"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message" example:"Operation successful"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"validation_error"`
	Message string `json:"message,omitempty" example:"Invalid input data"`
}

package repository

import (
	"golang-starter-kit/internal/models"

	"gorm.io/gorm"
)

// UserRepository interface defines user repository methods
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	GetAllWithFilter(req models.UserListRequest) ([]models.User, int64, error)
}

// userRepository implements UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID gets a user by ID
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail gets a user by email
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete soft deletes a user
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetAllWithFilter gets all users with filters and pagination
func (r *userRepository) GetAllWithFilter(req models.UserListRequest) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Build query with filters
	query := r.db.Model(&models.User{})

	// Apply filters
	if req.Filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+req.Filter.Name+"%")
	}
	if req.Filter.Email != "" {
		query = query.Where("email ILIKE ?", "%"+req.Filter.Email+"%")
	}

	// Get total count with filters
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Apply sorting
	orderBy := "created_at desc" // default sorting
	if req.Filter.SortBy != "" {
		sortOrder := "asc"
		if req.Filter.SortOrder == "desc" {
			sortOrder = "desc"
		}

		// Validate sort fields
		validSortFields := map[string]bool{
			"id":         true,
			"name":       true,
			"email":      true,
			"created_at": true,
			"updated_at": true,
		}

		if validSortFields[req.Filter.SortBy] {
			orderBy = req.Filter.SortBy + " " + sortOrder
		}
	}

	// Calculate offset
	offset := (req.Page - 1) * req.Limit

	// Execute query with pagination
	err = query.Order(orderBy).Limit(req.Limit).Offset(offset).Find(&users).Error
	return users, total, err
}

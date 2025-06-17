package repository

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDB()}
}

// FindByUsername finds a user by username
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("UserCredentials").
		Where("username = ? AND status = ?", username, true).
		First(&user).Error
	return &user, err
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("UserCredentials").
		First(&user, id).Error
	return &user, err
}

// SearchUsers searches users by name or username with pagination
func (r *UserRepository) SearchUsers(name string, page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{}).
		Preload("UserCredentials").
		Where("CONCAT(first_name, ' ', last_name) LIKE ? OR username LIKE ?", "%"+name+"%", "%"+name+"%").
		Where("deleted_at IS NULL")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Limit(pageSize).Offset(offset).Find(&users).Error
	return users, total, err
}

// CheckDuplicate checks for duplicate username, phone, or email
func (r *UserRepository) CheckDuplicate(username, phone, email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ? OR phone = ? OR email = ?", username, phone, email).
		First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

// CheckDuplicateForUpdate checks for duplicate phone or email excluding the given ID
func (r *UserRepository) CheckDuplicateForUpdate(id uint, phone, email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id != ? AND (phone = ? OR email = ?)", id, phone, email).
		First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// UpdatePassword updates the user's password
func (r *UserRepository) UpdatePassword(userID uint, password string) error {
	return r.db.Model(&models.UserCredentials{}).
		Where("user_id = ?", userID).
		Update("password", password).Error
}

// Delete soft deletes a user
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

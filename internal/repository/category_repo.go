package repository

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/pkg/database"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: database.GetDB()}
}

func (r *CategoryRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) FindAll(page, pageSize int, search string) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	query := r.db.Model(&models.Category{})
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&categories).Error
	return categories, total, err
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

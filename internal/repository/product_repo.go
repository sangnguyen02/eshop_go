package repository

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/pkg/database"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: database.GetDB()}
}

// FindByName searches products by name with pagination
func (r *ProductRepository) FindByName(name string, page, pageSize int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{}).
		Preload("Category").
		Preload("Brand").
		Preload("Images").
		Where("Name LIKE ?", "%"+name+"%")

	// Đếm tổng số sản phẩm
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Phân trang
	offset := (page - 1) * pageSize
	err := query.Limit(pageSize).Offset(offset).Find(&products).Error
	return products, total, err
}

// FindByID finds a product by ID with all related data
func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").
		Preload("Brand").
		Preload("Images").
		Preload("Variants").
		Preload("Reviews").
		First(&product, id).Error
	return &product, err
}

// Create creates a new product
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// Update updates an existing product
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete soft deletes a product
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

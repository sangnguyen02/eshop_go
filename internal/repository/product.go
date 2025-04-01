package repository

import (
	"go_ecommerce/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByName(name string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("Name LIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

package repository

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/pkg/database"

	"gorm.io/gorm"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository() *BrandRepository {
	return &BrandRepository{db: database.GetDB()}
}

func (r *BrandRepository) Create(brand *model.Brand) error {
	return r.db.Create(brand).Error
}

func (r *BrandRepository) FindByID(id uint) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.First(&brand, id).Error
	return &brand, err
}

func (r *BrandRepository) FindAll(page, pageSize int, search string, status bool) ([]model.Brand, int64, error) {
	var brands []model.Brand
	var total int64

	query := r.db.Model(&model.Brand{}).Where("deleted_at IS NULL")

	if !status {
		query = query.Where("status = ?", true)
	}

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	err = query.Offset(offset).Limit(pageSize).Find(&brands).Error
	return brands, total, err
}

func (r *BrandRepository) Update(brand *model.Brand) error {
	return r.db.Save(brand).Error
}

func (r *BrandRepository) Delete(id uint) error {
	return r.db.Delete(&model.Brand{}, id).Error
}

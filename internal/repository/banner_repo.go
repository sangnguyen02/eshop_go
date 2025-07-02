package repository

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/pkg/database"

	"gorm.io/gorm"
)

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository() *BannerRepository {
	return &BannerRepository{db: database.GetDB()}
}

func (r *BannerRepository) Create(banner *model.Banner) error {
	return r.db.Create(banner).Error
}

func (r *BannerRepository) FindByID(id uint) (*model.Banner, error) {
	var banner model.Banner
	err := r.db.First(&banner, id).Error
	return &banner, err
}

func (r *BannerRepository) FindAll(page, pageSize int, search string, status bool) ([]model.Banner, int64, error) {
	var banners []model.Banner
	var total int64

	query := r.db.Model(&model.Banner{}).Where("deleted_at IS NULL")

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

	err = query.Offset(offset).Limit(pageSize).Find(&banners).Error
	return banners, total, err
}

func (r *BannerRepository) Update(banner *model.Banner) error {
	return r.db.Save(banner).Error
}

func (r *BannerRepository) Delete(id uint) error {
	return r.db.Delete(&model.Banner{}, id).Error
}

package repository

import (
	"go_ecommerce/internal/dto"
	"go_ecommerce/internal/model"
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
func (r *ProductRepository) FindByName(name string, page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.Model(&model.Product{}).
		Preload("Category").
		Preload("Brand").
		Preload("Images").
		Where("products.deleted_at IS NULL")

	if name != "undefined" && name != "" {
		query = query.Where("products.name LIKE ?", "%"+name+"%")
	}

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
func (r *ProductRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Category").
		Preload("Brand").
		Preload("Images").
		Preload("Variants").
		Preload("Reviews").
		First(&product, id).Error
	return &product, err
}

// Create a new product
func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

// Update updates an existing product
func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

// Delete soft deletes a product
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

// #region customiztion

// FindForCard searches products for card display

type TempProductCard struct {
	ID                    uint                `gorm:"column:id"`
	Name                  string              `gorm:"column:name"`
	Price                 float64             `gorm:"column:price"`
	DiscountPrice         float64             `gorm:"column:discount_price"`
	BrandName             string              `gorm:"column:brand_name"`
	PrimaryImageURL       string              `gorm:"column:primary_image_url"`
	PrimaryImageIsPrimary bool                `gorm:"column:primary_image_is_primary"`
	Status                model.ProductStatus `gorm:"column:status"`
}

func (r *ProductRepository) FindForCard(name string, page, pageSize int) ([]dto.ProductCardResponse, int64, error) {
	var tempProducts []TempProductCard
	var total int64

	query := r.db.Debug().Model(&model.Product{}).
		Select("products.id, products.name, products.price, products.discount_price, brands.name as brand_name, product_images.url as primary_image_url, product_images.is_primary as primary_image_is_primary, products.status").
		Joins("LEFT JOIN brands ON products.brand_id = brands.id").
		Joins("LEFT JOIN product_images ON products.id = product_images.product_id AND product_images.is_primary = ?", true).
		Where("products.deleted_at IS NULL")

	if name != "undefined" && name != "" {
		query = query.Where("products.name ILIKE ?", "%"+name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Limit(pageSize).Offset(offset).Find(&tempProducts).Error

	if err != nil {
		return nil, 0, err
	}

	products := make([]dto.ProductCardResponse, len(tempProducts))
	for i, temp := range tempProducts {
		products[i] = dto.ProductCardResponse{
			ID:            temp.ID,
			Name:          temp.Name,
			Price:         temp.Price,
			DiscountPrice: temp.DiscountPrice,
			Brand: struct {
				Name string `json:"name"`
			}{Name: temp.BrandName},
			PrimaryImage: struct {
				URL string `json:"url"`
			}{URL: temp.PrimaryImageURL},
			Status: temp.Status,
		}
	}

	return products, total, err

}

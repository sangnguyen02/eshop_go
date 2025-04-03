package models

// ProductImage represents an image of a product
type ProductImage struct {
	BaseModel
	ProductID uint   `gorm:"column:product_id;index:idx_product_image,priority:1" json:"productId"` // Composite index với IsPrimary
	URL       string `gorm:"column:url" json:"url"`
	IsPrimary bool   `gorm:"column:is_primary;index:idx_product_image,priority:2" json:"isPrimary"` // Index cho ảnh chính
}

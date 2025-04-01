package models

// ProductImage represents an image of a product
type ProductImage struct {
	BaseModel
	ProductID uint   `gorm:"column:product_id" json:"productId"`
	URL       string `gorm:"column:url" json:"url"`
	IsPrimary bool   `gorm:"column:is_primary" json:"isPrimary"`
}

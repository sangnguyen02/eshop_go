package models

// ProductVariant represents a variant of a product (e.g., color, size)
type ProductVariant struct {
	BaseModel
	ProductID     uint    `gorm:"column:product_id" json:"productId"`
	SKU           string  `gorm:"column:sku" json:"sku"`
	Name          string  `gorm:"column:name" json:"name"`
	Price         float64 `gorm:"column:price" json:"price"`
	StockQuantity int     `gorm:"column:stock_quantity" json:"stockQuantity"`
	Attributes    string  `gorm:"column:attributes" json:"attributes"`
}

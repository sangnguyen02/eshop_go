package model

// import "gorm.io/gorm"

type ProductVariant struct {
	BaseModel
	ProductID     uint    `gorm:"column:product_id;index" json:"productId"`
	SKU           string  `gorm:"column:sku;type:VARCHAR(100);uniqueIndex" json:"sku"`
	Name          string  `gorm:"column:name" json:"name"`
	Price         float64 `gorm:"column:price;index" json:"price"`
	StockQuantity int     `gorm:"column:stock_quantity;index" json:"stockQuantity"`
	Attributes    string  `gorm:"column:attributes" json:"attributes"`
}

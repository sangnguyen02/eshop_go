package model

// import "gorm.io/gorm"

type ProductStatus string

const (
	ProductStatusActive     ProductStatus = "active"
	ProductStatusInactive   ProductStatus = "inactive"
	ProductStatusOutOfStock ProductStatus = "out_of_stock"
)

type Product struct {
	BaseModel
	SKU              string           `gorm:"column:sku;type:VARCHAR(100);uniqueIndex" json:"sku"`
	Name             string           `gorm:"column:name;index" json:"name"`
	Slug             string           `gorm:"column:slug;type:VARCHAR(255);uniqueIndex" json:"slug"`
	Description      string           `gorm:"column:description" json:"description"`
	ShortDescription string           `gorm:"column:short_description" json:"shortDescription"`
	Price            float64          `gorm:"column:price;index" json:"price"`
	DiscountPrice    float64          `gorm:"column:discount_price;index" json:"discountPrice"`
	StockQuantity    int              `gorm:"column:stock_quantity;index" json:"stockQuantity"`
	CategoryID       uint             `gorm:"column:category_id;index" json:"categoryId"`
	Category         Category         `gorm:"foreignKey:CategoryID" json:"category"`
	BrandID          uint             `gorm:"column:brand_id;index" json:"brandId"`
	Brand            Brand            `json:"brand"`
	Status           ProductStatus    `gorm:"column:status;index" json:"status"`
	Images           []ProductImage   `gorm:"foreignKey:ProductID" json:"images"`
	Variants         []ProductVariant `gorm:"foreignKey:ProductID" json:"variants"`
	Reviews          []ProductReview  `gorm:"foreignKey:ProductID" json:"reviews"`
}

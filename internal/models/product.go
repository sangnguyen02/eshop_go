package models

type ProductStatus string

const (
	ProductStatusActive     ProductStatus = "active"
	ProductStatusInactive   ProductStatus = "inactive"
	ProductStatusOutOfStock ProductStatus = "out_of_stock"
)

type Product struct {
	BaseModel
	SKU              string           `gorm:"column:sku" json:"sku"`
	Name             string           `gorm:"column:name" json:"name"`
	Slug             string           `gorm:"column:slug" json:"slug"`
	Description      string           `gorm:"column:description" json:"description"`
	ShortDescription string           `gorm:"column:short_description" json:"shortDescription"`
	Price            float64          `gorm:"column:price" json:"price"`
	DiscountPrice    float64          `gorm:"column:discount_price" json:"discountPrice"`
	StockQuantity    int              `gorm:"column:stock_quantity" json:"stockQuantity"`
	CategoryID       uint             `gorm:"column:category_id" json:"categoryId"`
	Category         Category         `json:"category"`
	BrandID          uint             `gorm:"column:brand_id" json:"brandId"`
	Brand            Brand            `json:"brand"`
	Status           ProductStatus    `gorm:"column:status" json:"status"`
	Images           []ProductImage   `json:"images"`
	Variants         []ProductVariant `json:"variants"`
	Reviews          []ProductReview  `json:"reviews"`
}

package model

// import "gorm.io/gorm"

// Brand represents a product brand
type Brand struct {
	BaseModel
	Name        string `gorm:"column:name;index" json:"name"`
	Slug        string `gorm:"column:slug;type:VARCHAR(255);uniqueIndex" json:"slug"`
	Description string `gorm:"column:description" json:"description"`
	LogoURL     string `gorm:"column:logo_url" json:"logoUrl"`
}

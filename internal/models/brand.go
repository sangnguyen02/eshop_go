package models

// Brand represents a product brand
type Brand struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`
	Slug        string `gorm:"column:slug" json:"slug"`
	Description string `gorm:"column:description" json:"description"`
	LogoURL     string `gorm:"column:logo_url" json:"logoUrl"`
}

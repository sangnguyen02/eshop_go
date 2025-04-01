package models

// Category represents a product category
type Category struct {
	BaseModel
	Name        string    `gorm:"column:name" json:"name"`
	Slug        string    `gorm:"column:slug" json:"slug"`
	Description string    `gorm:"column:description" json:"description"`
	ParentID    *uint     `gorm:"column:parent_id" json:"parentId"`
	Parent      *Category `json:"parent"`
}

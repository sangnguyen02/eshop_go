package models

// Category represents a product category
type Category struct {
	BaseModel
	Name        string    `gorm:"column:name;index" json:"name"`
	Slug        string    `gorm:"column:slug;type:VARCHAR(255);uniqueIndex" json:"slug"` // Giới hạn độ dài
	Description string    `gorm:"column:description" json:"description"`
	ParentID    *uint     `gorm:"column:parent_id;index" json:"parentId"`
	Parent      *Category `json:"parent"`
}

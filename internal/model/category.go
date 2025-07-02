package model

// import "gorm.io/gorm"

type Category struct {
	BaseModel
	Name        string     `gorm:"column:name;index" json:"name"`
	Slug        string     `gorm:"column:slug;type:VARCHAR(255);uniqueIndex" json:"slug"`
	Description string     `gorm:"column:description" json:"description"`
	Image       string     `gorm:"column:image" json:"image"`
	ParentID    *uint      `gorm:"column:parent_id;index" json:"parentId"`
	Parent      *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children    []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Status      bool       `gorm:"column:status" json:"status"`
}

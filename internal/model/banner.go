package model

type Banner struct {
	BaseModel
	Name   string `gorm:"column:name;index" json:"name"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

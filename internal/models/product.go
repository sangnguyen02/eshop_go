package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID     int    `gorm:"primaryKey"`
	Name          string `gorm:"type:nvarchar(50)"`
	ProductNumber string `gorm:"type:nvarchar(25)"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel định nghĩa các trường chung như ID, CreatedAt, UpdatedAt, DeletedAt
type BaseModel struct {
	ID        uint           `json:"id" example:"1" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" example:"2025-06-24T15:04:05Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2025-06-24T15:04:05Z"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

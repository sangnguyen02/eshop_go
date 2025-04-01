package database

import (
	"fmt"
	"go_ecommerce/pkg/setting"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Tạo chuỗi kết nối từ DatabaseSetting
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)

	// Kết nối đến database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	return db
}

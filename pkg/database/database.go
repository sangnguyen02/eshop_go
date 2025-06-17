package database

import (
	"fmt"
	"go_ecommerce/internal/models"
	"go_ecommerce/pkg/setting"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

// initializes the database instance and performs migration
func InitDB() *gorm.DB {

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)

	var err error
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("Failed to connect to database (db): " + err.Error())
	}

	// migration
	err = db.AutoMigrate(
		&models.User{},
		&models.UserCredentials{},
		&models.Product{},
		&models.Category{},
		&models.Brand{},
		&models.ProductImage{},
		&models.ProductVariant{},
		&models.ProductReview{},
	)
	if err != nil {
		panic("Failed to migrate database (db): " + err.Error())
	}
	fmt.Println("Migration completed for primary database (db)")
	return db
}

func GetDB() *gorm.DB {
	return db
}

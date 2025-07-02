package database

import (
	"fmt"
	"go_ecommerce/internal/model"
	"go_ecommerce/pkg/setting"
	"log"
	"strconv"

	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// initializes the database instance and performs migration
func InitDB() *gorm.DB {

	/* SQL Server DSN format
	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Name,
	// )

	// var err error
	// db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
	// 	PrepareStmt: true,
	// })
	// if err != nil {
	// 	panic("Failed to connect to database (db): " + err.Error())
	// }
	*/

	/* PostgreSQL DSN format */

	portStr := setting.DatabaseSetting.Port
	port, err1 := strconv.Atoi(portStr)
	if err1 != nil {
		log.Fatalf("Invalid DB_PORT: %v", err1)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		port,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("Failed to connect to PostgreSQL database: " + err.Error())
	}

	// migration
	err = db.AutoMigrate(
		&model.User{},
		&model.UserCredentials{},
		&model.Product{},
		&model.Category{},
		&model.Brand{},
		&model.ProductImage{},
		&model.ProductVariant{},
		&model.ProductReview{},
		&model.Banner{},
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

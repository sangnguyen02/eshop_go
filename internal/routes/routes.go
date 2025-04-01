package routes

import (
	"go_ecommerce/internal/handlers"
	"go_ecommerce/internal/repository"
	"go_ecommerce/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Khởi tạo repository
	productRepo := repository.NewProductRepository(db)

	// Khởi tạo service
	productService := services.NewProductService(productRepo)

	// Khởi tạo handler
	productHandler := handlers.NewProductHandler(productService)

	// Định nghĩa route
	r.GET("/products/search", productHandler.SearchProducts)
}

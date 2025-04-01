package routes

import (
	"go_ecommerce/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	productHandler := handlers.NewProductHandler()
	categoryHandler := handlers.NewCategoryHandler()
	brandHandler := handlers.NewBrandHandler()
	seedHandler := handlers.NewSeedHandler() // sample data

	// #region API V1
	apiV1 := r.Group("/api/v1")

	// Product
	products := apiV1.Group("/products")
	products.GET("/search", productHandler.SearchProducts)
	products.GET("/:id", productHandler.GetProductByID)
	products.POST("", productHandler.CreateProduct)
	products.PUT("/:id", productHandler.UpdateProduct)
	products.DELETE("/:id", productHandler.DeleteProduct)

	// Category
	categories := apiV1.Group("/categories")
	categories.GET("/", categoryHandler.GetAllCategories)
	categories.GET("/:id", categoryHandler.GetCategoryByID)
	categories.POST("", categoryHandler.CreateCategory)
	categories.PUT("/:id", categoryHandler.UpdateCategory)
	categories.DELETE("/:id", categoryHandler.DeleteCategory)

	// Brand
	brands := apiV1.Group("/brands")
	brands.GET("/", brandHandler.GetAllBrands)
	brands.GET("/:id", brandHandler.GetBrandByID)
	brands.POST("", brandHandler.CreateBrand)
	brands.PUT("/:id", brandHandler.UpdateBrand)
	brands.DELETE("/:id", brandHandler.DeleteBrand)

	// Seed
	apiV1.POST("/seed", seedHandler.SeedData)

	// #region API V2
}

package routes

import (
	"go_ecommerce/internal/handlers"
	"go_ecommerce/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	userHandler := handlers.NewUserHandler()
	productHandler := handlers.NewProductHandler()
	categoryHandler := handlers.NewCategoryHandler()
	brandHandler := handlers.NewBrandHandler()
	// seedHandler := handlers.NewSeedHandler() // sample data

	// #region API V1
	apiV1 := r.Group("/api/v1")
	apiV1.POST("/signup", userHandler.SignUp)
	// dang nhap
	apiV1.POST("/signin", userHandler.SignIn)

	// #region protect route

	protected := apiV1.Group("/")

	protected.Use(middleware.AuthMiddleware())
	{

		// #region user route
		users := protected.Group("/users")
		{
			users.GET("/search", userHandler.SearchUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.PUT("/pass/:id", userHandler.UpdatePassword)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// #region product route
		products := protected.Group("/products")
		{
			products.GET("/search", productHandler.SearchProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		// #region category route
		categories := protected.Group("/categories")
		{
			categories.GET("/", categoryHandler.GetAllCategories)
			categories.GET("/:id", categoryHandler.GetCategoryByID)
			categories.POST("", categoryHandler.CreateCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// #region brand route
		brands := protected.Group("/brands")
		{
			brands.GET("/", brandHandler.GetAllBrands)
			brands.GET("/:id", brandHandler.GetBrandByID)
			brands.POST("", brandHandler.CreateBrand)
			brands.PUT("/:id", brandHandler.UpdateBrand)
			brands.DELETE("/:id", brandHandler.DeleteBrand)
		}

	}
	// User
	// users := apiV1.Group("/users")
	// users.GET("/search", userHandler.SearchUsers)
	// users.GET("/:id", userHandler.GetUserByID)
	// users.PUT("/:id", userHandler.UpdateUser)
	// users.PUT("/pass/:id", userHandler.UpdatePassword)
	// users.DELETE("/:id", userHandler.DeleteUser)

	// Product
	// products := apiV1.Group("/products")
	// products.GET("/search", productHandler.SearchProducts)
	// products.GET("/:id", productHandler.GetProductByID)
	// products.POST("", productHandler.CreateProduct)
	// products.PUT("/:id", productHandler.UpdateProduct)
	// products.DELETE("/:id", productHandler.DeleteProduct)

	// Category
	// categories := apiV1.Group("/categories")
	// categories.GET("/", categoryHandler.GetAllCategories)
	// categories.GET("/:id", categoryHandler.GetCategoryByID)
	// categories.POST("", categoryHandler.CreateCategory)
	// categories.PUT("/:id", categoryHandler.UpdateCategory)
	// categories.DELETE("/:id", categoryHandler.DeleteCategory)

	// Brand
	// brands := apiV1.Group("/brands")
	// brands.GET("/", brandHandler.GetAllBrands)
	// brands.GET("/:id", brandHandler.GetBrandByID)
	// brands.POST("", brandHandler.CreateBrand)
	// brands.PUT("/:id", brandHandler.UpdateBrand)
	// brands.DELETE("/:id", brandHandler.DeleteBrand)

	// Seed
	// apiV1.POST("/seed", seedHandler.SeedData)

}

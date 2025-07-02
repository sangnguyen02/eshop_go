package routes

import (
	"go_ecommerce/internal/handler"
	"go_ecommerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// #region init hanlder
	uploadHandler := handler.NewUploadHandler()
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	categoryHandler := handler.NewCategoryHandler()
	brandHandler := handler.NewBrandHandler()
	bannerHandler := handler.NewBannerHandler()
	// seedHandler := handler.NewSeedHandler() // sample data

	// #region API V1
	apiV1 := r.Group("/api/v1")
	apiV1.POST("/signup", userHandler.SignUp)
	// dang nhap
	apiV1.POST("/signin", userHandler.SignIn)

	// #region upload route
	upload := apiV1.Group("/upload")
	{
		upload.POST("", uploadHandler.UploadFileSingle)
		upload.POST("/multiple", uploadHandler.UploadFileMultiple)
	}

	r.Static("/api/v1/upload/files/images", "./upload/files/images")

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
			products.POST("/", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		// #region category route
		categories := protected.Group("/categories")
		{
			categories.GET("/", categoryHandler.GetAllCategories)
			categories.GET("/:id", categoryHandler.GetCategoryByID)
			categories.POST("/", categoryHandler.CreateCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// #region brand route
		brands := protected.Group("/brands")
		{
			brands.GET("/", brandHandler.GetAllBrands)
			brands.GET("/:id", brandHandler.GetBrandByID)
			brands.POST("/", brandHandler.CreateBrand)
			brands.PUT("/:id", brandHandler.UpdateBrand)
			brands.DELETE("/:id", brandHandler.DeleteBrand)
		}

		banners := protected.Group("/banners")
		{
			banners.GET("/", bannerHandler.GetAllBanners)
			banners.GET("/:id", bannerHandler.GetBannerByID)
			banners.POST("/", bannerHandler.CreateBanner)
			banners.PUT("/:id", bannerHandler.UpdateBanner)
			banners.DELETE("/:id", bannerHandler.DeleteBanner)
		}

	}

	// Seed
	// apiV1.POST("/seed", seedHandler.SeedData)

}

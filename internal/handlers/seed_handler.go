package handlers

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeedHandler struct {
	categoryRepo *repository.CategoryRepository
	brandRepo    *repository.BrandRepository
	productRepo  *repository.ProductRepository
}

func NewSeedHandler() *SeedHandler {
	return &SeedHandler{
		categoryRepo: repository.NewCategoryRepository(),
		brandRepo:    repository.NewBrandRepository(),
		productRepo:  repository.NewProductRepository(),
	}
}

func (h *SeedHandler) SeedData(c *gin.Context) {

	// Tạo category Electronics trước
	electronics := models.Category{
		Name:        "Electronics",
		Slug:        "electronics",
		Description: "Category for electronic products",
	}
	if err := h.categoryRepo.Create(&electronics); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed electronics category: " + err.Error()})
		return
	}

	// Tạo các category còn lại
	categories := []models.Category{
		{
			Name:        "Phones",
			Slug:        "phones",
			Description: "Category for mobile phones",
			ParentID:    &electronics.ID, // Gán ParentID là ID của Electronics
		},
		{
			Name:        "Laptops",
			Slug:        "laptops",
			Description: "Category for laptops",
			ParentID:    &electronics.ID, // Gán ParentID là ID của Electronics
		},
		{
			Name:        "Clothing",
			Slug:        "clothing",
			Description: "Category for clothing products",
		},
	}

	// Thêm các category còn lại vào database
	for i := range categories {
		if err := h.categoryRepo.Create(&categories[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed categories: " + err.Error()})
			return
		}
	}

	// Thêm dữ liệu mẫu cho brands
	brands := []models.Brand{
		{
			Name:        "Apple",
			Slug:        "apple",
			Description: "Apple Inc.",
			LogoURL:     "https://example.com/apple-logo.png",
		},
		{
			Name:        "Samsung",
			Slug:        "samsung",
			Description: "Samsung Electronics",
			LogoURL:     "https://example.com/samsung-logo.png",
		},
		{
			Name:        "Dell",
			Slug:        "dell",
			Description: "Dell Technologies",
			LogoURL:     "https://example.com/dell-logo.png",
		},
		{
			Name:        "Nike",
			Slug:        "nike",
			Description: "Nike, Inc.",
			LogoURL:     "https://example.com/nike-logo.png",
		},
	}

	// Thêm brands vào database
	for i := range brands {
		if err := h.brandRepo.Create(&brands[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed brands: " + err.Error()})
			return
		}
	}

	// Thêm dữ liệu mẫu cho products
	products := []models.Product{
		{
			SKU:              "IPHONE14PRO",
			Name:             "iPhone 14 Pro",
			Slug:             "iphone-14-pro",
			Description:      "The latest iPhone 14 Pro with A16 Bionic chip.",
			ShortDescription: "iPhone 14 Pro, 128GB, Space Black",
			Price:            999.99,
			DiscountPrice:    899.99,
			StockQuantity:    100,
			CategoryID:       categories[0].ID, // Phones
			BrandID:          brands[0].ID,     // Apple
			Status:           models.ProductStatusActive,
			Images: []models.ProductImage{
				{
					URL:       "https://example.com/iphone14pro.jpg",
					IsPrimary: true,
				},
			},
		},
		{
			SKU:              "GALAXYS23",
			Name:             "Samsung Galaxy S23",
			Slug:             "galaxy-s23",
			Description:      "The latest Samsung Galaxy S23 with Snapdragon 8 Gen 2.",
			ShortDescription: "Galaxy S23, 256GB, Phantom Black",
			Price:            799.99,
			DiscountPrice:    749.99,
			StockQuantity:    150,
			CategoryID:       categories[0].ID, // Phones
			BrandID:          brands[1].ID,     // Samsung
			Status:           models.ProductStatusActive,
			Images: []models.ProductImage{
				{
					URL:       "https://example.com/galaxys23.jpg",
					IsPrimary: true,
				},
			},
		},
		{
			SKU:              "DELLXPS13",
			Name:             "Dell XPS 13",
			Slug:             "dell-xps-13",
			Description:      "The latest Dell XPS 13 with 11th Gen Intel processor.",
			ShortDescription: "Dell XPS 13, 512GB SSD, Platinum Silver",
			Price:            1299.99,
			DiscountPrice:    1199.99,
			StockQuantity:    50,
			CategoryID:       categories[1].ID, // Laptops
			BrandID:          brands[2].ID,     // Dell
			Status:           models.ProductStatusActive,
			Images: []models.ProductImage{
				{
					URL:       "https://example.com/dellxps13.jpg",
					IsPrimary: true,
				},
			},
		},
		{
			SKU:              "NIKEAIRMAX",
			Name:             "Nike Air Max",
			Slug:             "nike-air-max",
			Description:      "The latest Nike Air Max sneakers.",
			ShortDescription: "Nike Air Max, Size 10, Black",
			Price:            149.99,
			DiscountPrice:    129.99,
			StockQuantity:    200,
			CategoryID:       categories[2].ID, // Clothing
			BrandID:          brands[3].ID,     // Nike
			Status:           models.ProductStatusActive,
			Images: []models.ProductImage{
				{
					URL:       "https://example.com/nikeairmax.jpg",
					IsPrimary: true,
				},
			},
		},
	}

	// Thêm products vào database
	for i := range products {
		if err := h.productRepo.Create(&products[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed products: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sample data seeded successfully"})
}

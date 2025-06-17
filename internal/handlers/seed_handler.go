package handlers

import (
	"fmt"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repository"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

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

var (
	categoryNames    = []string{"Badminton Rackets", "Shuttlecocks", "Badminton Shoes", "Badminton Apparel", "Badminton Accessories"}
	subCategoryNames = map[string][]string{
		"Badminton Rackets":     {"Offensive Rackets", "Defensive Rackets", "All-Around Rackets"},
		"Shuttlecocks":          {"Feather Shuttlecocks", "Nylon Shuttlecocks"},
		"Badminton Shoes":       {"Men's Shoes", "Women's Shoes", "Unisex Shoes"},
		"Badminton Apparel":     {"Shirts", "Shorts", "Skirts"},
		"Badminton Accessories": {"Grips", "Bags", "Strings"},
	}
	brandNames   = []string{"Yonex", "Victor", "Li-Ning", "Carlton", "Ashaway", "FZ Forza", "Babolat", "Apacs", "Wilson", "Karakal"}
	productTypes = []string{"Racket", "Shuttlecock", "Shoes", "Shirt", "Shorts", "Skirt", "Grip", "Bag", "String"}
	colors       = []string{"Black", "White", "Blue", "Red", "Green", "Yellow", "Orange"}
	sizes        = []string{"XS", "S", "M", "L", "XL"}
)

func (h *SeedHandler) SeedData(c *gin.Context) {
	const (
		numCategories = 5
		numBrands     = 10
		numProducts   = 1000
		batchSize     = 100
	)

	rand.Seed(time.Now().UnixNano())
	tx := h.categoryRepo.GetDB().Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction: " + tx.Error.Error()})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "panic occurred: " + fmt.Sprint(r)})
		}
	}()

	// Tạo categories (cha và con)
	var categories []models.Category
	categoryMap := make(map[string]uint) // Lưu ID của danh mục cha

	// Tạo danh mục cha
	for i := 0; i < numCategories; i++ {
		categoryName := categoryNames[i]
		slug := strings.ToLower(strings.ReplaceAll(categoryName, " ", "-"))

		categories = append(categories, models.Category{
			Name:        categoryName,
			Slug:        slug,
			Description: "Category for " + categoryName,
		})
	}
	if err := tx.CreateInBatches(categories, batchSize).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed categories: " + err.Error()})
		return
	}

	// Lưu ID của danh mục cha
	for _, cat := range categories {
		categoryMap[cat.Name] = cat.ID
	}

	// Tạo danh mục con
	var subCategories []models.Category
	for parentName, subs := range subCategoryNames {
		parentID := categoryMap[parentName]
		for _, subName := range subs {
			slug := strings.ToLower(strings.ReplaceAll(subName, " ", "-"))
			subCategories = append(subCategories, models.Category{
				Name:        subName,
				Slug:        slug,
				Description: "Subcategory for " + subName,
				ParentID:    &parentID,
			})
		}
	}
	if err := tx.CreateInBatches(subCategories, batchSize).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed subcategories: " + err.Error()})
		return
	}

	// Lấy danh sách category IDs thực tế
	var categoryIDs []uint
	if err := tx.Model(&models.Category{}).Select("id").Find(&categoryIDs).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch category IDs: " + err.Error()})
		return
	}

	// Tạo brands
	var brands []models.Brand
	for i := 0; i < numBrands; i++ {
		brandName := brandNames[i]
		slug := strings.ToLower(brandName)

		brands = append(brands, models.Brand{
			Name:        brandName,
			Slug:        slug,
			Description: "Description for " + brandName,
			LogoURL:     "https://example.com/" + slug + "-logo.png",
		})
	}
	if err := tx.CreateInBatches(brands, batchSize).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed brands: " + err.Error()})
		return
	}

	// Lấy danh sách brand IDs thực tế
	var brandIDs []uint
	if err := tx.Model(&models.Brand{}).Select("id").Find(&brandIDs).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch brand IDs: " + err.Error()})
		return
	}

	// Tạo products
	var products []models.Product
	for i := 1; i <= numProducts; i++ {
		// Chọn CategoryID và BrandID từ danh sách thực tế
		categoryID := categoryIDs[rand.Intn(len(categoryIDs))]
		brandID := brandIDs[rand.Intn(len(brandIDs))]
		productType := productTypes[rand.Intn(len(productTypes))]
		brandName := brandNames[(int(brandID)-1)%len(brandNames)]
		productName := brandName + " " + productType
		slug := strings.ToLower(strings.ReplaceAll(productName, " ", "-")) + "-" + strconv.Itoa(i)

		// Tạo variants
		numVariants := rand.Intn(3) + 1
		variants := make([]models.ProductVariant, numVariants)
		for j := 0; j < numVariants; j++ {
			color := colors[rand.Intn(len(colors))]
			size := sizes[rand.Intn(len(sizes))]
			variants[j] = models.ProductVariant{
				SKU:           "SKU" + strconv.Itoa(i) + "-" + strconv.Itoa(j),
				Name:          color + " " + size,
				Price:         10.00 + float64(rand.Intn(150)),
				StockQuantity: 5 + rand.Intn(50),
				Attributes:    fmt.Sprintf(`{"color": "%s", "size": "%s"}`, color, size),
			}
		}

		// Tạo images
		images := []models.ProductImage{
			{URL: "https://example.com/" + slug + "-1.jpg", IsPrimary: true},
			{URL: "https://example.com/" + slug + "-2.jpg", IsPrimary: false},
		}

		// Tạo reviews
		numReviews := rand.Intn(5)
		reviews := make([]models.ProductReview, numReviews)
		for j := 0; j < numReviews; j++ {
			reviews[j] = models.ProductReview{
				UserID:  uint(rand.Intn(100) + 1),
				Rating:  rand.Intn(5) + 1,
				Comment: "Review " + strconv.Itoa(j+1) + " for " + productName,
			}
		}

		products = append(products, models.Product{
			SKU:              "SKU" + strconv.Itoa(i),
			Name:             productName,
			Slug:             slug,
			Description:      "Description for " + productName,
			ShortDescription: "Short desc for " + productName,
			Price:            15.00 + float64(rand.Intn(200)),
			DiscountPrice:    10.00 + float64(rand.Intn(150)),
			StockQuantity:    10 + rand.Intn(100),
			CategoryID:       categoryID,
			BrandID:          brandID,
			Status:           models.ProductStatusActive,
			Images:           images,
			Variants:         variants,
			Reviews:          reviews,
		})

		if len(products) == batchSize || i == numProducts {
			if err := tx.CreateInBatches(products, batchSize).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to seed products: " + err.Error()})
				return
			}
			products = nil
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Badminton sample data seeded successfully"})
}

package handler

import (
	"go_ecommerce/internal/dto"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/service"
	"go_ecommerce/internal/utils/formatter"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{service: service.ProductServiceFactory()}
}

// SearchProducts searches products by name with pagination
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	var req dto.SearchRequestParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": formatter.FormatValidationErrors(err)})
		return
	}
	products, total, err := h.service.SearchProducts(req.Name, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     products,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// GetProductByID gets a product by ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	var req dto.IDRequestParams
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID: " + err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": formatter.FormatValidationErrors(err)})
		return
	}

	product, err := h.service.GetProductByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates an existing product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var product model.Product // Sửa từ service.Product thành models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	if err := h.service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

// #region customiztion

func (h *ProductHandler) SearchForCard(c *gin.Context) {
	var req dto.SearchRequestParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": formatter.FormatValidationErrors(err)})
		return
	}

	products, total, err := h.service.SearchForCard(req.Name, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products: " + err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":     []dto.ProductCardResponse{},
			"total":    total,
			"page":     req.Page,
			"pageSize": req.PageSize,
			"message":  "No products found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     products,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

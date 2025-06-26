package handler

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	service *service.BrandService
}

func NewBrandHandler() *BrandHandler {
	return &BrandHandler{service: service.BrandServiceFactory()}
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var brand model.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateBrand(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand)
}

func (h *BrandHandler) GetBrandByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	brand, err := h.service.GetBrandByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "brand not found"})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func (h *BrandHandler) GetAllBrands(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	search := c.Query("name")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	brands, total, err := h.service.GetAllBrands(page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       brands,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var updatedBrand model.Brand
	if err := c.ShouldBindJSON(&updatedBrand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateBrand(uint(id), &updatedBrand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand updated"})
}

func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.DeleteBrand(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand deleted"})
}

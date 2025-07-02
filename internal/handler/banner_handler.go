package handler

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	service *service.BannerService
}

func NewBannerHandler() *BannerHandler {
	return &BannerHandler{service: service.BannerServiceFactory()}
}

func (h *BannerHandler) CreateBanner(c *gin.Context) {
	var banner model.Banner
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateBanner(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, banner)
}

func (h *BannerHandler) GetBannerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	banner, err := h.service.GetBannerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "banner not found"})
		return
	}

	c.JSON(http.StatusOK, banner)
}

func (h *BannerHandler) GetAllBanners(c *gin.Context) {
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

	statusStr := c.Query("status")
	var status bool
	if statusStr == "" {
		status = true
	} else {
		var err error
		status, err = strconv.ParseBool(statusStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid status value, must be true or false",
			})
			return
		}
	}

	banners, total, err := h.service.GetAllBanners(page, pageSize, search, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       banners,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var updatedBanner model.Banner
	if err := c.ShouldBindJSON(&updatedBanner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateBanner(uint(id), &updatedBanner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "banner updated"})
}

func (h *BannerHandler) DeleteBanner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.DeleteBanner(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "banner deleted"})
}

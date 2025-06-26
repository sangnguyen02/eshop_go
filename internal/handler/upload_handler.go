package handler

import (
	"fmt"
	utils "go_ecommerce/internal/utils/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) UploadFileSingle(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "Lỗi",
		})
		return
	}

	if file == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "INVALID_PARAMS",
		})
		return
	}

	fileName := utils.GetImageName(file.Filename)
	fullPath := utils.GetImageFullPath()
	savePath := utils.GetImagePath()
	src := fullPath + fileName

	// Check image validity
	if err := utils.CheckImage(fullPath); err != nil {
		fmt.Println("CheckImage error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FAIL",
		})
		return
	}

	// Save file
	if err := c.SaveUploadedFile(file, src); err != nil {
		fmt.Println("Error saving file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "ERROR_UPLOAD_SAVE_IMAGE_FAIL",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":         false,
		"msg":           "SUCCESS",
		"file_url":      utils.GetImageFullUrl(fileName),
		"file_save_url": savePath + fileName,
	})
}

func (h *UploadHandler) UploadFileMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"error": true,
			"msg":   "INVALID_PARAMS",
		})
		return
	}

	fullPath := utils.GetImageFullPath()
	savePath := utils.GetImagePath()

	for _, file := range files {
		fileName := utils.GetImageName(file.Filename)
		src := fullPath + fileName

		if !utils.CheckFileExt(fileName) || !utils.CheckFileSize(file) {
			fmt.Println(utils.CheckFileExt(fileName))
			fmt.Println(utils.CheckFileSize(file))
			c.JSON(http.StatusOK, gin.H{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_FILE_FORMAT",
			})
			return
		}

		if err := utils.CheckImage(fullPath); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_FILE_FAIL",
			})
			return
		}

		if err := c.SaveUploadedFile(file, src); err != nil {
			fmt.Println("Error saving file:", err)
			c.JSON(http.StatusOK, gin.H{
				"error": true,
				"msg":   "ERROR_UPLOAD_SAVE_FILE_FAIL",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error":         false,
		"msg":           "SUCCESS",
		"file_url":      utils.GetImagePath(),
		"file_save_url": savePath,
	})
}

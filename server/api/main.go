package main

import (
	"fmt"
	"go_ecommerce/internal/routes"
	"go_ecommerce/pkg/database"
	"go_ecommerce/pkg/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo cấu hình
	setting.Setup()

	// Khởi tạo kết nối database
	db := database.InitDB()

	// Khởi tạo Gin router
	r := gin.Default()

	// Đăng ký các route
	routes.RegisterRoutes(r, db)

	// Chạy server
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	fmt.Printf("Starting server on port %s\n", port)
	if err := r.Run(port); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

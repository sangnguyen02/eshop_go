package main

import (
	"fmt"
	"go_ecommerce/internal/routes"
	"go_ecommerce/pkg/database"
	"go_ecommerce/pkg/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	// init setting
	setting.Setup()

	// init db conn
	_ = database.InitDB()

	// Khởi tạo Gin router
	r := gin.Default()

	// Đăng ký các route
	routes.RegisterRoutes(r)

	// Chạy server
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	fmt.Printf("Starting server on port %s\n", port)
	if err := r.Run(port); err != nil {
		fmt.Println("Error starting server:", err)
	}

}

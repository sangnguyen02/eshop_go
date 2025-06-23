package main

import (
	"fmt"
	"go_ecommerce/internal/routes"
	"go_ecommerce/pkg/database"
	"go_ecommerce/pkg/setting"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// init setting
	setting.Setup()

	// init db connection
	_ = database.InitDB()

	// Khởi tạo Gin router
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// Global middleware
	r.Use(gin.Logger())                       // Log request
	r.Use(gin.Recovery())                     // Recover panic
	r.Use(cors.Default())                     // CORS
	r.Use(gzip.Gzip(gzip.DefaultCompression)) // Nén response
	r.Use(requestid.New())                    // Gắn ID cho mỗi request

	// Đăng ký các route
	routes.RegisterRoutes(r)

	// Chạy server
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	fmt.Printf("Starting server on port %s\n", port)
	if err := r.Run(port); err != nil {
		fmt.Println("Error starting server:", err)
	}

}

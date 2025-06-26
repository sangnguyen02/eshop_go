package main

import (
	"fmt"
	"go_ecommerce/internal/middleware"
	"go_ecommerce/internal/routes"
	"go_ecommerce/pkg/database"
	"go_ecommerce/pkg/setting"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_ecommerce/docs"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// init setting
	setting.Setup()

	// init db connection
	_ = database.InitDB()

	// init Gin
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// global middleware
	r.Use(middleware.InitMiddlewares()...)

	// init route
	routes.RegisterRoutes(r)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run server
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	fmt.Printf("Starting server on port %s\n", port)
	if err := r.Run(port); err != nil {
		fmt.Println("Error starting server:", err)
	}

}

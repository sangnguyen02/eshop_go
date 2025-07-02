package middleware

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),   // Log requests
		gin.Recovery(), // Recover from panics
		corsMiddleware(),
		gzip.Gzip(gzip.DefaultCompression), // Compress responses
		requestid.New(),                    // Generate request ID
		// cacheControlForStatic(),
	}
}

func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	})
}

func cacheControlForStatic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// chỉ áp dụng cho route tĩnh
		if strings.HasPrefix(c.Request.URL.Path, "/api/v1/upload/files/images") {
			c.Header("Cache-Control", "max-age=3600") // 1 tiếng
		}
		c.Next()
	}
}

package middleware

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// InitMiddlewares initializes and returns a slice of Gin middleware
func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),                       // Log requests
		gin.Recovery(),                     // Recover from panics
		cors.Default(),                     // CORS with default config
		gzip.Gzip(gzip.DefaultCompression), // Compress responses
		requestid.New(),                    // Generate request ID
		cacheControlForStatic(),            // Thêm middleware cho Cache-Control
	}
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

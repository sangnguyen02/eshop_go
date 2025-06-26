package middleware

import (
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
	}
}

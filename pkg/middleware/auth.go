package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Thêm logic authentication nếu cần
		c.Next()
	}
}

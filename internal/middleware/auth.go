package middleware

import (
	"go_ecommerce/internal/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
			return
		}

		// get tokenStr from prefix "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// verify token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// check token (expired)
		exp, ok := (*claims)["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		// get user data
		userInfo, err := utils.ExtractUserFromClaims(claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		// save to context
		c.Set("userID", userInfo.ID)
		c.Set("username", userInfo.Username)
		c.Set("role", userInfo.Role)

		c.Next()
	}
}

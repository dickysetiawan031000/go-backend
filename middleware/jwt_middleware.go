package middleware

import (
	"net/http"
	"strings"

	"github.com/dickysetiawan031000/go-backend/utils"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Cek apakah token valid (masih ada di validTokens)
		if !utils.IsTokenValid(tokenStr) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token has been invalidated"})
			return
		}

		claims, err := utils.VerifyJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Set user ID ke context
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

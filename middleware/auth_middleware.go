package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dickysetiawan031000/go-backend/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MASUK MIDDLEWARE")
		authHeader := c.GetHeader("Authorization")
		fmt.Println("Auth Header:", authHeader)

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("Token:", tokenStr)

		claims, err := utils.VerifyJWT(tokenStr)
		if err != nil {
			fmt.Println("JWT verification error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		fmt.Printf("Claims: %+v\n", claims)

		// simpan user_id di context untuk handler
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

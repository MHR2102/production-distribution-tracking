package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User role tidak ditemukan"})
			c.Abort()
			return
		}

		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Anda tidak memiliki akses ke resource ini"})
			c.Abort()
			return
		}

		c.Next()
	}
}

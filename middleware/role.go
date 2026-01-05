package middleware

import (
	"Be-Book-Padel/helper"

	"github.com/gin-gonic/gin"
)

func AdminOrStaffOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtData, exists := c.Get("Auth")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		data := jwtData.(helper.JwtData)
		if data.Role != "admin" && data.Role != "staff" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins and Staff only"})
			return
		}

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtData, exists := c.Get("Auth")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		data := jwtData.(helper.JwtData)
		if data.Role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins only"})
			return
		}

		c.Next()
	}
}

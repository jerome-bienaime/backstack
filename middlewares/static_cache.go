package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func StaticCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Apply the Cache-Control header to the static files
		if strings.HasPrefix(c.Request.URL.Path, "/public/") {
			c.Header("Cache-Control", "private, max-age=86400")
		}
		// Continue to the next middleware or handler
		c.Next()
	}
}

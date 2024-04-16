package config

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Csrf(secretKey string) gin.HandlerFunc {

	return csrf.Middleware(csrf.Options{
		Secret: secretKey,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	})
}

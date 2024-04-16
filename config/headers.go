package config

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

func Headers() (gin.HandlerFunc, gin.HandlerFunc, gin.HandlerFunc, gin.HandlerFunc, gin.HandlerFunc, gin.HandlerFunc) {
	return helmet.Default()
}

package main

import (
	"net/http"
	"time"
	"web/config"
	"web/controllers"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB, e config.EnvConfig) *gin.Engine {
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	r.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/", func(c *gin.Context) {
		controllers.Index(c, db, e)
	})

	r.GET("/home", cache.CacheByRequestURI(memoryStore, 2*time.Second), func(c *gin.Context) {
		controllers.Index(c, db, e)
	})

	return r
}

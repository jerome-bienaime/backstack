package config

import (
	"time"

	"github.com/gin-gonic/gin"
	cachecontrol "github.com/joeig/gin-cachecontrol"
)

func Cache() gin.HandlerFunc {
	return cachecontrol.New(&cachecontrol.Config{
		MustRevalidate:       true,
		NoCache:              false,
		NoStore:              false,
		NoTransform:          false,
		Public:               true,
		Private:              false,
		ProxyRevalidate:      true,
		MaxAge:               cachecontrol.Duration(30 * time.Minute),
		SMaxAge:              nil,
		Immutable:            false,
		StaleWhileRevalidate: cachecontrol.Duration(2 * time.Hour),
		StaleIfError:         cachecontrol.Duration(2 * time.Hour),
	})
}

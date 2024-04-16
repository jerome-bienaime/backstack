package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func timeoutResponse(c *gin.Context) {
	e := SetupEnv()
	log.WithFields(log.Fields{
		"timeout": e.HTTP_TIMEOUT,
	}).Warn("Request timed out")
	c.SetCookie("error", "Request timed out", 360, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/home")
}

func Timeout(timeoutDuration time.Duration) gin.HandlerFunc {
	return timeout.New(timeout.WithTimeout(timeoutDuration),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

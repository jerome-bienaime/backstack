package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleError(c *gin.Context, err error, message string) {
	log.WithError(err).Warn(message)
	c.SetCookie("error", message, 360, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

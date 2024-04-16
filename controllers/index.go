package controllers

import (
	"net/http"
	"web/config"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

func Index(c *gin.Context, db *gorm.DB, e config.EnvConfig) {
	err_message, _ := c.Cookie("error")
	success_message, _ := c.Cookie("success")
	c.SetCookie("error", "", -1, "/", "localhost", false, true)
	c.SetCookie("success", "", -1, "/", "localhost", false, true)
	csrfToken := "1"
	if e.MODE != "TEST" {
		csrfToken = csrf.GetToken(c)
	}

	c.HTML(http.StatusOK, "pages/home.html", gin.H{
		"title":     "Home",
		"error":     err_message,
		"success":   success_message,
		"csrfToken": csrfToken,
	})
}

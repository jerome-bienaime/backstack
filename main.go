package main

import (
	"web/config"
	DB "web/db"
	"web/middlewares"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func boot() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	e := middlewares.SetupEnv()

	r.Use(config.Cors())
	r.Use(config.Headers())
	r.Use(middlewares.Timeout(config.RequestTimeout(e.HTTP_TIMEOUT)))

	if e.MODE == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	if e.HTTP_CACHE == "1" {
		r.Use(config.Cache())
		r.Use(middlewares.StaticCache())
	}

	if e.HTTP_COMPRESSION == "1" {
		r.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	store := cookie.NewStore([]byte(e.SECRET_KEY))
	r.Use(sessions.Sessions("jobseek_session", store))

	if e.MODE != "TEST" {
		r.Use(config.Csrf(e.SECRET_KEY))
	}

	r.LoadHTMLGlob("views/**/*")

	dbFile := middlewares.SetupEnv().SQLITE_PATH
	db := DB.Connect(dbFile, &gorm.Config{})
	DB.Migrate(db)

	return Routes(r, db, *e)
}

func main() {
	e := middlewares.SetupEnv()
	r := boot()

	if e.HTTP_SSL == "1" {
		r.RunTLS(":"+e.HTTP_PORT, "cert.pem", "cert.key")
	} else {
		r.Run(":" + e.HTTP_PORT)
	}
}

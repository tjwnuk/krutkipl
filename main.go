package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"krutki.pl/controllers"
	"krutki.pl/models"
)

var db *gorm.DB
var ct *controllers.Controller

func main() {

	db, err := models.GetDB()

	if err != nil {
		panic("Error connecting database")
	}

	db.AutoMigrate(&models.Url{})

	ct = &controllers.Controller{db}

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	// Assets
	router.Static("/assets", "./assets")
	// 404
	router.NoRoute(ct.Error404Handler)

	// Routes
	router.GET("/", ct.IndexHandler)
	router.GET("/login", ct.LoginHandler)
	router.POST("/login", ct.LoginPostHandler)
	router.POST("/shorten", ct.ShortenHandler)

	// Redirect
	router.GET("/l/:token", ct.RedirectHandler)

	router.Run()

}

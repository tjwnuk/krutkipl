package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"krutki.pl/controllers"
	"krutki.pl/middleware"
	"krutki.pl/models"
)

var db *gorm.DB
var ct *controllers.Controller

func main() {

	db, err := models.GetDB()

	if err != nil {
		panic("Error connecting database")
	}

	db.AutoMigrate(&models.Url{}, &models.User{})

	// ct - controller object
	ct = &controllers.Controller{db}

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	// Assets
	router.Static("/assets", "./assets")
	// 404
	router.NoRoute(ct.Error404Handler)

	// Routes
	router.GET("/", ct.IndexHandler)
	router.GET("/about", middleware.AlreadyLogged, ct.AboutHandler)
	router.GET("/login", ct.LoginHandler)
	router.POST("/login", ct.LoginPostHandler)
	router.POST("/shorten", ct.ShortenHandler)

	router.GET("/register", ct.RegisterControllerHandler)
	router.POST("/register", ct.RegisterControllerPOST)

	// Mod panel
	router.GET("/mod", middleware.RequireAuth, ct.ModPanelListAllLinks)

	// Redirect all other routes
	// Check if route matches token in DB, if yes, redirect
	router.GET("/:token", ct.RedirectHandler)

	router.Run()

}

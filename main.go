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

	gin.SetMode(gin.DebugMode)

	db, err := models.GetDB()

	if err != nil {
		panic("Error connecting database")
	}

	db.AutoMigrate(&models.Url{}, &models.User{})

	// ct - controller object
	ct = &controllers.Controller{Db: db}

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	// Assets
	router.Static("/assets", "./assets")
	// 404
	router.NoRoute(ct.Error404Handler)

	// Routes
	router.Use(middleware.AlreadyLogged)

	router.GET("/", ct.IndexHandler)
	router.GET("/about", ct.AboutHandler)
	router.GET("/login", ct.LoginHandler)
	router.POST("/login", ct.LoginPostHandler)
	router.GET("/logout", ct.LogoutHandler)
	router.POST("/shorten", ct.ShortenHandler)
	router.GET("/register", ct.RegisterControllerHandler)
	router.POST("/register", ct.RegisterControllerPOST)

	// Mod panel
	router.GET("/manage-links", middleware.RequireAuth, ct.ManageLinks)

	// Redirect all other routes
	// Check if route matches token in DB, if yes, redirect
	router.GET("/:token", ct.RedirectHandler)

	router.Run()

}

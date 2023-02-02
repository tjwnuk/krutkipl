package main

import (
	"github.com/gin-gonic/gin"
	"krutki.pl/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	// Assets
	router.Static("/assets", "./assets")
	// 404
	router.NoRoute(controllers.Error404Handler)

	// Routes
	router.GET("/", controllers.IndexHandler)
	router.GET("/login", controllers.LoginHandler)
	router.POST("/login", controllers.LoginPostHandler)
	router.POST("/shorten", controllers.ShortenHandler)

	// Redirect
	router.GET("/l/:token", controllers.RedirectHandler)

	router.Run()
}

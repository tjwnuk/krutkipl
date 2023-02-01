package main

import (
	"github.com/gin-gonic/gin"
	"krutki.pl/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	// Routes
	router.Static("/assets", "./assets")
	router.GET("/", controllers.IndexHandler)
	router.GET("/login", controllers.LoginHandler)
	router.POST("/login", controllers.LoginPostHandler)

	router.Run()
}

package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
	"krutki.pl/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	router.Static("/assets", "./assets")

	router.GET("/", controllers.IndexHandler)
	router.GET("/login", controllers.LoginHandler)

	router.Run()
}

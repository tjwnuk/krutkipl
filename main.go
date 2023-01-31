package main

import (
	"net/http"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "Strona główna",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.tmpl.html")

	router.Static("/assets", "./assets")

	router.GET("/", IndexHandler)

	router.Run()
}

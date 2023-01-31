package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", gin.H{
		"title": "Zaloguj się",
	})
}

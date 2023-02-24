package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ct Controller) IndexHandler(c *gin.Context) {

	currentUser, ok := c.Get("User")

	if !ok {
		currentUser = nil
	}

	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "krutki.pl - link shortener",
		"User":  currentUser,
	})
}

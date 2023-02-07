package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ct Controller) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "Strona główna",
	})
}

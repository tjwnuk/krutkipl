package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectHandler(c *gin.Context) {
	token := c.Param("token")

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})

}

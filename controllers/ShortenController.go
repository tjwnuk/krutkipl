package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	originalURL := c.Request.PostFormValue("originalURL")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"StatusCode": http.StatusInternalServerError,
			"Response":   "Internal Server Error",
			"URL":        originalURL,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"URL": originalURL,
	})
}

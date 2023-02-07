package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

// Handle request for shortening URL

func (ct Controller) ShortenHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	originalURL := c.Request.PostFormValue("originalURL")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"StatusCode": http.StatusInternalServerError,
			"Response":   "Internal Server Error",
			"URL":        originalURL,
		})
	}

	model := models.Model{Db: ct.Db}

	shortenedUrl, err := model.CreateNewShortcut(originalURL)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": "error creating new shortcut",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ShortenedUrl": shortenedUrl,
	})
}

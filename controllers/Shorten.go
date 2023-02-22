package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

// Handle request for shortening URL
func (ct Controller) ShortenHandler(c *gin.Context) {
	// get user
	var shortenedUrl string

	var userStruct *models.User

	currentUser, userPresent := c.Get("User")

	if userPresent {
		userStruct = c.Keys["User"].(*models.User)
	}

	if !userPresent {
		currentUser = nil
	}

	//parse form
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

	if !userPresent {
		shortenedUrl, err = model.CreateNewShortcut(originalURL)
	} else {
		shortenedUrl, err = model.CreateNewShortcutByUser(originalURL, int(userStruct.ID))
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": "error creating new shortcut",
		})
		return
	}

	// full link to copy
	baseUrl := "http://" + c.Request.Host + "/"

	// render result page
	c.HTML(http.StatusOK, "shorten/result", gin.H{
		"ShortenedUrl": baseUrl + shortenedUrl,
		"User":         currentUser,
	})
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

func (ct Controller) AboutHandler(c *gin.Context) {

	var userString string
	_, exists := c.Get("User")
	var user *models.User

	if exists {
		user = c.Keys["User"].(*models.User)
	}

	if exists {
		userString += "ID:       " + strconv.Itoa(int(user.ID)) + "; "
		userString += "Username: " + user.Username + "; "
		userString += "Rank:     " + user.Rank + "; "
	} else {
		userString = "not logged in"
	}

	c.HTML(http.StatusOK, "pages/about", gin.H{
		"User":       user,
		"UserString": userString,
		"title":      "About this site",
	})
}

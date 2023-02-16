package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

func (ct *Controller) RegisterControllerPOST(c *gin.Context) {

	model := models.Model{Db: ct.Db}

	err := c.Request.ParseForm()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "error parsing form",
		})
		return
	}

	username := c.Request.PostFormValue("username")
	password := c.Request.PostFormValue("password")
	password1 := c.Request.PostFormValue("password1")

	if password != password1 {
		c.JSON(http.StatusOK, gin.H{
			"error": "passwords do not match",
		})
		return
	}

	ok := model.CreateUser(username, password)

	if ok {
		// c.JSON(http.StatusOK, gin.H{
		// 	"status":   "ok",
		// 	"msg":      "created new user",
		// 	"username": username,
		// })

		c.HTML(http.StatusOK, "login/registerSuccess", gin.H{
			"Username": username,
		})

		return
	} else {
		c.HTML(http.StatusOK, "login/registerFail", gin.H{
			"Username": username,
		})
	}

}

func (ct *Controller) RegisterControllerHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login/register", nil)
}

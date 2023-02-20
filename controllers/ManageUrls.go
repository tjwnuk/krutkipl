package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

func (ct Controller) ManageLinks(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "user error",
		})

		return
	}

	model := models.Model{Db: ct.Db}

	links := model.GetAllLinks()

	c.HTML(http.StatusOK, "manage/manageLinks", gin.H{
		"User":  user,
		"Links": links,
	})
}

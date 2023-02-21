package controllers

import (
	"net/http"
	"strconv"

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

func (ct Controller) ManageLinksDeleteLink(c *gin.Context) {

	linkId, err := strconv.Atoi(c.Param("link_id"))

	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "error parsing param",
		})

		return
	}

	model := models.Model{Db: ct.Db}

	model.DeleteUrl(linkId)

	c.Redirect(302, "/manage-links")
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

// Account management for user
func (ct Controller) AccountPanel(c *gin.Context) {

	// return an error if user is not set in session
	if _, ok := c.Get("User"); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "not logged in",
		})

		return
	}

	var user *models.User = c.Keys["User"].(*models.User)

	model := models.Model{Db: ct.Db}
	urls, ok := model.GetAllUserLinks(int(user.ID))

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "error getting links",
		})
	}

	c.HTML(200, "user/accountPanel", gin.H{
		"Links": urls,
		"User":  user,
	})
}

func (ct Controller) AccountPanelDeleteLink(c *gin.Context) {

	linkId, err := strconv.Atoi(c.Param("link_id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "internal server error",
		})

		return
	}

	// return an error if user is not set in session
	if _, ok := c.Get("User"); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "not logged in",
		})

		return
	}

	var user *models.User = c.Keys["User"].(*models.User)

	ct.Db.Where("user_id = ?", user.ID).Delete(&models.Url{}, linkId)

	c.Redirect(302, "/account-panel")

}

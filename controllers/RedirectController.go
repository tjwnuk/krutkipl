package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

func (ct Controller) RedirectHandler(c *gin.Context) {
	token := c.Param("token")

	model := models.Model{Db: ct.Db}

	ok, url := model.GetRedirect(token)

	if ok {
		c.Redirect(302, url)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "nie ma takiej strony",
		})
	}
}

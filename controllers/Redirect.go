package controllers

import (
	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

func (ct Controller) RedirectHandler(c *gin.Context) {
	token := c.Param("token")

	model := models.Model{Db: ct.Db}

	ok, url := model.GetRedirectUrl(token)

	if ok {
		// if the token is present in the database
		c.Redirect(302, url)
	} else {
		// or if it's not
		c.HTML(404, "error404", nil)
	}
}

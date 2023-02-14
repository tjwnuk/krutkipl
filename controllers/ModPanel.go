package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ct Controller) ModPanelListAllLinks(c *gin.Context) {

	user, exist := c.Get("user")

	if exist {
		c.HTML(http.StatusOK, "modpanel/alllinks", gin.H{
			"User": user,
		})

	}

}

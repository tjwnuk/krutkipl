package controllers

import "github.com/gin-gonic/gin"

func (ct Controller) AboutHandler(c *gin.Context) {
	currentUser, exists := c.Get("User")

	if exists {
		c.JSON(200, currentUser)
	} else {
		c.JSON(200, gin.H{
			"msg": "you are not logged in",
		})
	}
}

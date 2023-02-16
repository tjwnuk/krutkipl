package controllers

import "github.com/gin-gonic/gin"

func (ct Controller) AboutHandler(c *gin.Context) {
	currentUser, exists := c.Get("user")

	if exists {
		c.JSON(200, currentUser)
	}
}

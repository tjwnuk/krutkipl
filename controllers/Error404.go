package controllers

import "github.com/gin-gonic/gin"

func (ct Controller) Error404Handler(c *gin.Context) {
	c.HTML(404, "error404", nil)
}

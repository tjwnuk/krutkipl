package controllers

import "github.com/gin-gonic/gin"

func Error404Handler(c *gin.Context) {
	c.HTML(404, "error404", nil)
}

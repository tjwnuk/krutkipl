package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Email    string
	Password string
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", gin.H{
		"title": "Zaloguj się",
	})
}

func LoginPostHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		panic(err)
	}

	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")

	userData := LoginData{Email: email, Password: password}

	c.HTML(http.StatusOK, "login/example", userData)
}

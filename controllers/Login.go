package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"krutki.pl/models"
)

type LoginData struct {
	Email    string
	Password string
}

func (ct Controller) LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", gin.H{
		"title": "Zaloguj się",
	})
}

func (ct Controller) LoginPostHandler(c *gin.Context) {

	_ = godotenv.Load()

	err := c.Request.ParseForm()
	if err != nil {
		panic(err)
	}

	username := c.Request.PostFormValue("username")
	password := c.Request.PostFormValue("password")

	// userData := LoginData{Email: email, Password: password}

	model := models.Model{Db: ct.Db}

	ok, user := model.GetUser(username, password)

	if ok {
		// if success, username and password matches

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"msg":    "internal server error",
			})

			fmt.Println(err)

			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

		return
	} else {
		// if username and password does not match
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    "user not found",
		})

		return
	}
}

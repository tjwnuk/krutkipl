package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"krutki.pl/models"
)

func RequireAdmin(c *gin.Context) {

	db, err := models.GetDB()

	if err != nil {
		panic("cannot connect to the db from middleware")
	}

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.HTML(401, "errors/permissionDenied", nil)
		c.Next()
		return
	}

	token, err := parseToken(tokenString)

	if err != nil {
		panic("error processing the token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user *models.User

		db.First(&user, claims["sub"])

		// refuse request from non admin and non mod users
		var notAllowed bool = true

		if user.ID == 0 {
			notAllowed = true
		}

		if user.Rank == "admin" {
			notAllowed = false
		}

		if notAllowed {
			fmt.Println("Error 401: unauthorized user tried to access protected resource")
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Next()
			return
		}

		c.Set("User", user)

		// fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
		return
	}

	c.Next()
}

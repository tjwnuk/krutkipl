package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"krutki.pl/models"
)

// This middleware updates context with currently logged user
func AlreadyLogged(c *gin.Context) {

	// c.Set("user", models.User{})

	if cookie, err := c.Request.Cookie("Authorization"); err == nil {
		value := cookie.Value

		var currentUser *models.User

		token, err := parseToken(value)

		if err == nil {
			var ok bool

			ok, currentUser = getUserFromCookie(token)

			if ok {
				// execute when the user is logged in and cookie is set

				c.Set("User", currentUser)
			}
		}

	}

	c.Next()
}

func getUserFromCookie(token *jwt.Token) (bool, *models.User) {

	db, err := models.GetDB()

	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	_ = ok

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return false, nil
	}

	var user *models.User

	db.First(&user, claims["sub"])

	if user.ID != 0 {
		return true, user
	}

	return false, nil

}

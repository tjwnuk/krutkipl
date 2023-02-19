package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Deletes cookie and logouts user
func (ct Controller) LogoutHandler(c *gin.Context) {

	// ResponseWriter
	rw := c.Writer

	// set blank cookie
	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: "",
		// Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	// Delete cookie
	http.SetCookie(rw, cookie)

	// redirect to the main page
	// baseUrl := "http://" + c.Request.Host + "/"
	// c.Redirect(301, baseUrl)

	c.HTML(200, "login/logout", nil)
}

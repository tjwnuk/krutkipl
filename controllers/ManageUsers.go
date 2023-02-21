package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"krutki.pl/models"
)

func (ct Controller) ManageUsers(c *gin.Context) {

	user, exist := c.Get("User")

	// check middleware error
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "user error",
		})

		return
	}

	model := models.Model{Db: ct.Db}

	usersList := model.GetAllUsers()

	c.HTML(http.StatusOK, "manage/manageUsers", gin.H{
		"User":      user,
		"UsersList": usersList,
	})
}

func (ct Controller) ManageUsersDeleteUser(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "error parsing param",
		})

		return
	}

	model := models.Model{Db: ct.Db}

	model.DeleteUser(userID)

	c.Redirect(302, "/manage-users")
}

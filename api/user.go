package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/takeanote/takeanote-api/model"
	"github.com/takeanote/takeanote-api/router/middleware"
	"github.com/takeanote/takeanote-api/store"
)

type inputUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// ProfileView return the information of the user
func ProfileView(c *gin.Context) {
	user := c.Value(middleware.UserKey).(*model.User)

	c.JSON(http.StatusOK, user)
}

// ProfileEdit modify the information of the user
func ProfileEdit(c *gin.Context) {
	user := c.Value(middleware.UserKey).(*model.User)

	var in = &inputUser{}

	if err := json.NewDecoder(c.Request.Body).Decode(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	in.Email = strings.ToLower(in.Email)

	if in.Email != user.Email && len(in.Email) > 0 {
		user.Email = in.Email
	}
	if in.FirstName != user.FirstName && len(in.FirstName) > 0 {
		user.FirstName = in.FirstName
	}
	if in.LastName != user.LastName && len(in.LastName) > 0 {
		user.LastName = in.LastName
	}
	if HashPassword(in.Password) != user.Password && len(in.Password) > 0 {
		user.Password = HashPassword(in.Password)
	}

	err := store.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": ErrAccountAlreadyExist.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/takeanote/takeanote-api/store"
)

const (
	// UserKey is the key to reference the user in context.
	UserKey = "user"
)

var (
	// ErrInvalidCredentials is returned if an email and password don't match a db entry.
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// Session handles token checking.
func Session(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:7]) == "BEARER " {
		userToken := bearer[7:]
		email, err := store.GetEmailByToken(c, userToken)
		if err == nil && len(email) > 0 {
			user, err := store.GetUserByEmail(c, email)
			if err != nil {
				c.AbortWithError(http.StatusUnauthorized, ErrInvalidCredentials)
				c.Next()
				return
			}
			c.Set(UserKey, user)
		}
	} else {
		c.AbortWithError(http.StatusUnauthorized, ErrInvalidCredentials)
	}
	c.Next()
}

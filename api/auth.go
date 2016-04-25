package api

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/takeanote/takeanote-api/model"
	"github.com/takeanote/takeanote-api/store"

	"github.com/satori/go.uuid"
)

type newUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

const (
	// TokenDuration is validity duration of a token.
	TokenDuration time.Duration = 24 * time.Hour
)

// Token holds the user token
type Token struct {
	Token string `json:"token"`
}

// HashPassword hash password using sha512 and salt
func HashPassword(password string) string {
	shaHandler := sha512.New()
	shaHandler.Write([]byte(password))

	return hex.EncodeToString(shaHandler.Sum(nil))
}

// SignUp handles User creation
func SignUp(c *gin.Context) {
	in := &newUser{}

	if err := json.NewDecoder(c.Request.Body).Decode(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	in.Email = strings.ToLower(in.Email)
	user := &model.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  HashPassword(in.Password),
	}

	err := store.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
	return
}

// SignIn handles User creation
func SignIn(c *gin.Context) {
	in := &newUser{}

	if err := json.NewDecoder(c.Request.Body).Decode(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	in.Email = strings.ToLower(in.Email)

	user, err := store.GetUserByEmailPassword(c, in.Email, HashPassword(in.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": ErrInvalidCredentials.Error(),
		})
		return
	}

	uuid := uuid.NewV4()
	store.CreateToken(c, uuid.String(), user.Email, TokenDuration)

	c.JSON(http.StatusOK, Token{
		Token: uuid.String(),
	})
}

// SignOut handles User creation
func SignOut(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:7]) == "BEARER " {
		userToken := bearer[7:]
		email, err := store.GetEmailByToken(c, userToken)
		if err == nil && len(email) > 0 {
			err = store.DeleteToken(c, userToken)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": ErrInvalidCredentials.Error(),
	})
	return
}

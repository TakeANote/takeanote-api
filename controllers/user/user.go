package user

import (
	"errors"
	"net/http"

	"github.com/gorilla/context"
	"github.com/takeanote/takeanote-api/controllers/auth"
	"github.com/takeanote/takeanote-api/httputils"
	"github.com/takeanote/takeanote-api/models"

	"github.com/jinzhu/gorm"
)

var (
	// ErrTokenInvalid is returned if token is invalid or expired.
	ErrTokenInvalid = errors.New("token invalid or expired")
	// ErrAccountAlreadyExist is returned if an email already exists in the db.
	ErrAccountAlreadyExist = errors.New("the account already exist")
	// ErrInvalidCredentials is returned if an email and password don't match a db entry.
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrWrongPassword is returned if the provided password doesn't match.
	ErrWrongPassword = errors.New("wrong password")
)

type user struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Controller handles User routines.
type Controller struct {
	DB *gorm.DB
}

// NewController instantiate a new Controller that handles user routines.
func NewController(db *gorm.DB) *Controller {
	return &Controller{
		DB: db,
	}
}

// ProfileView return the information of the user
func (controller Controller) ProfileView(w http.ResponseWriter, r *http.Request, vars map[string]string) error {

	dbUser := context.Get(r, auth.UserKey).(models.User)
	httputils.WriteJSON(w, http.StatusOK, dbUser)

	return nil
}

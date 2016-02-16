package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/takeanote/takeanote-api/httputils"
	"github.com/takeanote/takeanote-api/models"
	"github.com/takeanote/takeanote-api/config"

	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v3"
)

var (
	// ErrTokenInvalid is returned if token is invalid or expired.
	ErrTokenInvalid = errors.New("token invalid or expired")
	// ErrAccountAlreadyExist is returned if an email already exists in the db.
	ErrAccountAlreadyExist = errors.New("the account already exist")
	// ErrWrongPassword is returned if the provided password doesn't match.
	ErrWrongPassword = errors.New("wrong password")
)

type user struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Controller handles authentication routines.
type Controller struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

// NewController instantiate a new Controller that handles auth routines.
func NewController(db *gorm.DB, cfg *config.Config) *Controller {
	rdis := redis.NewClient(&redis.Options{
		Addr: cfg.Redis,
	})
	return &Controller{
		DB:          db,
		RedisClient: rdis,
	}
}

// HashPassword hash password using sha512 and salt
func hashPassword(password string) string {
	shaHandler := sha512.New()
	shaHandler.Write([]byte(password))

	return hex.EncodeToString(shaHandler.Sum(nil))
}

// SignUp handles User creation
func (controller Controller) SignUp(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	user := &user{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		httputils.WriteError(w, models.NewError(http.StatusBadRequest, err))
		return err
	}
	user.Email = strings.ToLower(user.Email)
	dbUser := &models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashPassword(user.Password),
	}

	err := controller.DB.Create(dbUser).Error
	if err != nil {
		httputils.WriteError(w, models.NewError(http.StatusConflict, ErrAccountAlreadyExist))
		return err
	}

	httputils.WriteJSON(w, http.StatusNoContent, nil)

	return nil
}

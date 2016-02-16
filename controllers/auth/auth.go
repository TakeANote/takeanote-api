package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/httputils"
	"github.com/takeanote/takeanote-api/models"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v3"
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

const (
	// TokenDuration is validity duration of a token.
	TokenDuration = 24 * time.Hour
)

// Token holds the user token
type Token struct {
	Token string `json:"token"`
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

// SignIn handles User creation
func (controller Controller) SignIn(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	user := &user{}
	dbUser := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		httputils.WriteError(w, models.NewError(http.StatusBadRequest, err))
		return err
	}
	user.Email = strings.ToLower(user.Email)

	err := controller.DB.Where("email = ? AND password = ?",
		user.Email, hashPassword(user.Password)).Find(&dbUser).Error
	if err != nil {
		httputils.WriteError(w, models.NewError(http.StatusUnauthorized, ErrInvalidCredentials))
		return err
	}

	uuid := uuid.NewV4()
	controller.RedisClient.Set(uuid.String(), user.Email, TokenDuration)

	httputils.WriteJSON(w, http.StatusOK, Token{
		Token: uuid.String(),
	})

	return nil
}

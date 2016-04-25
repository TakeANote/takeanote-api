package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takeanote/takeanote-api/api"
	"github.com/takeanote/takeanote-api/router/middleware"
	"github.com/takeanote/takeanote-api/router/middleware/header"
)

// Load takes infinite number of middleware and apply them in order
// to the gin router.
func Load(middlewares ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(header.NoCache)
	e.Use(header.Options)
	e.Use(header.Secure)
	e.Use(middlewares...)

	v1 := e.Group("/v1")
	{
		v1.POST("/users", api.SignUp)

		v1.POST("/session", api.SignIn)
		v1.DELETE("/session", api.SignOut)

		auth := v1.Group("")
		{
			auth.Use(middleware.Session)
			auth.GET("/profile", api.ProfileView)
			auth.PATCH("/profile", api.ProfileEdit)
		}
	}
	return e
}

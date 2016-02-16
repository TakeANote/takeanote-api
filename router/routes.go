package router

import (
	"fmt"
	"net/http"

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/controllers/auth"
	"github.com/takeanote/takeanote-api/httputils"
	"github.com/takeanote/takeanote-api/models"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// route defines an individual API route to connect with the docker daemon.
// It implements router.Route.
type route struct {
	method  string
	path    string
	handler httputils.APIFunc
}

type router struct {
	routes []Route
}

func (r route) Register(m *mux.Router, handler http.Handler) {
	logrus.Debugf("Registering %s, %s", r.method, r.path)
	m.Path(r.path).Methods(r.method).Handler(handler)
}

// Handler returns the APIFunc to let the API wrap it in middlewares
func (r route) Handler() httputils.APIFunc {
	return r.handler
}

// NewRoute initialies a new local route for the router
func NewRoute(method, path string, handler httputils.APIFunc) Route {
	return route{method, path, handler}
}

// NewGetRoute initializes a new route with the http method GET.
func NewGetRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("GET", path, handler)
}

// NewPostRoute initializes a new route with the http method POST.
func NewPostRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("POST", path, handler)
}

// NewPutRoute initializes a new route with the http method PUT.
func NewPutRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("PUT", path, handler)
}

// NewPatchRoute initializes a new route with the http method PATCH.
func NewPatchRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("PATCH", path, handler)
}

// NewDeleteRoute initializes a new route with the http method DELETE
func NewDeleteRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("DELETE", path, handler)
}

// NewOptionsRoute initializes a new route with the http method OPTIONS
func NewOptionsRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("OPTIONS", path, handler)
}

// NewHeadRoute initializes a new route with the http method HEAD.
func NewHeadRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("HEAD", path, handler)
}

// New initializes a local router with a new daemon.
func New(config *config.Config) Router {
	r := &router{}
	if db, err := models.OpenDBWithConfig(config); err != nil {
		panic(fmt.Errorf("fatal error cannot connect to database: %s\n", err))
	} else {
		r.initRoutes(&db, config)
	}
	return r
}

// Routes returns the list of routes registered in the router.
func (r *router) Routes() []Route {
	return r.routes
}

// initRoutes initializes the routes in this router
// Use authController.OAuth2Middleware(httputils.APIFunc) to protect a route.
func (r *router) initRoutes(db *gorm.DB, cfg *config.Config) {
	db.AutoMigrate(&models.User{})
	auth := auth.NewController(db, cfg)
	r.routes = []Route{
		// GET
		// POST
		NewPostRoute("/signup", auth.SignUp),
		NewPostRoute("/signin", auth.SignIn),
		// PUT
		// PATCH
		// DELETE
		// HEAD
		// OPTIONS
	}
}

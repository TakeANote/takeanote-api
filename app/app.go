package app

import (
	"net/http"

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/httputils"
	"github.com/takeanote/takeanote-api/router"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// App contains what's needed to run the API.
type App struct {
	config *config.Config
	router router.Router
	logger *logrus.Logger
}

// NewApp create a new App.
func NewApp(config *config.Config,
	router router.Router,
	log *logrus.Logger) (*App, error) {
	var err error

	log.Level, err = logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	return &App{
		config: config,
		router: router,
		logger: log,
	}, nil
}

// GetLogger returns the App's Logger.
func (app App) GetLogger() *logrus.Logger {
	return app.logger
}

// GetPort returns the App's port.
func (app App) GetPort() int {
	return app.config.Port
}

// RegisterRoutes registers all controllers routes to a mux.Router.
func (app App) RegisterRoutes(m *mux.Router) {
	for _, r := range app.router.Routes() {
		f := makeHTTPHandler(r.Handler())
		r.Register(m, f)
	}
}

// makeHTTPHandler returns a httputils.APIFunc wrapped in a http.HandlerFunc.
func makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r, mux.Vars(r)); err != nil {
			logrus.Errorf("Handler for %s %s returned error: %s", r.Method, r.URL.Path, err.Error())
		}
	}
}

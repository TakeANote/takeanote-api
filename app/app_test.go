package app

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/router"
)

type TestRouter struct {
}

var m = mux.NewRouter()

var TestRoutes = []router.Route{
	router.NewGetRoute("/v1/test/5",
		func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return nil
		}),
	router.NewPostRoute("/v1/test/5",
		func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return nil
		}),
	router.NewPutRoute("/v1/test/5",
		func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return nil
		}),
	router.NewPatchRoute("/v1/test/5",
		func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return nil
		}),
	router.NewDeleteRoute("/v1/test/5",
		func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return nil
		}),
	router.NewHeadRoute("/v1/test/5",
		func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return nil
		}),
}

func (tr TestRouter) Routes() []router.Route {
	return TestRoutes
}

func TestNewApp(t *testing.T) {
	conf := &config.Config{
		LogLevel: "debug",
	}
	router := TestRouter{}
	logger := logrus.New()
	app, _ := NewApp(conf, router, logger)
	assert.Equal(t, conf, app.config)
	assert.Equal(t, router, app.router)
	assert.Equal(t, logger, app.logger)
}

func TestNewAppWithErrors(t *testing.T) {
	conf := &config.Config{
		LogLevel: "fail",
	}
	router := TestRouter{}
	logger := logrus.New()
	_, err := NewApp(conf, router, logger)
	if err == nil {
		assert.Fail(t, "Should return an error")
	}
}

func TestGetLogger(t *testing.T) {
	logger := logrus.New()
	app := App{
		logger: logger,
	}
	assert.Equal(t, logger, app.GetLogger())
}

func TestGetPort(t *testing.T) {
	logger := logrus.New()
	config := config.Config{
		Port: 9595,
	}
	app := App{
		config: &config,
		logger: logger,
	}
	assert.Equal(t, 9595, app.GetPort())
}

func TestRegisterRoutes(t *testing.T) {
	logger := logrus.New()
	config := config.Config{
		Port: 9595,
	}
	app := App{
		config: &config,
		logger: logger,
		router: TestRouter{},
	}
	app.RegisterRoutes(m)
	routeNb := 0
	m.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeNb++
		return nil
	})
	assert.Equal(t, 6, routeNb)
}

func TestMakeHTTPHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://api.takeanote.org/v1/test/5", nil)

	w := httptest.NewRecorder()
	testHandler := makeHTTPHandler(func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		assert.Equal(t, "5", mux.Vars(r)["id"])
		return nil
	})
	m.HandleFunc("/v1/test/{id:[0-9]+}", testHandler)

	m.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestMakeHTTPHandlerWithErrors(t *testing.T) {
	var buffer bytes.Buffer
	logrus.SetOutput(&buffer)
	req, _ := http.NewRequest("GET", "http://api.takeanote.org", nil)
	w := httptest.NewRecorder()
	testHandler := makeHTTPHandler(func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		return fmt.Errorf("Something happened")
	})
	testHandler(w, req)
	assert.NotEqual(t, 0, len(buffer.String()))
}

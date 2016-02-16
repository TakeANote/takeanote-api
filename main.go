package main

import (
	"fmt"

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/router"
	"github.com/takeanote/takeanote-api/app"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/phyber/negroni-gzip/gzip"
)

func main() {
	conf := config.New()
	app, err := app.NewApp(conf, router.New(conf), logrus.New())
	if err != nil {
		panic(err)
	}

	m := mux.NewRouter()
	m = m.StrictSlash(false).PathPrefix("/v1").Subrouter()

	app.RegisterRoutes(m)
	n := negroni.New()

	// Middlewares
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewMiddlewareFromLogger(app.GetLogger(), "takeanote"))

	n.UseHandler(m)
	n.Run(fmt.Sprintf(":%d", app.GetPort()))
}

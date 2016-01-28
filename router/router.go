package router

import (
	"net/http"

	"github.com/takeanote/takeanote-api/httputils"

	"github.com/gorilla/mux"
)

// Router defines an interface to specify a group of routes to add to API.
type Router interface {
	Routes() []Route
}

// Route defines an individual API route in the API.
type Route interface {
	// Register adds the handler route to the docker mux.
	Register(*mux.Router, http.Handler)
	// Handler returns the raw function to create the http handler.
	Handler() httputils.APIFunc
}

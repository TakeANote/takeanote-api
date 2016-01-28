package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/takeanote/takeanote-api/models"
)

// APIFunc is an adapter to allow the use of ordinary functions as TakeANote API endpoints.
// Any function that has the appropriate signature can be register as a API endpoint
type APIFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string) error

// WriteError decodes an error and sends it in the response.
func WriteError(w http.ResponseWriter, err models.Error) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err.HTTPStatusCode)
	return json.NewEncoder(w).Encode(err)
}

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

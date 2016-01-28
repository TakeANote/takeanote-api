package httputils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/takeanote/takeanote-api/models"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	integer int
	strng   string
}

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
	var err models.Error

	err = models.Error{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        "This is an error",
	}

	WriteError(w, err)
	jsonObj, _ := json.Marshal(err)
	jsonObj = append(jsonObj, '\n')

	assert.Equal(t, jsonObj, w.Body.Bytes())
}

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()

	obj := testObject{integer: 894, strng: "Hey you!"}
	jsonObj, _ := json.Marshal(obj)
	WriteJSON(w, 200, obj)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, []string{"application/json; charset=utf-8"}, w.HeaderMap["Content-Type"])
	jsonObj = append(jsonObj, '\n')
	assert.Equal(t, jsonObj, w.Body.Bytes())
}

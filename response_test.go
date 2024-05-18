package restkit

import (
	"encoding/json"
	"errors"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertJsonResponse(t *testing.T, w *httptest.ResponseRecorder, status int, data any) {
	ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	assert.NoError(t, err)
	assert.Equal(t, "application/json", ct)
	assert.Equal(t, status, w.Code)
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&data))
}

func Test_RespondJson(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}

	// Act
	w := httptest.NewRecorder()
	WriteResponseJson(w, http.StatusCreated, &data)

	// Assert
	var result map[string]string
	assertJsonResponse(t, w, http.StatusCreated, &result)
	assert.Equal(t, data, result)
}

func Test_ExecuteJsonEndpoint_Success(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}

	// Act
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	err := ExecuteJsonEndpoint(w, r, http.StatusCreated, func() (any, error) {
		return data, nil
	})

	// Assert
	assert.NoError(t, err)

	var result map[string]string
	assertJsonResponse(t, w, http.StatusCreated, &result)
	assert.Equal(t, data, result)
}

func Test_ExecuteJsonEndpoint_Failure(t *testing.T) {
	// Act
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	expected := errors.New("the world exploded")
	err := ExecuteJsonEndpoint(w, r, http.StatusCreated, func() (any, error) {
		return nil, expected
	})

	// Assert
	assert.Equal(t, expected, err)
	assert.Equal(t, http.StatusOK, w.Code) // The error is not written by this
}

func Test_ReadJsonResponse(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(data))

	// Act
	var result map[string]string
	err := ReadJsonResponse(w.Result(), &result)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, data, result)
}

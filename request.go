package restkit

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var ErrBadJson = errors.New("invalid json")

// Reads the content of a JSON request and binds it to the provided data
func ReadRequestBody(r *http.Request, data any) error {
	if err := validateContentType(r.Header.Get("Content-Type"), "application/json"); err != nil {
		return err
	} else if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return ErrBadJson
	}

	return nil
}

// Parses and validates the request's JSON content then binds it to the given type. If any step
// fails, the error is returned. This error will render a valid HTTP response with any restkit error filter
func WithJsonBody[T any](r *http.Request, f func(data T) error) error {
	var data T
	if err := ReadRequestBody(r, &data); err != nil {
		return err
	}

	return f(data)
}

// Creates a new HTTP request with JSON content
func NewJsonRequest(method, url string, data any) (*http.Request, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, ErrBadJson
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

package restkit

import (
	"encoding/json"
	"net/http"
)

// Writes a JSON response with the provided status code
func WriteResponseJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// Executes a basic REST JSON endpoint. Assumes that function provided returns an object to render as JSON
// or an error. If an error is returned, it will be returned from this method
func ExecuteJsonEndpoint(w http.ResponseWriter, r *http.Request, successStatus int, f func() (any, error)) error {
	data, err := f()
	if err != nil {
		return err
	}

	return WriteResponseJson(w, successStatus, data)
}

// Validates the given response is a JSON response and binds the data to the provided target
func ReadJsonResponse(r *http.Response, data any) error {
	if err := validateContentType(r.Header.Get("Content-Type"), "application/json"); err != nil {
		return err
	} else if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return ErrBadJson
	}

	return nil
}

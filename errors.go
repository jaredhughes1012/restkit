package restkit

import (
	"log/slog"
	"net/http"
)

type ErrorMap map[error]int

// ErrorFilter is a filter that writes an HTTP error response for a given error
type ErrorFilter struct {
	errors ErrorMap
}

// Creates a new error filter. This filter will include known errors provided by restkit
func NewErrorFilter(errMap ErrorMap) *ErrorFilter {
	errs := ErrorMap{
		ErrBadContentType:         http.StatusBadRequest,
		ErrUnsupportedContentType: http.StatusUnsupportedMediaType,
		ErrBadJson:                http.StatusBadRequest,
	}

	for k, v := range errMap {
		errs[k] = v
	}

	return &ErrorFilter{errors: errs}
}

// Writes an HTTP error response for a given error, if the error is not nil. Returns false if the error is nil.
func (ef *ErrorFilter) WriteIfError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	status, ok := ef.errors[err]
	if !ok {
		status = http.StatusInternalServerError
	}

	if status >= 500 {
		slog.Error("Error in request", "message", err.Error())
	}

	http.Error(w, http.StatusText(status), status)
	return true
}

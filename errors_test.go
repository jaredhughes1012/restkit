package restkit

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrorFilter_WriteIfError(t *testing.T) {
	customErr := errors.New("custom error")
	cases := []struct {
		name   string
		err    error
		status int
	}{
		{
			name:   "nil error",
			err:    nil,
			status: http.StatusOK,
		},
		{
			name:   "bad json",
			err:    ErrBadJson,
			status: http.StatusBadRequest,
		},
		{
			name:   "unsupported content type",
			err:    ErrUnsupportedContentType,
			status: http.StatusUnsupportedMediaType,
		},
		{
			name:   "custom error",
			err:    customErr,
			status: http.StatusNotFound,
		},
		{
			name:   "unknown error",
			err:    errors.New("unknown error"),
			status: http.StatusInternalServerError,
		},
	}

	errFilter := NewErrorFilter(ErrorMap{
		customErr: http.StatusNotFound,
	})

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ok := errFilter.WriteIfError(w, tc.err)
			assert.Equal(t, ok, tc.err != nil)
			assert.Equal(t, w.Code, tc.status)
		})
	}
}

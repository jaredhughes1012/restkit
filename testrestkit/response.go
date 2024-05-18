package testrestkit

import (
	"net/http/httptest"
	"testing"

	"github.com/jaredhughes1012/restkit"
	"github.com/stretchr/testify/require"
)

func RequireJsonResponse[T any](t *testing.T, w *httptest.ResponseRecorder, status int) T {
	t.Helper()

	var data T
	require.Equal(t, status, w.Code)
	require.NoError(t, restkit.ReadJsonResponse(w.Result(), &data))

	return data
}

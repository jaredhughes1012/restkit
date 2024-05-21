package testrestkit

import (
	"net/http"
	"testing"

	"github.com/jaredhughes1012/restkit"
	"github.com/stretchr/testify/require"
)

func RequireJsonResponse[T any](t *testing.T, res *http.Response, status int) T {
	t.Helper()

	var data T
	require.Equal(t, status, res.StatusCode)
	require.NoError(t, restkit.ReadJsonResponse(res, &data))

	return data
}

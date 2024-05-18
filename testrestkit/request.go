package testrestkit

import (
	"net/http"
	"testing"

	"github.com/jaredhughes1012/restkit"
	"github.com/stretchr/testify/require"
)

func NewJsonRequest(t *testing.T, method, url string, data any) *http.Request {
	t.Helper()
	req, err := restkit.NewJsonRequest(method, url, data)
	require.NoError(t, err)
	return req
}

package testrestkit

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/jaredhughes1012/restkit"
	"github.com/stretchr/testify/require"
)

// Creates a full URL from a base and a path. Ensures no errors occur while resolving.
// Useful for creating full URLs from httptest.Server URL
func ApiUrl(t *testing.T, baseUrl, path string) string {
	t.Helper()

	ubase, err := url.Parse(baseUrl)
	require.NoError(t, err)

	upath, err := url.Parse(path)
	require.NoError(t, err)

	return ubase.ResolveReference(upath).String()
}

// Executes a request with a json body. Should be used with the httptest.Server client
func DoJsonRequest(t *testing.T, client *http.Client, method, url string, data any) *http.Response {
	t.Helper()

	req := NewJsonRequest(t, method, url, data)
	res, err := client.Do(req)
	require.NoError(t, err)
	return res
}

func NewJsonRequest(t *testing.T, method, url string, data any) *http.Request {
	t.Helper()
	req, err := restkit.NewJsonRequest(method, url, data)
	require.NoError(t, err)
	return req
}

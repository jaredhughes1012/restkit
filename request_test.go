package restkit

import (
	"bytes"
	"encoding/json"
	"mime"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadRequestBody(t *testing.T) {
	b, _ := json.Marshal(map[string]string{"key": "value"})
	r := httptest.NewRequest("POST", "/test", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")

	var data map[string]string
	err := ReadRequestBody(r, &data)

	assert.NoError(t, err)
	assert.Equal(t, data["key"], "value")
}

func Test_ReadRequestBody_NonJson(t *testing.T) {
	b, _ := json.Marshal(map[string]string{"key": "value"})
	r := httptest.NewRequest("POST", "/test", bytes.NewReader(b))
	r.Header.Set("Content-Type", "text/plain")

	var data map[string]string
	err := ReadRequestBody(r, &data)

	assert.Equal(t, ErrUnsupportedContentType, err)
}

func Test_ReadRequestBody_BadJson(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", bytes.NewReader([]byte("{{")))
	r.Header.Set("Content-Type", "application/json")

	var data map[string]string
	err := ReadRequestBody(r, &data)

	assert.Equal(t, ErrBadJson, err)
}

func Test_WithJsonBody(t *testing.T) {
	b, _ := json.Marshal(map[string]string{"key": "value"})
	r := httptest.NewRequest("POST", "/test", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")

	err := WithJsonBody(r, func(result map[string]any) error {
		assert.Equal(t, result["key"], "value")
		return nil
	})

	assert.NoError(t, err)
}

func Test_NewJsonRequest(t *testing.T) {
	r, err := NewJsonRequest("POST", "/test", map[string]string{"key": "value"})
	assert.NoError(t, err)

	var data map[string]string
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	assert.NoError(t, err)
	assert.Equal(t, ct, "application/json")
	assert.NoError(t, json.NewDecoder(r.Body).Decode(&data))
	assert.Equal(t, data["key"], "value")
}

package restkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateContentType_Valid(t *testing.T) {
	err := validateContentType("application/json", "application/json")
	assert.NoError(t, err)
}

func Test_validateContentType_Invalid(t *testing.T) {
	err := validateContentType("application/xml", "application/json")
	assert.Error(t, err)
}

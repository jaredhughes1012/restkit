package restkit

import (
	"errors"
	"mime"
)

var (
	ErrUnsupportedContentType = errors.New("invalid content type")
	ErrBadContentType         = errors.New("bad content type")
)

func validateContentType(header string, contentType string) error {
	ct, _, err := mime.ParseMediaType(header)
	if err != nil {
		return ErrBadContentType
	}

	if ct != contentType {
		return ErrUnsupportedContentType
	}

	return nil
}

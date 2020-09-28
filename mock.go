package recaptcha

import (
	"net/http"
)

// Mock the recaptcha service during testing.
type Mock struct {
}

// NewMock returns a mock reCaptcha service instance.
func NewMock() (*Mock, error) {
	return &Mock{}, nil
}

// Verify the mock reCaptcha token.
func (Mock) Verify(response string) error {
	if response != "mock-recaptcha" {
		return errRequestFailure
	}

	return nil
}

// GetRecaptchaToken returns 'mock-recaptcha'.
func (Mock) GetRecaptchaToken(_ *http.Request) (string, error) {
	return "mock-recaptcha", nil
}

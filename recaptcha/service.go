package recaptcha

import (
	"errors"
	"net/http"
	"time"
)

const (
	siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

var (
	errMissingSecret = errors.New("missing recaptcha secret")

	errNilClient = errors.New("invalid http client, can't be nil")

	errInvalidVersion = errors.New("invalid version, must be 2 or 3")

	errInvalidScore = errors.New("invalid min score, must be between 0.0 and 1.0")

	errInvalidAction = errors.New("invalid action, can't be empty")

	errRequestFailure = errors.New("unsuccessful recaptcha verify request")

	errLowerScore = errors.New("lower received score than expected")

	errMismatchAction = errors.New("mismatched recaptcha action")
)

type siteVerifyRequest struct {
	RecaptchaResponse string `json:"g-recaptcha-response"`
}

type siteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// Service interface for reCaptcha package.
type Service interface {
	Verify(string) error
	GetRecaptchaToken(*http.Request) (string, error)
}

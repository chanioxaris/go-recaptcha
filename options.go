package recaptcha

import (
	"net/http"
)

// Option describes a functional parameter for the New constructor.
type Option func(*Recaptcha)

// WithHTTPClient allows for overriding of the http client.
func WithHTTPClient(client *http.Client) Option {
	return func(rec *Recaptcha) {
		rec.Client = client
	}
}

// WithVersion allows for overriding of the reCaptcha version.
// Default value is 3.
func WithVersion(version int) Option {
	return func(rec *Recaptcha) {
		rec.Version = version
	}
}

// WithAction allows for overriding of the reCaptcha action.
// Default value is empty string.
// Only applicable for reCaptcha V3.
func WithAction(action string) Option {
	return func(rec *Recaptcha) {
		rec.Action = action
	}
}

// WithMinScore allows for overriding of the minimum reCaptcha accepted score.
// Default value is 0.5.
// Only applicable for reCaptcha V3.
func WithMinScore(score float64) Option {
	return func(rec *Recaptcha) {
		rec.Score = score
	}
}

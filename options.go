package recaptcha

import (
	"net/http"
)

// Option describes a functional parameter for the New constructor.
type Option func(*Recaptcha) error

// WithHTTPClient allows for overriding of the http client.
func WithHTTPClient(client *http.Client) Option {
	return func(rec *Recaptcha) error {
		if client == nil {
			return errNilClient
		}

		rec.client = client
		return nil
	}
}

// WithVersion allows for overriding of the reCaptcha version.
// Default value is 3.
func WithVersion(version int) Option {
	return func(rec *Recaptcha) error {
		if version != 2 && version != 3 {
			return errInvalidVersion
		}

		rec.version = version
		return nil
	}
}

// WithAction allows for overriding of the reCaptcha action.
// Default value is empty string.
// Only applicable for reCaptcha V3.
func WithAction(action string) Option {
	return func(rec *Recaptcha) error {
		if action == "" {
			return errInvalidAction
		}

		rec.action = action
		return nil
	}
}

// WithScore allows for overriding of the minimum reCaptcha accepted score.
// Default value is 0.5.
// Only applicable for reCaptcha V3.
func WithScore(score float64) Option {
	return func(rec *Recaptcha) error {
		if score < 0.0 || score > 1.0 {
			return errInvalidScore
		}

		rec.score = score
		return nil
	}
}

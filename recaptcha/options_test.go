package recaptcha

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	testDefaultRecaptcha = Recaptcha{
		secret:  "secret",
		client:  http.DefaultClient,
		version: 3,
		action:  "",
		score:   0.5,
	}
)

func TestWithHTTPClient(t *testing.T) {
	recaptchaWithHTTPClient := testDefaultRecaptcha
	recaptchaWithHTTPClient.client = &http.Client{
		Timeout: time.Second * 10,
	}

	tests := []struct {
		name    string
		client  *http.Client
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name:    "Valid http client",
			client:  recaptchaWithHTTPClient.client,
			want:    &recaptchaWithHTTPClient,
			wantErr: false,
		},
		{
			name:    "Invalid http client (nil)",
			client:  nil,
			wantErr: true,
			err:     errNilClient,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New("secret", WithHTTPClient(tt.client))
			if err != nil && !tt.wantErr {
				t.Error(err)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithHTTPClient() = %+v, want %+v", got, tt.want)
				}
			} else {
				if err == nil {
					t.Errorf("WithHTTPClient() expected error got nil")
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithHTTPClient() error = %v, want %v", err, tt.err)
				}
			}
		})
	}
}

func TestWithVersion(t *testing.T) {
	recaptchaWithVersion := testDefaultRecaptcha
	recaptchaWithVersion.version = 2

	tests := []struct {
		name    string
		version int
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name:    "Valid version",
			version: recaptchaWithVersion.version,
			want:    &recaptchaWithVersion,
			wantErr: false,
		},
		{
			name:    "Invalid version (13)",
			version: 13,
			wantErr: true,
			err:     errInvalidVersion,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New("secret", WithVersion(tt.version))
			if err != nil && !tt.wantErr {
				t.Error(err)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithVersion() = %+v, want %+v", got, tt.want)
				}
			} else {
				if err == nil {
					t.Errorf("WithVersion() expected error got nil")
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithVersion() error = %v, want %v", err, tt.err)
				}
			}
		})
	}
}

func TestWithAction(t *testing.T) {
	recaptchaWithAction := testDefaultRecaptcha
	recaptchaWithAction.action = "test-action"

	tests := []struct {
		name    string
		action  string
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name:    "Valid action",
			action:  recaptchaWithAction.action,
			want:    &recaptchaWithAction,
			wantErr: false,
		},
		{
			name:    "Invalid action (empty string)",
			action:  "",
			wantErr: true,
			err:     errInvalidAction,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New("secret", WithAction(tt.action))
			if err != nil && !tt.wantErr {
				t.Error(err)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithAction() = %+v, want %+v", got, tt.want)
				}
			} else {
				if err == nil {
					t.Errorf("WithAction() expected error got nil")
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithAction() error = %v, want %v", err, tt.err)
				}
			}
		})
	}
}

func TestWithMinScore(t *testing.T) {
	recaptchaWithMinScore := testDefaultRecaptcha
	recaptchaWithMinScore.score = 0.8

	tests := []struct {
		name    string
		score   float64
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name:    "Valid min score",
			score:   recaptchaWithMinScore.score,
			want:    &recaptchaWithMinScore,
			wantErr: false,
		},
		{
			name:    "Invalid min score (1.3)",
			score:   1.3,
			wantErr: true,
			err:     errInvalidScore,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New("secret", WithMinScore(tt.score))
			if err != nil && !tt.wantErr {
				t.Error(err)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithMinScore() = %+v, want %+v", got, tt.want)
				}
			} else {
				if err == nil {
					t.Errorf("WithMinScore() expected error got nil")
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithMinScore() error = %v, want %v", err, tt.err)
				}
			}
		})
	}
}

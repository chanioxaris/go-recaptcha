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
		secret:  "test-secret",
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

	type args struct {
		secret  string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name: "Valid http client",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithHTTPClient(recaptchaWithHTTPClient.client),
				},
			},
			want:    &recaptchaWithHTTPClient,
			wantErr: false,
		},
		{
			name: "Invalid http client (nil)",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithHTTPClient(nil),
				},
			},
			wantErr: true,
			err:     errNilClient,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.secret, tt.args.options...)
			if err != nil && !tt.wantErr {
				t.Error(err)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithHTTPClient() = %+v, want %+v", got, tt.want)
					return
				}
			} else {
				if err == nil {
					t.Errorf("WithHTTPClient() expected error got nil")
					return
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithHTTPClient() error = %v, want %v", err, tt.err)
					return
				}
			}
		})
	}
}

func TestWithVersion(t *testing.T) {
	recaptchaWithVersion := testDefaultRecaptcha
	recaptchaWithVersion.version = 2

	type args struct {
		secret  string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name: "Valid version",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithVersion(recaptchaWithVersion.version),
				},
			},
			want:    &recaptchaWithVersion,
			wantErr: false,
		},
		{
			name: "Invalid version (13)",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithVersion(13),
				},
			},
			wantErr: true,
			err:     errInvalidVersion,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.secret, tt.args.options...)
			if err != nil && !tt.wantErr {
				t.Error(err)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithVersion() = %+v, want %+v", got, tt.want)
					return
				}
			} else {
				if err == nil {
					t.Errorf("WithVersion() expected error got nil")
					return
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithVersion() error = %v, want %v", err, tt.err)
					return
				}
			}
		})
	}
}

func TestWithAction(t *testing.T) {
	recaptchaWithAction := testDefaultRecaptcha
	recaptchaWithAction.action = "test-action"

	type args struct {
		secret  string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name: "Valid action",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithAction(recaptchaWithAction.action),
				},
			},
			want:    &recaptchaWithAction,
			wantErr: false,
		},
		{
			name: "Invalid action (empty string)",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithAction(""),
				},
			},
			wantErr: true,
			err:     errInvalidAction,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.secret, tt.args.options...)
			if err != nil && !tt.wantErr {
				t.Error(err)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithAction() = %+v, want %+v", got, tt.want)
					return
				}
			} else {
				if err == nil {
					t.Errorf("WithAction() expected error got nil")
					return
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithAction() error = %v, want %v", err, tt.err)
					return
				}
			}
		})
	}
}

func TestWithScore(t *testing.T) {
	recaptchaWithScore := testDefaultRecaptcha
	recaptchaWithScore.score = 0.8

	type args struct {
		secret  string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Recaptcha
		wantErr bool
		err     error
	}{
		{
			name: "Valid score",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithScore(recaptchaWithScore.score),
				},
			},
			want:    &recaptchaWithScore,
			wantErr: false,
		},
		{
			name: "Invalid score (1.3)",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithScore(1.3),
				},
			},
			wantErr: true,
			err:     errInvalidScore,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.secret, tt.args.options...)
			if err != nil && !tt.wantErr {
				t.Error(err)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("WithMinScore() = %+v, want %+v", got, tt.want)
					return
				}
			} else {
				if err == nil {
					t.Errorf("WithMinScore() expected error got nil")
					return
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("WithMinScore() error = %v, want %v", err, tt.err)
					return
				}
			}
		})
	}
}

package recaptcha

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
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
			name: "Valid secret and default values",
			args: args{
				secret: "test-secret",
			},
			want: &Recaptcha{
				secret:  "test-secret",
				client:  http.DefaultClient,
				version: 3,
				action:  "",
				score:   0.5,
			},
			wantErr: false,
		},
		{
			name: "Invalid secret (empty)",
			args: args{
				secret: "",
			},
			wantErr: true,
			err:     errMissingSecret,
		},
		{
			name: "Version 2 with custom http client",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithVersion(2),
					WithHTTPClient(&http.Client{
						Timeout: time.Second * 10,
					}),
				},
			},
			want: &Recaptcha{
				secret: "test-secret",
				client: &http.Client{
					Timeout: time.Second * 10,
				},
				version: 2,
				action:  "",
				score:   0.5,
			},
			wantErr: false,
		},
		{
			name: "Version 3 with custom http client",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithVersion(3),
					WithHTTPClient(&http.Client{
						Timeout: time.Second * 10,
					}),
				},
			},
			want: &Recaptcha{
				secret: "test-secret",
				client: &http.Client{
					Timeout: time.Second * 10,
				},
				version: 3,
				action:  "",
				score:   0.5,
			},
			wantErr: false,
		},
		{
			name: "Version 3 with custom action name",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithVersion(3),
					WithAction("test-action"),
				},
			},
			want: &Recaptcha{
				secret:  "test-secret",
				client:  http.DefaultClient,
				version: 3,
				action:  "test-action",
				score:   0.5,
			},
			wantErr: false,
		},
		{
			name: "Version 3 with custom score",
			args: args{
				secret: "test-secret",
				options: []Option{
					WithVersion(3),
					WithScore(0.8),
				},
			},
			want: &Recaptcha{
				secret:  "test-secret",
				client:  http.DefaultClient,
				version: 3,
				action:  "",
				score:   0.8,
			},
			wantErr: false,
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
					t.Errorf("New() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if err == nil {
					t.Errorf("New() expected error got nil")
					return
				}

				if !errors.Is(err, tt.err) {
					t.Errorf("New() error = %v, want %v", err, tt.err)
					return
				}
			}
		})
	}
}

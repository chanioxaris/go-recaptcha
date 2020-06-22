package middleware

import (
	"net/http"

	"github.com/chanioxaris/go-recaptcha/recaptcha"
)

// Middleware to handle Google reCaptcha verification.
func Middleware(svc recaptcha.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the reCaptcha token from default request body field 'g-recaptcha-response'.
			captchaToken, err := svc.GetRequestToken(r)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			// Verify the retrieved token.
			if err := svc.Verify(captchaToken); err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

package recaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Recaptcha holds all the necessary options to configure and verify the request.
type Recaptcha struct {
	secret  string
	client  *http.Client
	version int
	action  string
	score   float64
}

// New returns a recaptcha service instance.
func New(secret string, options ...Option) (*Recaptcha, error) {
	if secret == "" {
		return nil, errMissingSecret
	}

	newRecaptcha := &Recaptcha{
		secret:  secret,
		client:  http.DefaultClient,
		version: 3,
		action:  "",
		score:   0.5,
	}

	for _, option := range options {
		if err := option(newRecaptcha); err != nil {
			return nil, err
		}
	}

	return newRecaptcha, nil
}

// Verify the provided reCaptcha token depending on version.
func (c *Recaptcha) Verify(response string) error {
	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	// Add necessary request parameters.
	q := req.URL.Query()
	q.Add("secret", c.secret)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var body siteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	if !body.Success {
		return errRequestFailure
	}

	// Check additional response parameters applicable for V3.
	if c.version == 3 {
		if body.Score < c.score {
			return errLowerScore
		}

		if body.Action != c.action {
			return errMismatchAction
		}
	}

	return nil
}

// GetRecaptchaToken from request body 'g-recaptcha-response' field.
func (c *Recaptcha) GetRecaptchaToken(r *http.Request) (string, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var body siteVerifyRequest
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return "", err
	}

	// Restore request body to read more than once.
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return body.RecaptchaResponse, nil
}

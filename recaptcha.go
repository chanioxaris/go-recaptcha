package recaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Recaptcha holds all the necessary options to configure and verify the request.
type Recaptcha struct {
	Secret  string
	Client  *http.Client
	Version int
	Action  string
	Score   float64
}

// New returns a recaptcha service instance.
func New(secret string, options ...Option) (*Recaptcha, error) {
	if secret == "" {
		return nil, errMissingSecret
	}

	newRecaptcha := &Recaptcha{
		Secret:  secret,
		Client:  http.DefaultClient,
		Version: 3,
		Score:   0.5,
	}

	for _, option := range options {
		option(newRecaptcha)
	}

	return newRecaptcha, nil
}

// Verify the provided reCaptcha token depending on version.
func (c *Recaptcha) Verify(response string) error {
	if c.Version != 2 && c.Version != 3 {
		return errUnsupportedVersion
	}

	return c.verify(response)
}

func (c *Recaptcha) verify(response string) error {
	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	// Add necessary request parameters.
	q := req.URL.Query()
	q.Add("secret", c.Secret)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Client.Do(req)
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
	if c.Version == 3 {
		if body.Score < c.Score {
			return errLowerScore
		}

		if body.Action != c.Action {
			return errMismatchAction
		}
	}

	return nil
}

// GetRequestToken from request body 'g-recaptcha-response' field.
func (c *Recaptcha) GetRequestToken(r *http.Request) (string, error) {
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

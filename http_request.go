package chargebee

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cenkalti/backoff"
)

var httpClient = &http.Client{Timeout: DefaultHTTPTimeout}

//Do is used to execute an API Request.
func Do(req *http.Request, env Environment) (string, error) {
	client := httpClient
	if env.HTTPClient != nil {
		client = env.HTTPClient
	}

	backoffConfig := env.BackoffConfig
	if backoffConfig == nil { // if no backoff configured, use default one
		backoffConfig = backoff.NewExponentialBackOff()
	}

	if req.Method != http.MethodGet { // if the request is anything other than GET, don't use backoff
		backoffConfig = &backoff.StopBackOff{}
	}

	var response *http.Response
	var err error
	// Retry on HTTP 429 rate limit, or network error, see:
	// https://godoc.org/github.com/cenkalti/backoff#pkg-constants
	backoff.Retry(func() error {
		response, err = client.Do(req)
		if _, ok := err.(*url.Error); ok {
			return nil
		}
		if response == nil || response.StatusCode == http.StatusTooManyRequests {
			return errors.New("retrying")
		}
		return nil
	}, backoffConfig)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return "", ErrorHandling(resBody)
	}
	return string(resBody), nil
}

func ErrorHandling(resBody []byte) error {
	cbErr := &Error{}
	err := json.Unmarshal(resBody, cbErr)
	if err != nil {
		return err
	}
	if cbErr.APIErrorCode == "" {
		return errors.New("the api_error_code is not present - probably not a chargebee error")

	}
	switch cbErr.Type {

	case PaymentError:
		cbErr.Err = &paymentErr{cbErr: cbErr}
	case InvalidRequestError:
		cbErr.Err = &invalidRequestErr{cbErr: cbErr}
	case OperationFailedError:
		cbErr.Err = &operationFailedErr{cbErr: cbErr}
	}
	return cbErr
}

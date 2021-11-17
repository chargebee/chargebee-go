package chargebee

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// RequestOption allows you to pass along functions that customize a
// RequestOptions struct.
type RequestOption func(*RequestOptions)

// HTTPOption contains a struct of various request options that could be
//customized.
type RequestOptions struct {
	client *http.Client
}

// Client allows you to customize the http client on a per-request basis
func SetClient(hc *http.Client) RequestOption {
	return func(h *RequestOptions) {
		h.client = hc
	}
}

//Do is used to execute an API Request.
func Do(req *http.Request, opts ...RequestOption) (string, error) {
	reqOpts := RequestOptions{
		client: &http.Client{},
	}
	for _, o := range opts {
		o(&reqOpts)
	}
	response, err := reqOpts.client.Do(req)
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

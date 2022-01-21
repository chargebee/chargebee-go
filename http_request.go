package chargebee

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Do is used to execute an API Request.
func Do(req *http.Request) (string, error) {
	if httpClient == nil {
		httpClient = NewDefaultHTTPClient()
	}
	response, err := httpClient.Do(req)
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
		if strings.Contains(string(resBody), "503") {
			return fmt.Errorf("Sorry, the server is currently unable to handle the request due to a temporary overload or scheduled maintenance. Please retry after sometime. \n type: internal_temporary_error, \n http_status_code: 503, \n error_code: internal_temporary_error,\n content:%q", string(resBody))
		} else if strings.Contains(string(resBody), "504") {
			return fmt.Errorf("The server did not receive a timely response from an upstream server, request aborted. If this problem persists, contact us at support@chargebee.com. \n type: gateway_timeout, \n http_status_code: 504, \n error_code: gateway_timeout,\n content:%q", string(resBody))
		} else {
			return fmt.Errorf("Sorry, something went wrong when trying to process the request. If this problem persists, contact us at support@chargebee.com. \n type: internal_error, \n http_status_code: 500, \n error_code: internal_error,\n content:%q", string(resBody))
		}
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

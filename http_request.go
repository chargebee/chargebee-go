package chargebee

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Do(req *http.Request, isIdempotent bool) (*CBResponse, error) {
	client := httpClient
	if client == nil {
		client = &http.Client{Timeout: TotalHTTPTimeout}
	}

	env := DefaultEnv
	if req.Context() != nil {
		if v := req.Context().Value(cbEnvKey); v != nil {
			if e, ok := v.(Environment); ok {
				env = e
			}
		}
	}

	// Retry config defaults
	retryEnabled := false
	maxRetries := 3
	delayMs := 500
	retryOn := map[int]struct{}{500: {}, 502: {}, 503: {}, 504: {}}
	enableDebug := env.EnableDebugLogs

	if cfg := env.RetryConfig; cfg != nil {
		retryEnabled = cfg.Enabled
		if cfg.MaxRetries > 0 {
			maxRetries = cfg.MaxRetries
		}
		if cfg.DelayMs > 0 {
			delayMs = cfg.DelayMs
		}
		if cfg.RetryOn != nil {
			retryOn = cfg.RetryOn
		}
	}

	var bodyCopy []byte
	if req.Body != nil {
		var err error
		bodyCopy, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
	}

	var lastErr error
	var cbResponse *CBResponse

	for attempt := 1; attempt <= maxRetries+1; attempt++ {
		ensureRetryCountHeader(req, attempt)
		ensureIdempotencyKey(req, isIdempotent)
		clonedReq, err := cloneRequest(req, bodyCopy)
		if err != nil {
			return nil, err
		}
		resp, err := client.Do(clonedReq)
		if err != nil {
			lastErr = err
			if enableDebug {
				log.Printf("[Chargebee] HTTP request failed on attempt %d: %v", attempt, err)
			}
			if retryEnabled && attempt <= maxRetries {
				sleepWithBackoff(delayMs, attempt, enableDebug, 0, 0)
				continue
			}
			return nil, err
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			if enableDebug {
				log.Printf("[Chargebee] Failed to read response body on attempt %d: %v", attempt, err)
			}
			if retryEnabled && attempt <= maxRetries {
				sleepWithBackoff(delayMs, attempt, enableDebug, 0, 0)
				continue
			}
			return nil, err
		}

		cbResponse = &CBResponse{
			ResponseMeta: ResponseMeta{
				Headers:    resp.Header,
				Status:     resp.Status,
				StatusCode: resp.StatusCode,
			},
			Body: body,
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			_, shouldRetry := retryOn[resp.StatusCode]
			retryAfter := resp.Header.Get("Retry-After")

			if retryEnabled && shouldRetry && attempt <= maxRetries {
				var retryAfterSecs int
				if retryAfter != "" {
					if secs, err := strconv.Atoi(retryAfter); err == nil {
						retryAfterSecs = secs
					}
				}
				sleepWithBackoff(delayMs, attempt, enableDebug, retryAfterSecs, resp.StatusCode)
				continue
			}

			return cbResponse, ErrorHandling(body)
		}

		return cbResponse, nil
	}

	if cbResponse != nil {
		return cbResponse, lastErr
	}
	return nil, lastErr
}

func cloneRequest(r *http.Request, body []byte) (*http.Request, error) {
	cloned := r.Clone(r.Context())
	if body != nil {
		cloned.Body = io.NopCloser(bytes.NewReader(body))
	}
	return cloned, nil
}

func sleepWithBackoff(baseDelayMs, attempt int, debug bool, retryAfter, statusCode int) {
	var sleep time.Duration

	if retryAfter > 0 {
		sleep = time.Duration(retryAfter) * time.Second
	} else {
		delay := baseDelayMs * (1 << (attempt - 1))
		jitter := rand.Intn(100)
		sleep = time.Duration(delay+jitter) * time.Millisecond
	}

	if debug {
		log.Printf("[Chargebee] Retrying request (status %d) attempt %d after %v", statusCode, attempt, sleep)
	}
	time.Sleep(sleep)
}

func ErrorHandling(resBody []byte) error {
	cbErr := &Error{}
	err := json.Unmarshal(resBody, cbErr)
	if err != nil {
		bodyStr := string(resBody)
		if strings.Contains(bodyStr, "503") {
			return fmt.Errorf("Sorry, the server is currently unable to handle the request due to a temporary overload or scheduled maintenance. Please retry after sometime. \n type: internal_temporary_error, \n http_status_code: 503, \n error_code: internal_temporary_error,\n content: %q", bodyStr)
		} else if strings.Contains(bodyStr, "504") {
			return fmt.Errorf("The server did not receive a timely response from an upstream server, request aborted. If this problem persists, contact us at support@chargebee.com. \n type: gateway_timeout, \n http_status_code: 504, \n error_code: gateway_timeout,\n content:  %q", bodyStr)
		} else {
			return fmt.Errorf("Sorry, something went wrong when trying to process the request. If this problem persists, contact us at support@chargebee.com. \n type: internal_error, \n http_status_code: 500, \n error_code: internal_error,\n content:  %q", bodyStr)
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
	case UbbBatchIngestionInvalidRequestError:
		cbErr.Err = &ubbBatchIngestionInvalidRequestErr{cbErr: cbErr}
	}
	return cbErr
}

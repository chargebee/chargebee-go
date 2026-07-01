package chargebee

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chargebee/chargebee-go/v4/telemetry"
)

type telemetryExecuteInput struct {
	resource       string
	operation      string
	method         string
	httpURL        string
	requestHeaders map[string]string
}

func executeWithTelemetry(
	cfg *ClientConfig,
	input telemetryExecuteInput,
	fn func(headers map[string]string) (httpStatusCode int, err error),
) error {
	if cfg == nil {
		_, err := fn(nil)
		return err
	}

	adapter := cfg.TelemetryAdapter
	if adapter == nil || input.resource == "" || input.operation == "" {
		_, err := fn(nil)
		return err
	}

	start := time.Now()
	headers := map[string]string{}
	handle := startTelemetry(cfg, adapter, input, input.requestHeaders, headers)
	mergedHeaders := headers
	if len(mergedHeaders) == 0 {
		mergedHeaders = nil
	}

	statusCode, err := fn(mergedHeaders)
	if err != nil {
		endTelemetryFailure(adapter, handle, start, err)
		return err
	}

	endTelemetrySuccess(adapter, handle, start, statusCode)
	return nil
}

func startTelemetry(
	cfg *ClientConfig,
	adapter telemetry.TelemetryAdapter,
	input telemetryExecuteInput,
	requestHeaders map[string]string,
	headers map[string]string,
) any {
	parsed, parseErr := url.Parse(input.httpURL)
	if parseErr != nil {
		return nil
	}
	host := parsed.Hostname()
	if host == "" {
		host = parsed.Host
	}
	httpURL := parsed.Scheme + "://" + parsed.Host + parsed.Path
	apiPath := "/api/" + APIVersion

	context := telemetry.BuildRequestTelemetryContext(
		input.resource,
		input.operation,
		strings.ToUpper(input.method),
		httpURL,
		host,
		cfg.SiteName,
		telemetry.ResolveChargebeeAPIVersion(apiPath),
		Version,
		requestHeaders,
	)

	return safeOnRequestStart(adapter, context, headers)
}

func safeOnRequestStart(
	adapter telemetry.TelemetryAdapter,
	context telemetry.RequestTelemetryContext,
	headers map[string]string,
) any {
	defer func() {
		if recovered := recover(); recovered != nil {
			log.Printf("[Chargebee] Telemetry adapter OnRequestStart failed: %v. Continuing without telemetry.", recovered)
		}
	}()
	return adapter.OnRequestStart(context, headers)
}

func endTelemetrySuccess(adapter telemetry.TelemetryAdapter, handle any, start time.Time, httpStatusCode int) {
	result := telemetry.BuildRequestTelemetryResult(httpStatusCode, time.Since(start).Milliseconds(), nil)
	safeOnRequestEnd(adapter, handle, result)
}

func endTelemetryFailure(adapter telemetry.TelemetryAdapter, handle any, start time.Time, err error) {
	statusCode := 500
	if code, ok := extractTelemetryHTTPStatusCode(err); ok {
		statusCode = code
	}
	result := telemetry.BuildRequestTelemetryResult(
		statusCode,
		time.Since(start).Milliseconds(),
		extractRequestTelemetryError(err),
	)
	safeOnRequestEnd(adapter, handle, result)
}

func safeOnRequestEnd(adapter telemetry.TelemetryAdapter, handle any, result telemetry.RequestTelemetryResult) {
	defer func() {
		if recovered := recover(); recovered != nil {
			log.Printf("[Chargebee] Telemetry adapter OnRequestEnd failed: %v", recovered)
		}
	}()
	adapter.OnRequestEnd(handle, result)
}

func extractRequestTelemetryError(err error) *telemetry.RequestTelemetryError {
	if err == nil {
		return nil
	}
	message := err.Error()
	if message == "" {
		message = "Chargebee API request failed"
	}
	var cbErr *Error
	if errors.As(err, &cbErr) {
		return &telemetry.RequestTelemetryError{
			Message:               message,
			ChargebeeErrorCode:    cbErr.APIErrorCode,
			ChargebeeAPIErrorType: string(cbErr.Type),
			ChargebeeErrorParam:   cbErr.Param,
		}
	}
	return &telemetry.RequestTelemetryError{Message: message}
}

func extractTelemetryHTTPStatusCode(err error) (int, bool) {
	var cbErr *Error
	if errors.As(err, &cbErr) {
		return cbErr.HTTPStatusCode, true
	}
	return 0, false
}

func buildTelemetryHTTPURL(cfg *ClientConfig, subDomain, path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	fullURL := cfg.apiBaseUrl(subDomain) + path
	parsed, err := url.Parse(fullURL)
	if err != nil {
		return fullURL
	}
	return parsed.Scheme + "://" + parsed.Host + parsed.Path
}

func httpHeaderToMap(headers *http.Header) map[string]string {
	if headers == nil {
		return nil
	}
	result := make(map[string]string, len(*headers))
	for key, values := range *headers {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}
	return result
}

func mergeExtraHTTPHeaders(base *http.Header, extra map[string]string) *http.Header {
	if len(extra) == 0 {
		return base
	}
	var merged http.Header
	if base != nil {
		merged = base.Clone()
	} else {
		merged = make(http.Header)
	}
	for key, value := range extra {
		merged.Set(key, value)
	}
	return &merged
}

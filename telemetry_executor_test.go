package chargebee

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chargebee/chargebee-go/v4/telemetry"
)

type recordingAdapter struct {
	events       []string
	startContext telemetry.RequestTelemetryContext
	endResult    telemetry.RequestTelemetryResult
}

func (a *recordingAdapter) OnRequestStart(context telemetry.RequestTelemetryContext, requestHeaders map[string]string) any {
	a.events = append(a.events, "start")
	a.startContext = context
	requestHeaders["traceparent"] = "00-test-trace"
	return "span-1"
}

func (a *recordingAdapter) OnRequestEnd(handle any, result telemetry.RequestTelemetryResult) {
	a.events = append(a.events, "end")
	a.endResult = result
}

func testClientConfig(adapter telemetry.TelemetryAdapter) *ClientConfig {
	return &ClientConfig{
		SiteName:         "acme",
		ApiKey:           "test_key",
		TelemetryAdapter: adapter,
	}
}

func TestExecuteWithTelemetrySkipsWithoutAdapter(t *testing.T) {
	cfg := &ClientConfig{SiteName: "acme", ApiKey: "test_key"}

	err := executeWithTelemetry(cfg, telemetryExecuteInput{
		resource:  "customer",
		operation: "list",
		method:    "GET",
		httpURL:   "https://acme.chargebee.com/api/v2/customers",
	}, func(headers map[string]string) (int, error) {
		require.Nil(t, headers)
		return 200, nil
	})
	require.NoError(t, err)
}

func TestExecuteWithTelemetrySkipsWithoutMetadata(t *testing.T) {
	adapter := &recordingAdapter{}
	cfg := testClientConfig(adapter)

	err := executeWithTelemetry(cfg, telemetryExecuteInput{
		method:  "GET",
		httpURL: "https://acme.chargebee.com/api/v2/customers",
	}, func(headers map[string]string) (int, error) {
		return 200, nil
	})
	require.NoError(t, err)
	require.Empty(t, adapter.events)
}

func TestExecuteWithTelemetryCallsAdapterOnce(t *testing.T) {
	adapter := &recordingAdapter{}
	cfg := testClientConfig(adapter)

	err := executeWithTelemetry(cfg, telemetryExecuteInput{
		resource:  "customer",
		operation: "list",
		method:    "GET",
		httpURL:   "https://acme.chargebee.com/api/v2/customers",
	}, func(headers map[string]string) (int, error) {
		require.Equal(t, "00-test-trace", headers["traceparent"])
		return 200, nil
	})
	require.NoError(t, err)
	require.Equal(t, []string{"start", "end"}, adapter.events)
	require.Equal(t, "chargebee.customer.list", adapter.startContext.SpanName)
	require.Equal(t, 200, adapter.endResult.HTTPStatusCode)
}

func TestExecuteWithTelemetryRecordsChargebeeError(t *testing.T) {
	adapter := &recordingAdapter{}
	cfg := testClientConfig(adapter)

	cbErr := &Error{
		HTTPStatusCode: 404,
		Msg:            "Not found",
		APIErrorCode:   "resource_not_found",
		Type:           InvalidRequestError,
		Param:          "customer_id",
	}

	err := executeWithTelemetry(cfg, telemetryExecuteInput{
		resource:  "customer",
		operation: "retrieve",
		method:    "GET",
		httpURL:   "https://acme.chargebee.com/api/v2/customers/1",
	}, func(headers map[string]string) (int, error) {
		return 0, cbErr
	})
	require.Error(t, err)
	require.Equal(t, cbErr, err)
	require.Equal(t, []string{"start", "end"}, adapter.events)
	require.Equal(t, 404, adapter.endResult.HTTPStatusCode)
	require.Equal(t, "resource_not_found", adapter.endResult.Error.ChargebeeErrorCode)
}

func TestExecuteWithTelemetryPromotesChargebeeRequestHeaders(t *testing.T) {
	adapter := &recordingAdapter{}
	cfg := testClientConfig(adapter)

	err := executeWithTelemetry(cfg, telemetryExecuteInput{
		resource:  "customer",
		operation: "list",
		method:    "GET",
		httpURL:   "https://acme.chargebee.com/api/v2/customers",
		requestHeaders: map[string]string{
			"chargebee-foo":               "bar",
			"chargebee-request-origin-ip": "202.170.207.70",
		},
	}, func(headers map[string]string) (int, error) {
		return 200, nil
	})
	require.NoError(t, err)
	require.Equal(t, "bar", adapter.startContext.StartAttributes["http.request.header.chargebee-foo"])
	require.NotContains(t, adapter.startContext.StartAttributes, "http.request.header.chargebee-request-origin-ip")
}

func TestMergeExtraHTTPHeadersDoesNotMutateBase(t *testing.T) {
	base := http.Header{"Chargebee-Foo": []string{"bar"}}
	original := base.Clone()

	merged := mergeExtraHTTPHeaders(&base, map[string]string{"traceparent": "00-test-trace"})

	require.Equal(t, original, base)
	require.Equal(t, "bar", merged.Get("Chargebee-Foo"))
	require.Equal(t, "00-test-trace", merged.Get("traceparent"))
}

func TestExecuteWithTelemetryDoesNotMutateRequestHeaders(t *testing.T) {
	adapter := &recordingAdapter{}
	cfg := testClientConfig(adapter)
	originalHeaders := map[string]string{"chargebee-foo": "bar"}

	err := executeWithTelemetry(cfg, telemetryExecuteInput{
		resource:       "customer",
		operation:      "list",
		method:         "GET",
		httpURL:        "https://acme.chargebee.com/api/v2/customers",
		requestHeaders: originalHeaders,
	}, func(headers map[string]string) (int, error) {
		require.Equal(t, "00-test-trace", headers["traceparent"])
		return 200, nil
	})
	require.NoError(t, err)
	require.Equal(t, map[string]string{"chargebee-foo": "bar"}, originalHeaders)
	require.NotContains(t, originalHeaders, "traceparent")
}

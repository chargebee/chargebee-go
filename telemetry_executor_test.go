package chargebee

import (
	"testing"

	"github.com/chargebee/chargebee-go/v3/telemetry"
	"github.com/stretchr/testify/require"
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

func TestExecuteWithTelemetrySkipsWithoutAdapter(t *testing.T) {
	env := Environment{SiteName: "acme", Key: "test_key"}

	err := executeWithTelemetry(env, telemetryExecuteInput{
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
	env := Environment{
		SiteName:         "acme",
		Key:              "test_key",
		TelemetryAdapter: adapter,
	}

	err := executeWithTelemetry(env, telemetryExecuteInput{
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
	env := Environment{
		SiteName:         "acme",
		Key:              "test_key",
		TelemetryAdapter: adapter,
	}

	err := executeWithTelemetry(env, telemetryExecuteInput{
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
	env := Environment{
		SiteName:         "acme",
		Key:              "test_key",
		TelemetryAdapter: adapter,
	}

	cbErr := &Error{
		HTTPStatusCode: 404,
		Msg:            "Not found",
		APIErrorCode:   "resource_not_found",
		Type:           InvalidRequestError,
		Param:          "customer_id",
	}

	err := executeWithTelemetry(env, telemetryExecuteInput{
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

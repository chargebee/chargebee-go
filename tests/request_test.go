package tests

import (
	"bytes"
	_ "encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/telemetry"
)

type mockTransport struct {
	server *httptest.Server
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	mockReq := req.Clone(req.Context())
	mockReq.URL.Scheme = "http"
	mockReq.URL.Host = strings.TrimPrefix(m.server.URL, "http://")
	mockReq.RequestURI = ""
	return http.DefaultTransport.RoundTrip(mockReq)
}

func TestDo_SuccessFirstTry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, `{"message":"ok"}`)
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := chargebee.Do(req, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !bytes.Contains(resp.Body, []byte("ok")) {
		t.Errorf("expected ok body, got: %s", resp.Body)
	}
}

func TestDo_RetryOn503(t *testing.T) {
	var count int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&count, 1)
		w.WriteHeader(503)
		io.WriteString(w, `{"type":"operation_failed","api_error_code":"temporary_failure"}`)
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	ctx := chargebee.WithEnvironment(req.Context(), chargebee.Environment{
		RetryConfig: &chargebee.RetryConfig{
			Enabled:    true,
			MaxRetries: 2,
			DelayMs:    10,
			RetryOn:    map[int]struct{}{503: {}},
		},
	})
	req = req.WithContext(ctx)

	_, err := chargebee.Do(req, true)
	if err == nil || !strings.Contains(err.Error(), "operation_failed") {
		t.Errorf("expected retryable error, got: %v", err)
	}
	if atomic.LoadInt32(&count) != 3 {
		t.Errorf("expected 3 attempts, got: %d", count)
	}
}

func TestDo_RetryAfterHeader(t *testing.T) {
	var callCount int32
	firstCall := true
	start := time.Now()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&callCount, 1)
		if firstCall {
			firstCall = false
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(503)
			io.WriteString(w, `{"type":"operation_failed","api_error_code":"retry_after_test"}`)
			return
		}
		w.WriteHeader(200)
		io.WriteString(w, `{"message":"ok"}`)
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	ctx := chargebee.WithEnvironment(req.Context(), chargebee.Environment{
		RetryConfig: &chargebee.RetryConfig{
			Enabled:    true,
			MaxRetries: 2,
			DelayMs:    10,
			RetryOn:    map[int]struct{}{503: {}},
		},
	})
	req = req.WithContext(ctx)

	resp, err := chargebee.Do(req, false)
	if err != nil {
		t.Fatalf("expected success after retry, got error: %v", err)
	}

	if !bytes.Contains(resp.Body, []byte("ok")) {
		t.Errorf("expected success body, got: %s", resp.Body)
	}

	elapsed := time.Since(start)
	if elapsed < time.Second {
		t.Errorf("expected delay due to Retry-After header, got %v", elapsed)
	}
}

func TestDo_RetryDisabled(t *testing.T) {
	var callCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&callCount, 1)
		w.WriteHeader(503)
		io.WriteString(w, `{"type":"operation_failed","api_error_code":"disabled_retry"}`)
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	ctx := chargebee.WithEnvironment(req.Context(), chargebee.Environment{
		RetryConfig: &chargebee.RetryConfig{
			Enabled:    false,
			MaxRetries: 5,
			DelayMs:    10,
			RetryOn:    map[int]struct{}{503: {}},
		},
	})
	req = req.WithContext(ctx)

	_, err := chargebee.Do(req, false)
	if err == nil || !strings.Contains(err.Error(), "disabled_retry") {
		t.Errorf("expected error without retries, got: %v", err)
	}
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("expected 1 attempt, got: %d", callCount)
	}
}

func TestRequestWithEnv_RetryOverride(t *testing.T) {
	var count int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&count, 1)
		w.WriteHeader(503)
		io.WriteString(w, `{"type":"operation_failed","api_error_code":"temporary_failure"}`)
	}))
	defer server.Close()
	req := chargebee.Send("GET", "/customers", nil)
	mockClient := &http.Client{
		Transport: &mockTransport{server: server},
	}
	chargebee.WithHTTPClient(mockClient)

	env := chargebee.Environment{
		Key:      "test_key",
		SiteName: "test_site",
		RetryConfig: &chargebee.RetryConfig{
			Enabled:    true,
			MaxRetries: 3,
			DelayMs:    10,
			RetryOn:    map[int]struct{}{503: {}},
		},
	}

	_, err := req.RequestWithEnv(env)
	if err == nil || !strings.Contains(err.Error(), "operation_failed") {
		t.Errorf("expected retryable error, got: %v", err)
	}
	if atomic.LoadInt32(&count) != 4 { // 1 initial + 3 retries
		t.Errorf("expected 4 attempts, got: %d", count)
	}
}

type injectTraceAdapter struct{}

func (injectTraceAdapter) OnRequestStart(_ telemetry.RequestTelemetryContext, requestHeaders map[string]string) any {
	requestHeaders["traceparent"] = "00-test-trace"
	return nil
}

func (injectTraceAdapter) OnRequestEnd(any, telemetry.RequestTelemetryResult) {}

func snapshotHeaders(headers map[string]string) map[string]string {
	snapshot := make(map[string]string, len(headers))
	for key, value := range headers {
		snapshot[key] = value
	}
	return snapshot
}

func TestRequestWithEnv_DoesNotMutateRequestHeadersWithTelemetry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("traceparent") != "00-test-trace" {
			t.Errorf("expected traceparent on outbound request, got %q", r.Header.Get("traceparent"))
		}
		if r.Header.Get("chargebee-foo") != "bar" {
			t.Errorf("expected chargebee-foo on outbound request, got %q", r.Header.Get("chargebee-foo"))
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"list":[]}`)
	}))
	defer server.Close()

	chargebee.WithHTTPClient(&http.Client{Transport: &mockTransport{server: server}})

	env := chargebee.Environment{
		Key:              "test_key",
		SiteName:         "test_site",
		TelemetryAdapter: injectTraceAdapter{},
	}

	t.Run("RequestWithEnv", func(t *testing.T) {
		req := chargebee.Send("GET", "/customers/1", nil).
			Headers("chargebee-foo", "bar").
			WithTelemetryResource("customer").
			WithTelemetryOperation("retrieve")
		before := snapshotHeaders(req.Header)

		_, err := req.RequestWithEnv(env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if after := snapshotHeaders(req.Header); !headersEqual(before, after) {
			t.Fatalf("request headers mutated: before=%v after=%v", before, after)
		}
	})

	t.Run("ListRequestWithEnv", func(t *testing.T) {
		req := chargebee.SendList("GET", "/customers", nil).
			Headers("chargebee-foo", "bar").
			WithTelemetryResource("customer").
			WithTelemetryOperation("list")
		before := snapshotHeaders(req.Header)

		_, err := req.ListRequestWithEnv(env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if after := snapshotHeaders(req.Header); !headersEqual(before, after) {
			t.Fatalf("request headers mutated: before=%v after=%v", before, after)
		}
	})
}

func headersEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	return true
}

package chargebee

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

func TestDo_NilConfig(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	_, err := Do(req, true, nil)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDo_SuccessFirstTry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, `{"message":"ok"}`)
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := Do(req, true, &ClientConfig{})
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
	cfg := &ClientConfig{
		RetryConfig: &RetryConfig{
			Enabled:    true,
			MaxRetries: 2,
			DelayMs:    10,
			RetryOn:    []int{503},
		},
	}

	_, err := Do(req, true, cfg)
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
	cfg := &ClientConfig{
		RetryConfig: &RetryConfig{
			Enabled:    true,
			MaxRetries: 2,
			DelayMs:    10,
			RetryOn:    []int{503},
		},
	}

	resp, err := Do(req, false, cfg)
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
	cfg := &ClientConfig{
		RetryConfig: &RetryConfig{
			Enabled:    false,
			MaxRetries: 5,
			DelayMs:    10,
			RetryOn:    []int{503},
		},
	}

	_, err := Do(req, false, cfg)
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
	req, _ := http.NewRequest("GET", server.URL, nil)

	mockClient := &http.Client{
		Transport: &mockTransport{server: server},
	}

	cfg := &ClientConfig{
		ApiKey:   "test_key",
		SiteName: "test_site",
		RetryConfig: &RetryConfig{
			Enabled:    true,
			MaxRetries: 3,
			DelayMs:    10,
			RetryOn:    []int{503},
		},
		HTTPClient: mockClient,
	}

	_, err := Do(req, false, cfg)
	if err == nil || !strings.Contains(err.Error(), "operation_failed") {
		t.Errorf("expected retryable error, got: %v", err)
	}
	if atomic.LoadInt32(&count) != 4 { // 1 initial + 3 retries
		t.Errorf("expected 4 attempts, got: %d", count)
	}
}

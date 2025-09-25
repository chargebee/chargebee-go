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
)

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

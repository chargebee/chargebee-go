package tests

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/webhook"
	"github.com/stretchr/testify/assert"
)

func makeEventBody(eventType string, content string) []byte {
	if content == "" {
		content = "{}"
	}
	now := time.Now().Unix()
	body := []byte(`{
		"id":"evt_test_1",
		"occurred_at":` + toIntString(now) + `,
		"event_type":"` + eventType + `",
		"api_version":"v2",
		"content":` + content + `
	}`)
	return body
}

func toIntString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func TestParseEventType_OK(t *testing.T) {
	body := makeEventBody("pending_invoice_created", "{}")
	evtType, err := webhook.ParseEventType(body)
	assert.NoError(t, err)
	assert.Equal(t, enum.EventTypePendingInvoiceCreated, evtType)
}

func TestParseEventType_InvalidJSON(t *testing.T) {
	body := []byte(`invalid json`)
	_, err := webhook.ParseEventType(body)
	assert.Error(t, err)
}

func TestParseEventType_ApiVersionMismatch(t *testing.T) {
	body := []byte(`{"event_type":"pending_invoice_created","api_version":"v1"}`)
	_, err := webhook.ParseEventType(body)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API version")
}

func TestParseEventType_MissingEventType(t *testing.T) {
	body := []byte(`{"api_version":"v2"}`)
	evtType, err := webhook.ParseEventType(body)
	assert.NoError(t, err)
	assert.Equal(t, enum.EventType(""), evtType)
}

func TestHTTPHandler_RoutesToCallback(t *testing.T) {
	var called bool
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			t.Fatalf("unexpected error: %v", err)
		},
		OnPendingInvoiceCreated: func(e webhook.PendingInvoiceCreatedEvent) error {
			called = true
			assert.NotEmpty(t, e.Id)
			assert.NotEmpty(t, e.EventType)
			assert.NotZero(t, e.OccurredAt)
			assert.NotNil(t, e.Content)
			return nil
		},
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, called)
}

func TestHTTPHandler_ValidatorError(t *testing.T) {
	var onErrorCalled bool
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return errors.New("bad signature") },
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			onErrorCalled = true
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.True(t, onErrorCalled)
}

func TestHTTPHandler_CallbackError(t *testing.T) {
	var onErrorCalled bool
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			onErrorCalled = true
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
		OnPendingInvoiceCreated: func(e webhook.PendingInvoiceCreatedEvent) error {
			return errors.New("user code failed")
		},
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.True(t, onErrorCalled)
}

func TestHTTPHandler_UnknownEvent_OK(t *testing.T) {
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError:          func(w http.ResponseWriter, r *http.Request, err error) { t.Fatalf("unexpected error: %v", err) },
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("non_existing_event", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestHTTPHandler_NoHandlerRegistered_OK(t *testing.T) {
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError:          func(w http.ResponseWriter, r *http.Request, err error) { t.Fatalf("unexpected error: %v", err) },
		// No OnPendingInvoiceCreated handler registered
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestHTTPHandler_MultipleEventTypes(t *testing.T) {
	var pendingInvoiceCalled, subscriptionCalled bool
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError:          func(w http.ResponseWriter, r *http.Request, err error) { t.Fatalf("unexpected error: %v", err) },
		OnPendingInvoiceCreated: func(e webhook.PendingInvoiceCreatedEvent) error {
			pendingInvoiceCalled = true
			return nil
		},
		OnSubscriptionCreated: func(e webhook.SubscriptionCreatedEvent) error {
			subscriptionCalled = true
			return nil
		},
	}

	// Test pending_invoice_created
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)
	assert.True(t, pendingInvoiceCalled)
	assert.False(t, subscriptionCalled)

	// Test subscription_created
	pendingInvoiceCalled = false
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("subscription_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.False(t, pendingInvoiceCalled)
	assert.True(t, subscriptionCalled)
}

func TestHTTPHandler_InvalidJSONBody(t *testing.T) {
	var onErrorCalled bool
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			onErrorCalled = true
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader([]byte("invalid json")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.True(t, onErrorCalled)
}

func TestHTTPHandler_CustomValidator(t *testing.T) {
	var validatorCalled bool
	h := &webhook.WebhookHandler{
		RequestValidator: func(r *http.Request) error {
			validatorCalled = true
			// Custom validation logic - e.g., check headers
			if r.Header.Get("X-Custom-Header") != "expected-value" {
				return errors.New("missing required header")
			}
			return nil
		},
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		},
		OnPendingInvoiceCreated: func(e webhook.PendingInvoiceCreatedEvent) error {
			return nil
		},
	}

	// Test without required header
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusUnauthorized, rec1.Code)
	assert.True(t, validatorCalled)

	// Test with required header
	validatorCalled = false
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	req2.Header.Set("X-Custom-Header", "expected-value")
	h.HTTPHandler().ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.True(t, validatorCalled)
}


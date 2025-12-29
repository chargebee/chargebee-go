package webhook

import (
	"bytes"
	"encoding/base64"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/chargebee/chargebee-go/v4"
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
	evtType, err := ParseEventType(body)
	assert.NoError(t, err)
	assert.Equal(t, chargebee.EventTypePendingInvoiceCreated, evtType)
}

func TestParseEventType_InvalidJSON(t *testing.T) {
	body := []byte(`invalid json`)
	_, err := ParseEventType(body)
	assert.Error(t, err)
}

func TestParseEventType_ApiVersionMismatch(t *testing.T) {
	body := []byte(`{"event_type":"pending_invoice_created","api_version":"v1"}`)
	_, err := ParseEventType(body)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API version")
}

func TestParseEventType_MissingEventType(t *testing.T) {
	body := []byte(`{"api_version":"v2"}`)
	evtType, err := ParseEventType(body)
	assert.NoError(t, err)
	assert.Equal(t, chargebee.EventType(""), evtType)
}

func TestHTTPHandler_RoutesToCallback(t *testing.T) {
	var called bool
	h := &WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			t.Fatalf("unexpected error: %v", err)
		},
		OnPendingInvoiceCreated: func(e PendingInvoiceCreatedEvent) error {
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
	h := &WebhookHandler{
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
	h := &WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError: func(w http.ResponseWriter, r *http.Request, err error) {
			onErrorCalled = true
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
		OnPendingInvoiceCreated: func(e PendingInvoiceCreatedEvent) error {
			return errors.New("user code failed")
		},
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.True(t, onErrorCalled)
}

func TestHTTPHandler_UnknownEvent_Error(t *testing.T) {
	h := &WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		// Use default error handler which writes 500
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("non_existing_event", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestHTTPHandler_NoHandlerRegistered_Error(t *testing.T) {
	h := &WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		// No OnPendingInvoiceCreated handler registered; expect 500 via default error handler
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/cb", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	h.HTTPHandler().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestHTTPHandler_MultipleEventTypes(t *testing.T) {
	var pendingInvoiceCalled, subscriptionCalled bool
	h := &WebhookHandler{
		RequestValidator: func(r *http.Request) error { return nil },
		OnError:          func(w http.ResponseWriter, r *http.Request, err error) { t.Fatalf("unexpected error: %v", err) },
		OnPendingInvoiceCreated: func(e PendingInvoiceCreatedEvent) error {
			pendingInvoiceCalled = true
			return nil
		},
		OnSubscriptionCreated: func(e SubscriptionCreatedEvent) error {
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
	h := &WebhookHandler{
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
	h := &WebhookHandler{
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
		OnPendingInvoiceCreated: func(e PendingInvoiceCreatedEvent) error {
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

// Helper function to create Basic Auth header value
func makeBasicAuth(username, password string) string {
	credentials := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	return "Basic " + encoded
}

func TestBasicAuthValidator_ValidCredentials(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		return username == "testuser" && password == "testpass"
	})

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", makeBasicAuth("testuser", "testpass"))

	err := validator(req)
	assert.NoError(t, err)
}

func TestBasicAuthValidator_InvalidCredentials(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		return username == "testuser" && password == "testpass"
	})

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", makeBasicAuth("wronguser", "wrongpass"))

	err := validator(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid credentials")
}

func TestBasicAuthValidator_MissingAuthorizationHeader(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		return true
	})

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	// No Authorization header set

	err := validator(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid basic authorization header")
}

func TestBasicAuthValidator_InvalidScheme(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		return true
	})

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", "Bearer token123")

	err := validator(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid basic authorization header")
}

func TestBasicAuthValidator_InvalidBase64Encoding(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		return true
	})

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", "Basic invalid-base64!!!")

	err := validator(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid basic authorization header")
}

func TestBasicAuthValidator_InvalidCredentialsFormat(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		return true
	})

	// Encode a string without colon separator
	encoded := base64.StdEncoding.EncodeToString([]byte("nocolonseparator"))
	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", "Basic "+encoded)

	err := validator(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid basic authorization header")
}

func TestBasicAuthValidator_EmptyCredentialsAllowed(t *testing.T) {
	// When validateCredentials always returns true (empty credentials allowed)
	validator := BasicAuthValidator(func(username, password string) bool {
		return true // Allow all credentials
	})

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", makeBasicAuth("anyuser", "anypass"))

	err := validator(req)
	assert.NoError(t, err)
}

func TestBasicAuthValidator_UsernameWithColon(t *testing.T) {
	validator := BasicAuthValidator(func(username, password string) bool {
		// When encoding "user:name:pass:word", SplitN creates username="user" and password="name:pass:word"
		return username == "user" && password == "name:pass:word"
	})

	// Create auth header with colon in credentials
	credentials := "user:name:pass:word"
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	req.Header.Set("Authorization", "Basic "+encoded)

	err := validator(req)
	assert.NoError(t, err)
}

func TestBasicAuthErrorHandler_AuthError(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)

	// Test with missing Authorization header error
	err := errors.New("missing Authorization header")
	BasicAuthErrorHandler(rec, req, err)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, `Basic realm="Webhook"`, rec.Header().Get("WWW-Authenticate"))
	assert.Contains(t, rec.Body.String(), "unauthorized")
}

func TestBasicAuthErrorHandler_InvalidCredentialsError(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)

	err := errors.New("invalid credentials")
	BasicAuthErrorHandler(rec, req, err)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, `Basic realm="Webhook"`, rec.Header().Get("WWW-Authenticate"))
	assert.Contains(t, rec.Body.String(), "unauthorized")
}

func TestBasicAuthErrorHandler_InvalidSchemeError(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)

	err := errors.New("invalid authorization scheme, expected Basic")
	BasicAuthErrorHandler(rec, req, err)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, `Basic realm="Webhook"`, rec.Header().Get("WWW-Authenticate"))
	assert.Contains(t, rec.Body.String(), "unauthorized")
}

func TestBasicAuthErrorHandler_NonAuthError(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)

	// Test with a non-auth error
	err := errors.New("internal server error")
	BasicAuthErrorHandler(rec, req, err)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, `Basic realm="Webhook"`, rec.Header().Get("WWW-Authenticate"))
	assert.Contains(t, rec.Body.String(), "unauthorized")
}

func TestBasicAuthValidator_IntegrationWithWebhookHandler(t *testing.T) {
	var callbackCalled bool
	validator := BasicAuthValidator(func(username, password string) bool {
		return username == "admin" && password == "secret"
	})

	h := &WebhookHandler{
		RequestValidator: validator,
		OnError:          BasicAuthErrorHandler,
		OnPendingInvoiceCreated: func(e PendingInvoiceCreatedEvent) error {
			callbackCalled = true
			return nil
		},
	}

	// Test with valid credentials
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	req1.Header.Set("Authorization", makeBasicAuth("admin", "secret"))
	h.HTTPHandler().ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)
	assert.True(t, callbackCalled)

	// Test with invalid credentials
	callbackCalled = false
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(makeEventBody("pending_invoice_created", "{}")))
	req2.Header.Set("Authorization", makeBasicAuth("wrong", "wrong"))
	h.HTTPHandler().ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusUnauthorized, rec2.Code)
	assert.Equal(t, `Basic realm="Webhook"`, rec2.Header().Get("WWW-Authenticate"))
	assert.False(t, callbackCalled)
}

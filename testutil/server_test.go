package testutil

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chargebee "github.com/chargebee/chargebee-go/v4"
)

func TestServer_QueuedResponseAndRequestCapture(t *testing.T) {
	mock := NewServer()
	defer mock.Close()

	mock.Enqueue(MockResponse{
		StatusCode: http.StatusOK,
		Body:       []byte(`{"ok":true}`),
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Assert: func(req *http.Request, body []byte) error {
			if req.Method != http.MethodPost {
				return fmt.Errorf("expected POST, got %s", req.Method)
			}
			if req.URL.Path != "/api/v2/customers" {
				return fmt.Errorf("unexpected path: %s", req.URL.Path)
			}
			if string(body) != "id=cust_1" {
				return fmt.Errorf("unexpected body: %s", string(body))
			}
			return nil
		},
	})

	req, err := http.NewRequest(http.MethodPost, "https://example.chargebee.com/api/v2/customers", bytes.NewBufferString("id=cust_1"))
	require.NoError(t, err)

	resp, err := chargebee.Do(req, false, mock.ClientConfig())
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, `{"ok":true}`, string(resp.Body))
	assert.Nil(t, mock.AssertionError())
	assert.Equal(t, 0, mock.PendingResponses())

	last, ok := mock.LastRequest()
	require.True(t, ok)
	assert.Equal(t, http.MethodPost, last.Method)
	assert.Equal(t, "/api/v2/customers", last.Path)
	assert.Equal(t, "id=cust_1", string(last.Body))
}

func TestServer_EnqueueJSON(t *testing.T) {
	mock := NewServer()
	defer mock.Close()
	client := mock.NewClient()

	err := mock.EnqueueJSON(http.StatusOK, map[string]any{
		"customer": map[string]any{
			"id":              "cust_123",
			"first_name":      "John",
			"last_name":       "Doe",
			"email":           "john.doe@example.com",
			"auto_collection": "on",
			"billing_address": map[string]any{
				"city":              "New York",
				"state":             "NY",
				"country":           "USA",
				"zip":               "10001",
				"line1":             "123 Main St",
				"line2":             "Apt 4B",
				"line3":             "Suite 100",
				"validation_status": "valid",
			},
		},
	})

	require.NoError(t, err)

	res, err := client.Customer.Retrieve("cust_123")
	require.NoError(t, err)
	assert.Equal(t, "cust_123", res.Customer.Id)
	assert.Equal(t, "john.doe@example.com", res.Customer.Email)
	assert.Equal(t, chargebee.AutoCollectionOn, res.Customer.AutoCollection)
	assert.Equal(t, "New York", res.Customer.BillingAddress.City)
	assert.Equal(t, chargebee.ValidationStatusValid, res.Customer.BillingAddress.ValidationStatus)
}

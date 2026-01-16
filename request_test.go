package chargebee

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRoundTripper struct {
	response *http.Response
	err      error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

func TestPrepareJsonRequest(t *testing.T) {
	type testRequest struct {
		ID         string `json:"id"`
		apiRequest `json:"-"`
	}

	req := testRequest{
		ID: "123",
	}
	req.method = "POST"
	req.isJsonRequest = true

	err := req.prepare(req)
	assert.NoError(t, err)
	assert.Nil(t, req.urlParams)
	assert.Equal(t, `{"id":"123"}`, req.jsonBody)
}

func TestPrepareListRequest(t *testing.T) {
	type testRequest struct {
		Limit      *int32        `json:"limit,omitempty"`
		Offset     string        `json:"offset,omitempty"`
		Name       *StringFilter `json:"name,omitempty"`
		apiRequest `json:"-"`
	}

	req := &testRequest{
		Limit:  Int32(10),
		Offset: "100",
		Name: &StringFilter{
			StartsWith: "test",
		},
	}
	req.method = "GET"
	req.isListRequest = true

	err := req.prepare(req)
	assert.NoError(t, err)
	assert.Equal(t, &url.Values{
		"limit":             {"10"},
		"name[starts_with]": {"test"},
		"offset":            {"100"},
	}, req.urlParams)
}

func TestPrepareRequest(t *testing.T) {
	type testRequest struct {
		Limit      *int32        `json:"limit,omitempty"`
		Offset     string        `json:"offset,omitempty"`
		Name       *StringFilter `json:"name,omitempty"`
		apiRequest `json:"-"`
	}

	req := &testRequest{
		Limit:  Int32(10),
		Offset: "100",
		Name: &StringFilter{
			StartsWith: "test",
		},
	}
	req.method = "GET"

	err := req.prepare(req)
	assert.NoError(t, err)
	assert.Equal(t, &url.Values{
		"limit":             {"10"},
		"name[starts_with]": {"test"},
		"offset":            {"100"},
	}, req.urlParams)
}

func TestPrepareNilRequest(t *testing.T) {
	type testRequest struct {
		Limit      *int32        `json:"limit,omitempty"`
		Offset     string        `json:"offset,omitempty"`
		Name       *StringFilter `json:"name,omitempty"`
		apiRequest `json:"-"`
	}

	req := &testRequest{}
	req.method = "GET"

	err := req.prepare(nil)
	assert.NoError(t, err)
	assert.Nil(t, req.urlParams)
	assert.Equal(t, "", req.jsonBody)
}

func TestSendRequestWithCustomFields(t *testing.T) {
	// Mock response body with custom fields
	responseBody := `{
		"customer": {
			"id": "cust_123",
			"first_name": "John",
			"cf_string": "value_string",
			"cf_another": "value_another"
		}
	}`

	// Setup mock transport
	mockTransport := &mockRoundTripper{
		response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(responseBody)),
			Header:     make(http.Header),
		},
	}

	// Setup client config with mock client
	cfg := &ClientConfig{
		ApiKey:   "test_key",
		SiteName: "test_site",
		HTTPClient: &http.Client{
			Transport: mockTransport,
		},
	}

	// Create a request (using CustomerCreateRequest as it supports custom fields)
	req := &CustomerCreateRequest{}
	req.method = "POST"
	req.path = "/customers"

	// Execute request using send()
	// We expect a CustomerRetrieveResponse which contains the Customer object
	result, err := send[*CustomerRetrieveResponse](req, cfg)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Customer)
	assert.Equal(t, "cust_123", result.Customer.Id)

	// Verify custom fields
	// Note: CustomFields methods are not directly exposed on Customer but accessible via getter or field if public.
	// Customer struct has `CustomFields *customFields`.
	assert.NotNil(t, result.Customer.CustomFields)
	assert.Equal(t, "value_string", result.Customer.CustomFields.Get("cf_string"))
	assert.Equal(t, "value_another", result.Customer.CustomFields.Get("cf_another"))
}

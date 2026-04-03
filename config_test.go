package chargebee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientConfig(t *testing.T) {
	cfg := NewClientConfig("test_site", "test_api_key")
	assert.Equal(t, cfg.SiteName, "test_site")
	assert.Equal(t, cfg.ApiKey, "test_api_key")
	assert.Equal(t, cfg.HTTPClient, defaultHTTPClient)
	assert.Equal(t, cfg.RetryConfig, &RetryConfig{Enabled: true, MaxRetries: 3, DelayMs: 500, RetryOn: []int{500, 502, 503, 504}})
}

func TestApiBaseUrl(t *testing.T) {
	cfg := NewClientConfig("test_site", "test_api_key")
	cfg.Protocol = "http"
	cfg.ChargebeeDomain = "mock.com"
	assert.Equal(t, cfg.apiBaseUrl(""), "http://test_site.mock.com/api/v2")
	assert.Equal(t, cfg.apiBaseUrl("test_domain"), "http://test_site.test_domain.mock.com/api/v2")
}

func TestRequestHeaders(t *testing.T) {
	cfg := NewClientConfig("test_site", "test_api_key")
	cfg.RequestHeaders.Add("X-Test-Header", "test_value")
	assert.Equal(t, cfg.RequestHeaders.Get("X-Test-Header"), "test_value")
}

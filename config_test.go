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

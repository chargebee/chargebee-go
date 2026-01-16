package chargebee

import (
	"fmt"
	"net/http"
	"time"
)

type ClientConfig struct {
	ApiKey          string
	SiteName        string
	ChargebeeDomain string
	Protocol        string
	RetryConfig     *RetryConfig
	EnableDebugLogs bool
	HTTPClient      *http.Client
}

type RetryConfig struct {
	Enabled    bool
	MaxRetries int
	DelayMs    int
	RetryOn    []int
}

var (
	ConfigHTTPTimeout           = 80 * time.Second
	ConfigExportWaitInSecs      = 3 * time.Second
	ConfigTimeMachineWaitInSecs = 3 * time.Second
)

var defaultHTTPClient = &http.Client{Timeout: ConfigHTTPTimeout}

const (
	APIVersion = "v2"
	Charset    = "UTF-8"
)

type ctxKey string

const configCtxKey ctxKey = "config"

// NewClientConfig creates a new ClientConfig with the default HTTP client
// and retry config.
func NewClientConfig(siteName string, apiKey string) *ClientConfig {
	if siteName == "" || apiKey == "" {
		panic("Chargebee siteName and apiKey cannot be empty")
	}
	return &ClientConfig{
		SiteName:   siteName,
		ApiKey:     apiKey,
		HTTPClient: defaultHTTPClient,
		RetryConfig: &RetryConfig{
			Enabled:    true,
			MaxRetries: 3,
			DelayMs:    500,
			RetryOn:    []int{500, 502, 503, 504},
		},
	}
}

func (config *ClientConfig) apiBaseUrl(subDomain string) string {
	if config.Protocol == "" {
		config.Protocol = "https"
	}
	if config.ChargebeeDomain != "" {
		if subDomain != "" {
			return fmt.Sprintf("%v://%v.%v.%v/api/%v", config.Protocol, config.SiteName, subDomain, config.ChargebeeDomain, APIVersion)
		}
		return fmt.Sprintf("%v://%v.%v/api/%v", config.Protocol, config.SiteName, config.ChargebeeDomain, APIVersion)
	}
	if subDomain != "" {
		return fmt.Sprintf("https://%v.%v.chargebee.com/api/%v", config.SiteName, subDomain, APIVersion)
	}
	return fmt.Sprintf("https://%v.chargebee.com/api/%v", config.SiteName, APIVersion)
}

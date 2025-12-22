package chargebee

import (
	"fmt"
	"net/http"
	"time"
)

type ApiKey string
type SiteName string

type ClientConfig struct {
	ApiKey          ApiKey
	SiteName        SiteName
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
	RetryOn    map[int]struct{}
}

var (
	TotalHTTPTimeout      = 80 * time.Second
	ExportWaitInSecs      = 3 * time.Second
	TimeMachineWaitInSecs = 3 * time.Second
)

const (
	APIVersion = "v2"
	Charset    = "UTF-8"
)

type cbCtxKey string

const cbEnvKey cbCtxKey = "cb_env"

// func WithEnvironment(ctx context.Context, env ClientConfig) context.Context {
// 	return context.WithValue(ctx, cbEnvKey, env)
// }

// func Configure(key string, siteName string) {
// 	if key == "" || siteName == "" {
// 		panic(errors.New("Key/siteName cannot be empty"))
// 	}
// 	DefaultEnv = ClientConfig{Key: key, SiteName: siteName}
// }

// func WithHTTPClient(c *http.Client) {
// 	if c.Timeout == 0 {
// 		c.Timeout = TotalHTTPTimeout
// 	}
// 	httpClient = c
// }

// func WithRetryConfig(c *RetryConfig) {
// 	DefaultEnv.RetryConfig = c
// }

// func WithEnableDebugLogs(enabled bool) {
// 	DefaultEnv.EnableDebugLogs = enabled
// }

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

// func DefaultConfig() ClientConfig {
// 	if DefaultEnv == (ClientConfig{}) {
// 		panic(errors.New("The default environment has not been configured"))
// 	}
// 	return DefaultEnv
// }

// func NewDefaultHTTPClient() *http.Client {
// 	return &http.Client{Timeout: TotalHTTPTimeout}
// }

// func UpdateTotalHTTPTimeout(timeout time.Duration) {
// 	TotalHTTPTimeout = timeout
// 	if httpClient != nil {
// 		httpClient.Timeout = TotalHTTPTimeout
// 	}
// }

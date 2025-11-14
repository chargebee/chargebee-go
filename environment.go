package chargebee

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type RetryConfig struct {
	Enabled    bool
	MaxRetries int
	DelayMs    int
	RetryOn    map[int]struct{}
}

type Environment struct {
	Key             string
	SiteName        string
	ChargebeeDomain string
	Protocol        string
	RetryConfig     *RetryConfig
	EnableDebugLogs bool
}

type cbCtxKey string

var (
	TotalHTTPTimeout      = 80 * time.Second
	ExportWaitInSecs      = 3 * time.Second
	TimeMachineWaitInSecs = 3 * time.Second
	DefaultEnv            Environment
	httpClient            *http.Client
)

const (
	APIVersion = "v2"
	Charset    = "UTF-8"
)

const cbEnvKey cbCtxKey = "cb_env"

func WithEnvironment(ctx context.Context, env Environment) context.Context {
	return context.WithValue(ctx, cbEnvKey, env)
}

func Configure(key string, siteName string) error {
	if key == "" || siteName == "" {
		return errors.New("Key/siteName cannot be empty")
	}
	DefaultEnv = Environment{Key: key, SiteName: siteName}
	return nil
}
func WithHTTPClient(c *http.Client) {
	if c.Timeout == 0 {
		c.Timeout = TotalHTTPTimeout
	}
	httpClient = c
}
func WithRetryConfig(c *RetryConfig) {
	DefaultEnv.RetryConfig = c
}

func WithEnableDebugLogs(enabled bool) {
	DefaultEnv.EnableDebugLogs = enabled
}

func (env *Environment) apiBaseUrl(subDomain string) string {
	if env.Protocol == "" {
		env.Protocol = "https"
	}
	if env.ChargebeeDomain != "" {
		if subDomain != "" {
			return fmt.Sprintf("%v://%v.%v.%v/api/%v", env.Protocol, env.SiteName, subDomain, env.ChargebeeDomain, APIVersion)
		}
		return fmt.Sprintf("%v://%v.%v/api/%v", env.Protocol, env.SiteName, env.ChargebeeDomain, APIVersion)
	}
	if subDomain != "" {
		return fmt.Sprintf("https://%v.%v.chargebee.com/api/%v", env.SiteName, subDomain, APIVersion)
	}
	return fmt.Sprintf("https://%v.chargebee.com/api/%v", env.SiteName, APIVersion)
}

func DefaultConfig() (Environment, error) {
	if DefaultEnv == (Environment{}) {
		return Environment{}, errors.New("The default environment has not been configured")
	}
	return DefaultEnv, nil
}

func NewDefaultHTTPClient() *http.Client {
	return &http.Client{Timeout: TotalHTTPTimeout}
}

func UpdateTotalHTTPTimeout(timeout time.Duration) {
	TotalHTTPTimeout = timeout
	if httpClient != nil {
		httpClient.Timeout = TotalHTTPTimeout
	}
}

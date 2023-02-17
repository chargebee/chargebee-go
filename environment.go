package chargebee

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Environment struct {
	Key             string
	SiteName        string
	ChargebeeDomain string
	Protocol        string
}

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

func Configure(key string, siteName string) {
	if key == "" || siteName == "" {
		panic(errors.New("Key/siteName cannot be empty"))
	}
	DefaultEnv = Environment{Key: key, SiteName: siteName}
}
func WithHTTPClient(c *http.Client) {
	if c.Timeout == 0 {
		c.Timeout = TotalHTTPTimeout
	}
	httpClient = c
}
func (env *Environment) apiBaseUrl() string {
	if env.Protocol == "" {
		env.Protocol = "https"
	}
	if env.ChargebeeDomain != "" {
		return fmt.Sprintf("%v://%v.%v/api/%v", env.Protocol, env.SiteName, env.ChargebeeDomain, APIVersion)
	}
	return fmt.Sprintf("https://%v.chargebee.com/api/%v", env.SiteName, APIVersion)
}

func DefaultConfig() Environment {
	if DefaultEnv == (Environment{}) {
		panic(errors.New("The default environment has not been configured"))
	}
	return DefaultEnv
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

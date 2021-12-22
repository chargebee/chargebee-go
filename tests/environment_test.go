package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/chargebee/chargebee-go"
	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	chargebee.Configure("Your site Key", "Your site Name")
	defaultEnv := chargebee.DefaultConfig()
	assert.Equal(t, defaultEnv.SiteName, "Your site Name")
	assert.Equal(t, defaultEnv.Key, "Your site Key")
}

func TestDefaultHTTPClient(t *testing.T) {
	c := chargebee.NewDefaultHTTPClient()
	chargebee.WithHTTPClient(c)
	assert.Equal(t, chargebee.DefaultHTTPTimeout, c.Timeout)

	// UpdateTotalHTTPTimeout should set the timeout
	// on new and the existing http.Clients
	// as TotalHTTPTimeout is only read on http.Client creation
	timeout := 3 * time.Second
	chargebee.UpdateTotalHTTPTimeout(timeout)
	c2 := chargebee.NewDefaultHTTPClient()
	assert.Equal(t, timeout, c2.Timeout)
	assert.Equal(t, c2.Timeout, c.Timeout)

	// http.Client should always have a timeout
	c3 := &http.Client{}
	chargebee.WithHTTPClient(c3)
	assert.Equal(t, timeout, c3.Timeout)

	// respect the timeout passed in?
	timeout2 := 5 * time.Second
	c4 := &http.Client{Timeout: timeout2}
	chargebee.WithHTTPClient(c4)
	assert.Equal(t, timeout2, c4.Timeout)
	assert.NotEqual(t, timeout, c4.Timeout)
}

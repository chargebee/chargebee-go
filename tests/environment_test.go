package tests

import (
	"github.com/chargebee/chargebee-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnv(t *testing.T) {
	chargebee.Configure("Your site Key", "Your site Name")
	defaultEnv := chargebee.DefaultConfig()
	assert.Equal(t, defaultEnv.SiteName, "Your site Name")
	assert.Equal(t, defaultEnv.Key, "Your site Key")
}

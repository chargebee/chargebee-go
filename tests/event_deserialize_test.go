package tests

import (
	"testing"

	"github.com/chargebee/chargebee-go/v3/actions/event"
	"github.com/stretchr/testify/assert"
)

func TestDeserialize(t *testing.T) {
	evt := eventData()
	res := event.Deserialize(evt)
	assert.NotNil(t, res.Content)
	content := event.Content(*res)
	assert.NotNil(t, content.Customer)
	assert.NotNil(t, content.Subscription)
	assert.NotNil(t, content.Card)
	assert.NotNil(t, content.Invoice)
	assert.NotNil(t, content.Transaction)

	// Check that custom fields are here
	assert.Equal(t, content.Subscription.CustomField["cf_subscription_comment"], "test_value_subscription")
	assert.Equal(t, content.Customer.CustomField["cf_customer_comment"], "test_value_customer")
}

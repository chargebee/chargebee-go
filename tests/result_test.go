package tests

import (
	"github.com/chargebee/chargebee-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleResponseToStructConversion(t *testing.T) {
	data := simpleSubscription()
	result := &chargebee.Result{}
	chargebee.UnmarshalJSON(data, result)
	s := result.Subscription
	c := result.Customer
	assert.Equal(t, "simple_subscription", s.Id)
	assert.Equal(t, "basic", s.PlanId)
	assert.Equal(t, "simple", c.FirstName)
	assert.Equal(t, "subscription", c.LastName)
}

func TestNestedResponseToStructConversion(t *testing.T) {
	data := nestedSubscription()
	result := &chargebee.Result{}
	chargebee.UnmarshalJSON(data, result)
	s := result.Subscription
	assert.Equal(t, "nested_subscription", s.Id)
	addons := s.Addons
	assert.Equal(t, 2, len(addons))
	for i := range addons {
		assert.Contains(t, []string{"monitor", "ssl"}, addons[i].Id)
	}
}

package tests

import (
	"testing"

	"github.com/chargebee/chargebee-go"
	"github.com/stretchr/testify/assert"
)

func TestResponseToListConversion(t *testing.T) {
	data := listSubscription()
	listResult := &chargebee.ResultList{}
	chargebee.UnmarshalJSON(data, listResult)
	assert.Equal(t, 2, len(listResult.List))
	for i := range listResult.List {
		assert.Equal(t, "simple_subscription", listResult.List[i].Subscription.Id)
	}
}

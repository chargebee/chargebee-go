package tests

import (
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/models/subscription"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestUtil(t *testing.T) {

	before := &subscription.CreateRequestParams{
		PlanId: "cbdemo_grow",
		Customer: &subscription.CreateCustomerParams{
			Email:          "john@user.com",
			FirstName:      "John",
			LastName:       "Doe",
			Locale:         "frCA",
			Phone:          "+19499999999",
			AutoCollection: enum.AutoCollectionOn,
		},
		Addons: []*subscription.CreateAddonParams{
			{
				Id: "cbdemo_conciergesupport",
			},
			{
				Id:       "cbdemo_additionaluser",
				Quantity: chargebee.Int32(2),
			},
		},
		CouponIds: []string{"cbdemo_earlybird"},
	}
	after := &url.Values{
		"coupon_ids[0]":             []string{"cbdemo_earlybird"},
		"customer[phone]":           []string{"+19499999999"},
		"plan_id":                   []string{"cbdemo_grow"},
		"customer[last_name]":       []string{"Doe"},
		"customer[locale]":          []string{"frCA"},
		"addons[id][0]":             []string{"cbdemo_conciergesupport"},
		"addons[id][1]":             []string{"cbdemo_additionaluser"},
		"addons[quantity][1]":       []string{"2"},
		"customer[email]":           []string{"john@user.com"},
		"customer[auto_collection]": []string{"on"},
		"customer[first_name]":      []string{"John"},
	}
	assert.Equal(t, chargebee.SerializeParams(before), after)
}

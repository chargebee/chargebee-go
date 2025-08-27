package pricingpagesession

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/pricingpagesession"
)

func CreateForNewSubscription(params *pricingpagesession.CreateForNewSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/pricing_page_sessions/create_for_new_subscription"), params).SetIdempotency(true)
}
func CreateForExistingSubscription(params *pricingpagesession.CreateForExistingSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/pricing_page_sessions/create_for_existing_subscription"), params).SetIdempotency(true)
}

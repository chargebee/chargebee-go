package ramp

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/ramp"
	"net/url"
)

func CreateForSubscription(id string, params *ramp.CreateForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/create_ramp", url.PathEscape(id)), params).SetIdempotency(true)
}
func Update(id string, params *ramp.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/ramps/%v/update", url.PathEscape(id)), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/ramps/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/ramps/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func List(params *ramp.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/ramps"), params)
}

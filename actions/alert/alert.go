package alert

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/alert"
	"net/url"
)

func Create(params *alert.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/alerts"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/alerts/%v", url.PathEscape(id)), nil)
}
func List(params *alert.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/alerts"), params)
}
func Update(id string, params *alert.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/alerts/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/alerts/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func ApplicationAlertsForSubscription(id string, params *alert.ApplicationAlertsForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/applicable_alerts", url.PathEscape(id)), params)
}

package omnichannelsubscription

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"
	"net/url"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/omnichannel_subscriptions/%v", url.PathEscape(id)), nil)
}
func List(params *omnichannelsubscription.ListRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/omnichannel_subscriptions"), params)
}
func OmnichannelTransactionsForOmnichannelSubscription(id string, params *omnichannelsubscription.OmnichannelTransactionsForOmnichannelSubscriptionRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/omnichannel_subscriptions/%v/omnichannel_transactions", url.PathEscape(id)), params)
}

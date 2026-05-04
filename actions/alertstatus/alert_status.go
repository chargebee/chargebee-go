package alertstatus

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/alertstatus"
	"net/url"
)

func AlertStatusesForSubscription(id string, params *alertstatus.AlertStatusesForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/alert_statuses", url.PathEscape(id)), params)
}
func AlertStatusesForAlert(id string, params *alertstatus.AlertStatusesForAlertRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/alerts/%v/alert_statuses", url.PathEscape(id)), params)
}

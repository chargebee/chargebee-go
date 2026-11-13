package gatewaywebhooklog

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/gatewaywebhooklog"
	"net/url"
)

func List(params *gatewaywebhooklog.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/gateway_webhook_logs"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/gateway_webhook_logs/%v", url.PathEscape(id)), nil)
}

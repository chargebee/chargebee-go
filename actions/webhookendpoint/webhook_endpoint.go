package webhookendpoint

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/webhookendpoint"
	"net/url"
)

func Create(params *webhookendpoint.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/webhook_endpoints"), params).SetIdempotency(true)
}
func Update(id string, params *webhookendpoint.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/webhook_endpoints/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/webhook_endpoints/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/webhook_endpoints/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func List(params *webhookendpoint.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/webhook_endpoints"), params)
}

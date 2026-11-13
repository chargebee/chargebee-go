package gatewayrequestlog

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/gatewayrequestlog"
	"net/url"
)

func List(params *gatewayrequestlog.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/gateway_request_logs"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/gateway_request_logs/%v", url.PathEscape(id)), nil)
}

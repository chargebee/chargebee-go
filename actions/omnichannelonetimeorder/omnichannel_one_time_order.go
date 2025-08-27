package omnichannelonetimeorder

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorder"
	"net/url"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/omnichannel_one_time_orders/%v", url.PathEscape(id)), nil)
}
func List(params *omnichannelonetimeorder.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/omnichannel_one_time_orders"), params)
}

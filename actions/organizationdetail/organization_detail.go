package organizationdetail

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/organizationdetail"
	"net/url"
)

func Create(params *organizationdetail.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/organization_details"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/organization_details/%v", url.PathEscape(id)), nil)
}
func List(params *organizationdetail.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/organization_details"), params)
}

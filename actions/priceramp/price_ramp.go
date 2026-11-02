package priceramp

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/priceramp"
	"net/url"
)

func Create(params *priceramp.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/price_ramps"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/price_ramps/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *priceramp.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/price_ramps/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/price_ramps/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func List(params *priceramp.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/price_ramps"), params)
}

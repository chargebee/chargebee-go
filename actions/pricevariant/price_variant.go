package pricevariant

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/pricevariant"
	"net/url"
)

func Create(params *pricevariant.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/price_variants"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/price_variants/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *pricevariant.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/price_variants/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/price_variants/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func List(params *pricevariant.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/price_variants"), params)
}

package pricevariant

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/pricevariant"
	"net/url"
)

func Create(params *pricevariant.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/price_variants"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/price_variants/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *pricevariant.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/price_variants/%v", url.PathEscape(id)), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/price_variants/%v/delete", url.PathEscape(id)), nil)
}
func List(params *pricevariant.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/price_variants"), params)
}

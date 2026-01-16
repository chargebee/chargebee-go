package itemprice

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/itemprice"
	"net/url"
)

func Create(params *itemprice.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/item_prices/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *itemprice.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *itemprice.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/item_prices"), params)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func FindApplicableItems(id string, params *itemprice.FindApplicableItemsRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/item_prices/%v/applicable_items", url.PathEscape(id)), params)
}
func FindApplicableItemPrices(id string, params *itemprice.FindApplicableItemPricesRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/item_prices/%v/applicable_item_prices", url.PathEscape(id)), params)
}

package itemprice

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/itemprice"
)

func Create(params *itemprice.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/item_prices/%v", chargebee.IDEscape(id)), nil)
}
func Update(id string, params *itemprice.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices/%v", chargebee.IDEscape(id)), params)
}
func List(params *itemprice.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/item_prices"), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices/%v/delete", chargebee.IDEscape(id)), nil)
}
func FindApplicableItems(id string, params *itemprice.FindApplicableItemsRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/item_prices/%v/applicable_items", chargebee.IDEscape(id)), params)
}
func FindApplicableItemPrices(id string, params *itemprice.FindApplicableItemPricesRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/item_prices/%v/applicable_item_prices", chargebee.IDEscape(id)), params)
}

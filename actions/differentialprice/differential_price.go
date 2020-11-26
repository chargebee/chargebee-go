package differentialprice

import (
	"fmt"

	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/differentialprice"
)

func Create(id string, params *differentialprice.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices/%v/differential_prices", id), params)
}
func Retrieve(id string, params *differentialprice.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/differential_prices/%v", id), params)
}
func Update(id string, params *differentialprice.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/differential_prices/%v", id), params)
}
func Delete(id string, params *differentialprice.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/differential_prices/%v/delete", id), params)
}
func List(params *differentialprice.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/differential_prices"), params)
}

package differentialprice

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/differentialprice"
	"net/url"
)

func Create(id string, params *differentialprice.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_prices/%v/differential_prices", url.PathEscape(id)), params).SetIdempotency(true)
}
func Retrieve(id string, params *differentialprice.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/differential_prices/%v", url.PathEscape(id)), params)
}
func Update(id string, params *differentialprice.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/differential_prices/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string, params *differentialprice.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/differential_prices/%v/delete", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *differentialprice.ListRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/differential_prices"), params)
}

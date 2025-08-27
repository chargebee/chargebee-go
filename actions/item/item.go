package item

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/item"
	"net/url"
)

func Create(params *item.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/items"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/items/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *item.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *item.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/items"), params)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

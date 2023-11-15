package item

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/item"
)

func Create(params *item.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/items/%v", chargebee.IDEscape(id)), nil)
}
func Update(id string, params *item.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v", chargebee.IDEscape(id)), params)
}
func List(params *item.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/items"), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v/delete", chargebee.IDEscape(id)), nil)
}

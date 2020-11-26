package itemfamily

import (
	"fmt"

	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/itemfamily"
)

func Create(params *itemfamily.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_families"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/item_families/%v", id), nil)
}
func List(params *itemfamily.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/item_families"), params)
}
func Update(id string, params *itemfamily.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_families/%v", id), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/item_families/%v/delete", id), nil)
}

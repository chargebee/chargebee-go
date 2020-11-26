package attacheditem

import (
	"fmt"

	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/attacheditem"
)

func Create(id string, params *attacheditem.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v/attached_items", id), params)
}
func Update(id string, params *attacheditem.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/attached_items/%v", id), params)
}
func Retrieve(id string, params *attacheditem.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/attached_items/%v", id), params)
}
func Delete(id string, params *attacheditem.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/attached_items/%v/delete", id), params)
}
func List(id string, params *attacheditem.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/items/%v/attached_items", id), params)
}
func ListInternal(params *attacheditem.ListInternalRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/attached_items/list_internal"), params)
}

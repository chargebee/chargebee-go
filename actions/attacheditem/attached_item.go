package attacheditem

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/attacheditem"
	"net/url"
)

func Create(id string, params *attacheditem.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v/attached_items", url.PathEscape(id)), params)
}
func Update(id string, params *attacheditem.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/attached_items/%v", url.PathEscape(id)), params)
}
func Retrieve(id string, params *attacheditem.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/attached_items/%v", url.PathEscape(id)), params)
}
func Delete(id string, params *attacheditem.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/attached_items/%v/delete", url.PathEscape(id)), params)
}
func List(id string, params *attacheditem.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/items/%v/attached_items", url.PathEscape(id)), params)
}

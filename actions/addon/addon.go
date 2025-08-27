package addon

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/addon"
	"net/url"
)

func Create(params *addon.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons"), params).SetIdempotency(true)
}
func Update(id string, params *addon.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *addon.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/addons"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/addons/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Copy(params *addon.CopyRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/copy"), params).SetIdempotency(true)
}
func Unarchive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v/unarchive", url.PathEscape(id)), nil).SetIdempotency(true)
}

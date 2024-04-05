package addon

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/addon"
)

func Create(params *addon.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/addons"), params)
}
func Update(id string, params *addon.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v", chargebee.IDEscape(id)), params)
}
func List(params *addon.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/addons"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/addons/%v", chargebee.IDEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v/delete", chargebee.IDEscape(id)), nil)
}
func Copy(params *addon.CopyRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/addons/copy"), params)
}
func Unarchive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v/unarchive", chargebee.IDEscape(id)), nil)
}

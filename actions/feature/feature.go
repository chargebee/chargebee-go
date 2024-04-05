package feature

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/feature"
)

func List(params *feature.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/features"), params)
}
func Create(params *feature.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features"), params)
}
func Update(id string, params *feature.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v", chargebee.IDEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/features/%v", chargebee.IDEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/delete", chargebee.IDEscape(id)), nil)
}
func Activate(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/activate_command", chargebee.IDEscape(id)), nil)
}
func Archive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/archive_command", chargebee.IDEscape(id)), nil)
}
func Reactivate(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/reactivate_command", chargebee.IDEscape(id)), nil)
}

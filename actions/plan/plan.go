package plan

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/plan"
)

func Create(params *plan.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans"), params)
}
func Update(id string, params *plan.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v", chargebee.IDEscape(id)), params)
}
func List(params *plan.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/plans"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/plans/%v", chargebee.IDEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v/delete", chargebee.IDEscape(id)), nil)
}
func Copy(params *plan.CopyRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/copy"), params)
}
func Unarchive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v/unarchive", chargebee.IDEscape(id)), nil)
}

package plan

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/plan"
	"net/url"
)

func Create(params *plan.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans"), params)
}
func Update(id string, params *plan.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v", url.PathEscape(id)), params)
}
func List(params *plan.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/plans"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/plans/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v/delete", url.PathEscape(id)), nil)
}
func Copy(params *plan.CopyRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/copy"), params)
}
func Unarchive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v/unarchive", url.PathEscape(id)), nil)
}

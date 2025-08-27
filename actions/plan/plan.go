package plan

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/plan"
	"net/url"
)

func Create(params *plan.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/plans"), params).SetIdempotency(true)
}
func Update(id string, params *plan.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *plan.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/plans"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/plans/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Copy(params *plan.CopyRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/plans/copy"), params).SetIdempotency(true)
}
func Unarchive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/plans/%v/unarchive", url.PathEscape(id)), nil).SetIdempotency(true)
}

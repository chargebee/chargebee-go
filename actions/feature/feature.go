package feature

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/feature"
	"net/url"
)

func List(params *feature.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/features"), params)
}
func Create(params *feature.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features"), params).SetIdempotency(true)
}
func Update(id string, params *feature.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/features/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Activate(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/activate_command", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Archive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/archive_command", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Reactivate(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/reactivate_command", url.PathEscape(id)), nil).SetIdempotency(true)
}

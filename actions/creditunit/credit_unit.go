package creditunit

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/creditunit"
	"net/url"
)

func List(params *creditunit.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/credit_units"), params)
}
func Create(params *creditunit.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/credit_units"), params).SetIdempotency(true)
}
func Update(id string, params *creditunit.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/credit_units/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Archive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/credit_units/%v/archive_command", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Reactivate(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/credit_units/%v/reactivate_command", url.PathEscape(id)), nil).SetIdempotency(true)
}

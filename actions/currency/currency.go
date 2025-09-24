package currency

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/currency"
	"net/url"
)

func List(params *currency.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/currencies/list"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/currencies/%v", url.PathEscape(id)), nil)
}
func Create(params *currency.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/currencies"), params).SetIdempotency(true)
}
func Update(id string, params *currency.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/currencies/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func AddSchedule(id string, params *currency.AddScheduleRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/currencies/%v/add_schedule", url.PathEscape(id)), params).SetIdempotency(true)
}
func RemoveSchedule(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/currencies/%v/remove_schedule", url.PathEscape(id)), nil).SetIdempotency(true)
}

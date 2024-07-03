package currency

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/currency"
	"net/url"
)

func List() chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/currencies/list"), nil)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/currencies/%v", url.PathEscape(id)), nil)
}
func Create(params *currency.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/currencies"), params)
}
func Update(id string, params *currency.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/currencies/%v", url.PathEscape(id)), params)
}
func AddSchedule(id string, params *currency.AddScheduleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/currencies/%v/add_schedule", url.PathEscape(id)), params)
}
func RemoveSchedule(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/currencies/%v/remove_schedule", url.PathEscape(id)), nil)
}

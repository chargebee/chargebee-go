package paymentschedule

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentschedule"
)

func List(params *paymentschedule.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/payment_schedules"), params)
}

package promotionalcredit

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"
	"net/url"
)

func Add(params *promotionalcredit.AddRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/add"), params).SetIdempotency(true)
}
func Deduct(params *promotionalcredit.DeductRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/deduct"), params).SetIdempotency(true)
}
func Set(params *promotionalcredit.SetRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/set"), params).SetIdempotency(true)
}
func List(params *promotionalcredit.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/promotional_credits"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/promotional_credits/%v", url.PathEscape(id)), nil)
}

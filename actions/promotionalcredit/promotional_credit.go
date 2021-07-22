package promotionalcredit

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/promotionalcredit"
)

func Add(params *promotionalcredit.AddRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/add"), params)
}
func Deduct(params *promotionalcredit.DeductRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/deduct"), params)
}
func Set(params *promotionalcredit.SetRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/set"), params)
}
func List(params *promotionalcredit.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/promotional_credits"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/promotional_credits/%v", id), nil)
}

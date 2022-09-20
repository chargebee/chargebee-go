package purchase

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/purchase"
)

func Create(params *purchase.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases"), params)
}
func Estimate(params *purchase.EstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases/estimate"), params)
}

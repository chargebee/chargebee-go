package purchase

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/purchase"
)

func Create(params *purchase.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases"), params).SetIdempotency(true)
}
func Estimate(params *purchase.EstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases/estimate"), params)
}

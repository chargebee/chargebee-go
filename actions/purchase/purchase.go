package purchase

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/purchase"
	"net/url"
)

func Create(params *purchase.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/purchases/%v", url.PathEscape(id)), nil)
}
func Estimate(params *purchase.EstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases/estimate"), params)
}

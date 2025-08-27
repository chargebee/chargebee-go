package recordedpurchase

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/recordedpurchase"
	"net/url"
)

func Create(params *recordedpurchase.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/recorded_purchases"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/recorded_purchases/%v", url.PathEscape(id)), nil)
}

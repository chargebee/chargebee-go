package recordedpurchase

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/recordedpurchase"
	"net/url"
)

func Create(params *recordedpurchase.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/recorded_purchases"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/recorded_purchases/%v", url.PathEscape(id)), nil)
}

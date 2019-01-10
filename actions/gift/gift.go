package gift

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/gift"
)

func Create(params *gift.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/gifts/%v", id), nil)
}
func List(params *gift.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/gifts"), params)
}
func Claim(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/claim", id), nil)
}
func Cancel(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/cancel", id), nil)
}

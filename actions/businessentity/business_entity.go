package businessentity

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/businessentity"
	"net/url"
)

func Create(params *businessentity.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_entities"), params).SetIdempotency(true)
}
func List(params *businessentity.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/business_entities"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/business_entities/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *businessentity.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_entities/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func CreateTransfers(params *businessentity.CreateTransfersRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_entities/transfers"), params).SetIdempotency(true)
}
func GetTransfers(params *businessentity.GetTransfersRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/business_entities/transfers"), params)
}

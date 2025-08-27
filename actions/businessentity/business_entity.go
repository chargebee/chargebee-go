package businessentity

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/businessentity"
)

func CreateTransfers(params *businessentity.CreateTransfersRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_entities/transfers"), params).SetIdempotency(true)
}
func GetTransfers(params *businessentity.GetTransfersRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/business_entities/transfers"), params)
}

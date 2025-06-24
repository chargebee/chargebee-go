package businessentity

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/businessentity"
)

func CreateTransfers(params *businessentity.CreateTransfersRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/business_entities/transfers"), params).SetIdempotency(true)
}
func GetTransfers(params *businessentity.GetTransfersRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/business_entities/transfers"), params)
}

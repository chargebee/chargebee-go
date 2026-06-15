package promotionalgrant

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/promotionalgrant"
)

func PromotionalGrants(params *promotionalgrant.PromotionalGrantsRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/promotional_grants"), params).SetIdempotency(true)
}

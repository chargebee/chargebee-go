package personalizedoffer

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/personalizedoffer"
)

func PersonalizedOffers(params *personalizedoffer.PersonalizedOffersRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/personalized_offers"), params).SetSubDomain("grow")
}

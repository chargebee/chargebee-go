package gift

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/gift"
	"net/url"
)

func Create(params *gift.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts"), params).SetIdempotency(true)
}
func CreateForItems(params *gift.CreateForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/create_for_items"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/gifts/%v", url.PathEscape(id)), nil)
}
func List(params *gift.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/gifts"), params)
}
func Claim(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/claim", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Cancel(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/cancel", url.PathEscape(id)), nil).SetIdempotency(true)
}
func UpdateGift(id string, params *gift.UpdateGiftRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/update_gift", url.PathEscape(id)), params).SetIdempotency(true)
}

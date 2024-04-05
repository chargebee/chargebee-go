package gift

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/gift"
)

func Create(params *gift.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts"), params)
}
func CreateForItems(params *gift.CreateForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/create_for_items"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/gifts/%v", chargebee.IDEscape(id)), nil)
}
func List(params *gift.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/gifts"), params)
}
func Claim(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/claim", chargebee.IDEscape(id)), nil)
}
func Cancel(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/cancel", chargebee.IDEscape(id)), nil)
}
func UpdateGift(id string, params *gift.UpdateGiftRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/update_gift", chargebee.IDEscape(id)), params)
}

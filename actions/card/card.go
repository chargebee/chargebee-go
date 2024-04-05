package card

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/card"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/cards/%v", chargebee.IDEscape(id)), nil)
}
func UpdateCardForCustomer(id string, params *card.UpdateCardForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/credit_card", chargebee.IDEscape(id)), params)
}
func SwitchGatewayForCustomer(id string, params *card.SwitchGatewayForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/switch_gateway", chargebee.IDEscape(id)), params)
}
func CopyCardForCustomer(id string, params *card.CopyCardForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/copy_card", chargebee.IDEscape(id)), params)
}
func DeleteCardForCustomer(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_card", chargebee.IDEscape(id)), nil)
}

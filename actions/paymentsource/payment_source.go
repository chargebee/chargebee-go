package paymentsource

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentsource"
	"net/url"
)

func CreateUsingTempToken(params *paymentsource.CreateUsingTempTokenRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_temp_token"), params)
}
func CreateUsingPermanentToken(params *paymentsource.CreateUsingPermanentTokenRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_permanent_token"), params)
}
func CreateUsingToken(params *paymentsource.CreateUsingTokenRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_token"), params)
}
func CreateUsingPaymentIntent(params *paymentsource.CreateUsingPaymentIntentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_payment_intent"), params)
}
func CreateCard(params *paymentsource.CreateCardRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_card"), params)
}
func CreateBankAccount(params *paymentsource.CreateBankAccountRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_bank_account"), params)
}
func UpdateCard(id string, params *paymentsource.UpdateCardRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/update_card", url.PathEscape(id)), params)
}
func UpdateBankAccount(id string, params *paymentsource.UpdateBankAccountRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/update_bank_account", url.PathEscape(id)), params)
}
func VerifyBankAccount(id string, params *paymentsource.VerifyBankAccountRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/verify_bank_account", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/payment_sources/%v", url.PathEscape(id)), nil)
}
func List(params *paymentsource.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/payment_sources"), params)
}
func SwitchGatewayAccount(id string, params *paymentsource.SwitchGatewayAccountRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/switch_gateway_account", url.PathEscape(id)), params)
}
func ExportPaymentSource(id string, params *paymentsource.ExportPaymentSourceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/export_payment_source", url.PathEscape(id)), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/delete", url.PathEscape(id)), nil)
}
func DeleteLocal(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/delete_local", url.PathEscape(id)), nil)
}

package paymentsource

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentsource"
	"net/url"
)

func CreateUsingTempToken(params *paymentsource.CreateUsingTempTokenRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_temp_token"), params).SetIdempotency(true)
}
func CreateUsingPermanentToken(params *paymentsource.CreateUsingPermanentTokenRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_permanent_token"), params).SetIdempotency(true)
}
func CreateUsingToken(params *paymentsource.CreateUsingTokenRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_token"), params).SetIdempotency(true)
}
func CreateUsingPaymentIntent(params *paymentsource.CreateUsingPaymentIntentRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_using_payment_intent"), params).SetIdempotency(true)
}
func CreateVoucherPaymentSource(params *paymentsource.CreateVoucherPaymentSourceRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_voucher_payment_source"), params).SetIdempotency(true)
}
func CreateCard(params *paymentsource.CreateCardRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_card"), params).SetIdempotency(true)
}
func CreateBankAccount(params *paymentsource.CreateBankAccountRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/create_bank_account"), params).SetIdempotency(true)
}
func UpdateCard(id string, params *paymentsource.UpdateCardRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/update_card", url.PathEscape(id)), params).SetIdempotency(true)
}
func UpdateBankAccount(id string, params *paymentsource.UpdateBankAccountRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/update_bank_account", url.PathEscape(id)), params).SetIdempotency(true)
}
func VerifyBankAccount(id string, params *paymentsource.VerifyBankAccountRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/verify_bank_account", url.PathEscape(id)), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/payment_sources/%v", url.PathEscape(id)), nil)
}
func List(params *paymentsource.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/payment_sources"), params)
}
func SwitchGatewayAccount(id string, params *paymentsource.SwitchGatewayAccountRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/switch_gateway_account", url.PathEscape(id)), params).SetIdempotency(true)
}
func ExportPaymentSource(id string, params *paymentsource.ExportPaymentSourceRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/export_payment_source", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func DeleteLocal(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_sources/%v/delete_local", url.PathEscape(id)), nil).SetIdempotency(true)
}

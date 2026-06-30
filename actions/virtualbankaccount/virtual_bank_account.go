package virtualbankaccount

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"
	"net/url"
)

func CreateUsingPermanentToken(params *virtualbankaccount.CreateUsingPermanentTokenRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/create_using_permanent_token"), params).WithTelemetryResource("virtualBankAccount").WithTelemetryOperation("createUsingPermanentToken").SetIdempotency(true)
}
func Create(params *virtualbankaccount.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts"), params).WithTelemetryResource("virtualBankAccount").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/virtual_bank_accounts/%v", url.PathEscape(id)), nil).WithTelemetryResource("virtualBankAccount").WithTelemetryOperation("retrieve")
}
func List(params *virtualbankaccount.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/virtual_bank_accounts"), params).WithTelemetryResource("virtualBankAccount").WithTelemetryOperation("list")
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("virtualBankAccount").WithTelemetryOperation("delete").SetIdempotency(true)
}
func DeleteLocal(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete_local", url.PathEscape(id)), nil).WithTelemetryResource("virtualBankAccount").WithTelemetryOperation("deleteLocal").SetIdempotency(true)
}

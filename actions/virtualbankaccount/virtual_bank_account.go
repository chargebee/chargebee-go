package virtualbankaccount

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"
	"net/url"
)

func CreateUsingPermanentToken(params *virtualbankaccount.CreateUsingPermanentTokenRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/create_using_permanent_token"), params).SetIdempotency(true)
}
func Create(params *virtualbankaccount.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/virtual_bank_accounts/%v", url.PathEscape(id)), nil)
}
func List(params *virtualbankaccount.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/virtual_bank_accounts"), params)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func DeleteLocal(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete_local", url.PathEscape(id)), nil).SetIdempotency(true)
}

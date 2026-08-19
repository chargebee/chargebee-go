package vaultedpaymentmethod

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"net/url"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/vaulted_payment_methods/%v", url.PathEscape(id)), nil)
}

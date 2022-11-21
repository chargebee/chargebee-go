package inappsubscription

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/inappsubscription"
	"net/url"
)

func ProcessReceipt(id string, params *inappsubscription.ProcessReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", url.PathEscape(id)), params)
}
func ImportReceipt(id string, params *inappsubscription.ImportReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", url.PathEscape(id)), params)
}

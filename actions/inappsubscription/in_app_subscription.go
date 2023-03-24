package inappsubscription

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/inappsubscription"
	"net/url"
)

func ProcessReceipt(id string, params *inappsubscription.ProcessReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", url.PathEscape(id)), params)
}
func ImportReceipt(id string, params *inappsubscription.ImportReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", url.PathEscape(id)), params)
}
func ImportSubscription(id string, params *inappsubscription.ImportSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_subscription", url.PathEscape(id)), params)
}
func RetrieveStoreSubs(id string, params *inappsubscription.RetrieveStoreSubsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/retrieve", url.PathEscape(id)), params)
}

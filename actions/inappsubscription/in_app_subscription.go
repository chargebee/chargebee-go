package inappsubscription

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/inappsubscription"
)

func ProcessReceipt(id string, params *inappsubscription.ProcessReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", chargebee.IDEscape(id)), params)
}
func ImportReceipt(id string, params *inappsubscription.ImportReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", chargebee.IDEscape(id)), params)
}
func ImportSubscription(id string, params *inappsubscription.ImportSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_subscription", chargebee.IDEscape(id)), params)
}
func RetrieveStoreSubs(id string, params *inappsubscription.RetrieveStoreSubsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/retrieve", chargebee.IDEscape(id)), params)
}

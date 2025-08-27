package inappsubscription

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/inappsubscription"
	"net/url"
)

func ProcessReceipt(id string, params *inappsubscription.ProcessReceiptRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportReceipt(id string, params *inappsubscription.ImportReceiptRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportSubscription(id string, params *inappsubscription.ImportSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_subscription", url.PathEscape(id)), params).SetIdempotency(true)
}
func RetrieveStoreSubs(id string, params *inappsubscription.RetrieveStoreSubsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/retrieve", url.PathEscape(id)), params).SetIdempotency(true)
}

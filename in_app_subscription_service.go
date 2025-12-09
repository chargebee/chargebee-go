package chargebee

import (
	"fmt"
	"net/url"
)

type InAppSubscriptionService struct {
	transport *Transport
}

func (s *InAppSubscriptionService) ProcessReceipt(id string, req *ProcessReceiptRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InAppSubscriptionService) ImportReceipt(id string, req *ImportReceiptRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InAppSubscriptionService) ImportSubscription(id string, req *ImportSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_subscription", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InAppSubscriptionService) RetrieveStoreSubs(id string, req *RetrieveStoreSubsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/retrieve", url.PathEscape(id)), req).SetIdempotency(true)
}

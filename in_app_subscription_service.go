package chargebee

import (
	"fmt"
	"net/url"
)

type InAppSubscriptionService struct {
	config *ClientConfig
}

func (s *InAppSubscriptionService) ProcessReceipt(id string, req *InAppSubscriptionProcessReceiptRequest) (*InAppSubscriptionProcessReceiptResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", url.PathEscape(id))
	req.isIdempotent = true
	return send[*InAppSubscriptionProcessReceiptResponse](req, s.config)
}

func (s *InAppSubscriptionService) ImportReceipt(id string, req *InAppSubscriptionImportReceiptRequest) (*InAppSubscriptionImportReceiptResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", url.PathEscape(id))
	req.isIdempotent = true
	return send[*InAppSubscriptionImportReceiptResponse](req, s.config)
}

func (s *InAppSubscriptionService) ImportSubscription(id string, req *InAppSubscriptionImportSubscriptionRequest) (*InAppSubscriptionImportSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/in_app_subscriptions/%v/import_subscription", url.PathEscape(id))
	req.isIdempotent = true
	return send[*InAppSubscriptionImportSubscriptionResponse](req, s.config)
}

func (s *InAppSubscriptionService) RetrieveStoreSubs(id string, req *InAppSubscriptionRetrieveStoreSubsRequest) (*InAppSubscriptionRetrieveStoreSubsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/in_app_subscriptions/%v/retrieve", url.PathEscape(id))
	req.isIdempotent = true
	return send[*InAppSubscriptionRetrieveStoreSubsResponse](req, s.config)
}

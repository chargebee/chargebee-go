package chargebee

import (
	"fmt"
	"net/url"
)

type EstimateService struct {
	config *ClientConfig
}

func (s *EstimateService) CreateSubscription(req *EstimateCreateSubscriptionRequest) (*EstimateCreateSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/create_subscription")
	return send[*EstimateCreateSubscriptionResponse](req, s.config)
}

func (s *EstimateService) CreateSubItemEstimate(req *EstimateCreateSubItemEstimateRequest) (*EstimateCreateSubItemEstimateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/create_subscription_for_items")
	return send[*EstimateCreateSubItemEstimateResponse](req, s.config)
}

func (s *EstimateService) CreateSubForCustomerEstimate(id string, req *EstimateCreateSubForCustomerEstimateRequest) (*EstimateCreateSubForCustomerEstimateResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/customers/%v/create_subscription_estimate", url.PathEscape(id))
	return send[*EstimateCreateSubForCustomerEstimateResponse](req, s.config)
}

func (s *EstimateService) CreateSubItemForCustomerEstimate(id string, req *EstimateCreateSubItemForCustomerEstimateRequest) (*EstimateCreateSubItemForCustomerEstimateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/create_subscription_for_items_estimate", url.PathEscape(id))
	return send[*EstimateCreateSubItemForCustomerEstimateResponse](req, s.config)
}

func (s *EstimateService) UpdateSubscription(req *EstimateUpdateSubscriptionRequest) (*EstimateUpdateSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/update_subscription")
	return send[*EstimateUpdateSubscriptionResponse](req, s.config)
}

func (s *EstimateService) UpdateSubscriptionForItems(req *EstimateUpdateSubscriptionForItemsRequest) (*EstimateUpdateSubscriptionForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/update_subscription_for_items")
	return send[*EstimateUpdateSubscriptionForItemsResponse](req, s.config)
}

func (s *EstimateService) RenewalEstimate(id string, req *EstimateRenewalEstimateRequest) (*EstimateRenewalEstimateResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions/%v/renewal_estimate", url.PathEscape(id))
	return send[*EstimateRenewalEstimateResponse](req, s.config)
}

func (s *EstimateService) AdvanceInvoiceEstimate(id string, req *EstimateAdvanceInvoiceEstimateRequest) (*EstimateAdvanceInvoiceEstimateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/advance_invoice_estimate", url.PathEscape(id))
	return send[*EstimateAdvanceInvoiceEstimateResponse](req, s.config)
}

func (s *EstimateService) RegenerateInvoiceEstimate(id string, req *EstimateRegenerateInvoiceEstimateRequest) (*EstimateRegenerateInvoiceEstimateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/regenerate_invoice_estimate", url.PathEscape(id))
	return send[*EstimateRegenerateInvoiceEstimateResponse](req, s.config)
}

func (s *EstimateService) UpcomingInvoicesEstimate(id string) (*EstimateUpcomingInvoicesEstimateResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/customers/%v/upcoming_invoices_estimate", url.PathEscape(id))
	return send[*EstimateUpcomingInvoicesEstimateResponse](req, s.config)
}

func (s *EstimateService) ChangeTermEnd(id string, req *EstimateChangeTermEndRequest) (*EstimateChangeTermEndResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/change_term_end_estimate", url.PathEscape(id))
	return send[*EstimateChangeTermEndResponse](req, s.config)
}

func (s *EstimateService) CancelSubscription(id string, req *EstimateCancelSubscriptionRequest) (*EstimateCancelSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/cancel_subscription_estimate", url.PathEscape(id))
	return send[*EstimateCancelSubscriptionResponse](req, s.config)
}

func (s *EstimateService) CancelSubscriptionForItems(id string, req *EstimateCancelSubscriptionForItemsRequest) (*EstimateCancelSubscriptionForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/cancel_subscription_for_items_estimate", url.PathEscape(id))
	return send[*EstimateCancelSubscriptionForItemsResponse](req, s.config)
}

func (s *EstimateService) PauseSubscription(id string, req *EstimatePauseSubscriptionRequest) (*EstimatePauseSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/pause_subscription_estimate", url.PathEscape(id))
	return send[*EstimatePauseSubscriptionResponse](req, s.config)
}

func (s *EstimateService) ResumeSubscription(id string, req *EstimateResumeSubscriptionRequest) (*EstimateResumeSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/resume_subscription_estimate", url.PathEscape(id))
	return send[*EstimateResumeSubscriptionResponse](req, s.config)
}

func (s *EstimateService) GiftSubscription(req *EstimateGiftSubscriptionRequest) (*EstimateGiftSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/gift_subscription")
	return send[*EstimateGiftSubscriptionResponse](req, s.config)
}

func (s *EstimateService) GiftSubscriptionForItems(req *EstimateGiftSubscriptionForItemsRequest) (*EstimateGiftSubscriptionForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/gift_subscription_for_items")
	return send[*EstimateGiftSubscriptionForItemsResponse](req, s.config)
}

func (s *EstimateService) CreateInvoice(req *EstimateCreateInvoiceRequest) (*EstimateCreateInvoiceResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/create_invoice")
	return send[*EstimateCreateInvoiceResponse](req, s.config)
}

func (s *EstimateService) CreateInvoiceForItems(req *EstimateCreateInvoiceForItemsRequest) (*EstimateCreateInvoiceForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/create_invoice_for_items")
	return send[*EstimateCreateInvoiceForItemsResponse](req, s.config)
}

func (s *EstimateService) PaymentSchedules(req *EstimatePaymentSchedulesRequest) (*EstimatePaymentSchedulesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/estimates/payment_schedules")
	req.isIdempotent = true
	return send[*EstimatePaymentSchedulesResponse](req, s.config)
}

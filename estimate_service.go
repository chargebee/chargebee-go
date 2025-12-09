package chargebee

import (
	"fmt"
	"net/url"
)

type EstimateService struct {
	transport *Transport
}

func (s *EstimateService) CreateSubscription(req *CreateSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/create_subscription"), req)
}

func (s *EstimateService) CreateSubItemEstimate(req *CreateSubItemEstimateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/create_subscription_for_items"), req)
}

func (s *EstimateService) CreateSubForCustomerEstimate(id string, req *CreateSubForCustomerEstimateRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/customers/%v/create_subscription_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) CreateSubItemForCustomerEstimate(id string, req *CreateSubItemForCustomerEstimateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_for_items_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) UpdateSubscription(req *UpdateSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/update_subscription"), req)
}

func (s *EstimateService) UpdateSubscriptionForItems(req *UpdateSubscriptionForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/update_subscription_for_items"), req)
}

func (s *EstimateService) RenewalEstimate(id string, req *RenewalEstimateRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/subscriptions/%v/renewal_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) AdvanceInvoiceEstimate(id string, req *AdvanceInvoiceEstimateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/advance_invoice_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) RegenerateInvoiceEstimate(id string, req *RegenerateInvoiceEstimateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/regenerate_invoice_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) UpcomingInvoicesEstimate(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/customers/%v/upcoming_invoices_estimate", url.PathEscape(id)), nil)
}

func (s *EstimateService) ChangeTermEnd(id string, req *ChangeTermEndRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) CancelSubscription(id string, req *CancelSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) CancelSubscriptionForItems(id string, req *CancelSubscriptionForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_for_items_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) PauseSubscription(id string, req *PauseSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/pause_subscription_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) ResumeSubscription(id string, req *ResumeSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/resume_subscription_estimate", url.PathEscape(id)), req)
}

func (s *EstimateService) GiftSubscription(req *GiftSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/gift_subscription"), req)
}

func (s *EstimateService) GiftSubscriptionForItems(req *GiftSubscriptionForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/gift_subscription_for_items"), req)
}

func (s *EstimateService) CreateInvoice(req *CreateInvoiceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/create_invoice"), req)
}

func (s *EstimateService) CreateInvoiceForItems(req *CreateInvoiceForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/create_invoice_for_items"), req)
}

func (s *EstimateService) PaymentSchedules(req *PaymentSchedulesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/estimates/payment_schedules"), req).SetIdempotency(true)
}

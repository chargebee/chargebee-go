package chargebee

import (
	"fmt"
	"net/url"
)

type SubscriptionService struct {
	transport *Transport
}

func (s *SubscriptionService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions"), req).SetIdempotency(true)
}

func (s *SubscriptionService) CreateForCustomer(id string, req *CreateForCustomerRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/subscriptions", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) CreateWithItems(id string, req *CreateWithItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/subscription_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions"), req)
}

func (s *SubscriptionService) SubscriptionsForCustomer(id string, req *SubscriptionsForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/subscriptions", url.PathEscape(id)), req)
}

func (s *SubscriptionService) ContractTermsForSubscription(id string, req *ContractTermsForSubscriptionRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions/%v/contract_terms", url.PathEscape(id)), req)
}

func (s *SubscriptionService) ListDiscounts(id string, req *ListDiscountsRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions/%v/discounts", url.PathEscape(id)), req)
}

func (s *SubscriptionService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/subscriptions/%v", url.PathEscape(id)), nil)
}

func (s *SubscriptionService) RetrieveWithScheduledChanges(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_with_scheduled_changes", url.PathEscape(id)), nil)
}

func (s *SubscriptionService) RemoveScheduledChanges(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_changes", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *SubscriptionService) RemoveScheduledCancellation(id string, req *RemoveScheduledCancellationRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_cancellation", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) RemoveCoupons(id string, req *RemoveCouponsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_coupons", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) UpdateForItems(id string, req *UpdateForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/update_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ChangeTermEnd(id string, req *ChangeTermEndRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) Reactivate(id string, req *ReactivateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/reactivate", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) AddChargeAtTermEnd(id string, req *AddChargeAtTermEndRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/add_charge_at_term_end", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ChargeAddonAtTermEnd(id string, req *ChargeAddonAtTermEndRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_addon_at_term_end", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ChargeFutureRenewals(id string, req *ChargeFutureRenewalsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_future_renewals", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) EditAdvanceInvoiceSchedule(id string, req *EditAdvanceInvoiceScheduleRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/edit_advance_invoice_schedule", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) RetrieveAdvanceInvoiceSchedule(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_advance_invoice_schedule", url.PathEscape(id)), nil)
}

func (s *SubscriptionService) RemoveAdvanceInvoiceSchedule(id string, req *RemoveAdvanceInvoiceScheduleRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_advance_invoice_schedule", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) RegenerateInvoice(id string, req *RegenerateInvoiceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/regenerate_invoice", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ImportSubscription(req *ImportSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/import_subscription"), req).SetIdempotency(true)
}

func (s *SubscriptionService) ImportForCustomer(id string, req *ImportForCustomerRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/import_subscription", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ImportContractTerm(id string, req *ImportContractTermRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/import_contract_term", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ImportUnbilledCharges(id string, req *ImportUnbilledChargesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/import_unbilled_charges", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) ImportForItems(id string, req *ImportForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/import_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) OverrideBillingProfile(id string, req *OverrideBillingProfileRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/override_billing_profile", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *SubscriptionService) Pause(id string, req *PauseRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/pause", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) Cancel(id string, req *CancelRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) CancelForItems(id string, req *CancelForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) Resume(id string, req *ResumeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/resume", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *SubscriptionService) RemoveScheduledPause(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_pause", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *SubscriptionService) RemoveScheduledResumption(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_resumption", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *SubscriptionService) Move(id string, req *MoveRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/move", url.PathEscape(id)), req).SetIdempotency(true)
}

package chargebee

import (
	"fmt"
	"net/url"
)

type SubscriptionService struct {
	config *ClientConfig
}

func (s *SubscriptionService) Create(req *SubscriptionCreateRequest) (*SubscriptionCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions")
	req.isIdempotent = true
	return send[*SubscriptionCreateResponse](req, s.config)
}

func (s *SubscriptionService) CreateForCustomer(id string, req *SubscriptionCreateForCustomerRequest) (*SubscriptionCreateForCustomerResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/subscriptions", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionCreateForCustomerResponse](req, s.config)
}

func (s *SubscriptionService) CreateWithItems(id string, req *SubscriptionCreateWithItemsRequest) (*SubscriptionCreateWithItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/subscription_for_items", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionCreateWithItemsResponse](req, s.config)
}

func (s *SubscriptionService) List(req *SubscriptionListRequest) (*SubscriptionListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions")
	req.isListRequest = true
	return send[*SubscriptionListResponse](req, s.config)
}

func (s *SubscriptionService) SubscriptionsForCustomer(id string, req *SubscriptionSubscriptionsForCustomerRequest) (*SubscriptionSubscriptionsForCustomerResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/customers/%v/subscriptions", url.PathEscape(id))
	req.isListRequest = true
	return send[*SubscriptionSubscriptionsForCustomerResponse](req, s.config)
}

func (s *SubscriptionService) ContractTermsForSubscription(id string, req *SubscriptionContractTermsForSubscriptionRequest) (*SubscriptionContractTermsForSubscriptionResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions/%v/contract_terms", url.PathEscape(id))
	req.isListRequest = true
	return send[*SubscriptionContractTermsForSubscriptionResponse](req, s.config)
}

func (s *SubscriptionService) ListDiscounts(id string, req *SubscriptionListDiscountsRequest) (*SubscriptionListDiscountsResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions/%v/discounts", url.PathEscape(id))
	req.isListRequest = true
	return send[*SubscriptionListDiscountsResponse](req, s.config)
}

func (s *SubscriptionService) Retrieve(id string) (*SubscriptionRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions/%v", url.PathEscape(id))
	return send[*SubscriptionRetrieveResponse](req, s.config)
}

func (s *SubscriptionService) RetrieveWithScheduledChanges(id string) (*SubscriptionRetrieveWithScheduledChangesResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions/%v/retrieve_with_scheduled_changes", url.PathEscape(id))
	return send[*SubscriptionRetrieveWithScheduledChangesResponse](req, s.config)
}

func (s *SubscriptionService) RemoveScheduledChanges(id string) (*SubscriptionRemoveScheduledChangesResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/remove_scheduled_changes", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRemoveScheduledChangesResponse](req, s.config)
}

func (s *SubscriptionService) RemoveScheduledCancellation(id string, req *SubscriptionRemoveScheduledCancellationRequest) (*SubscriptionRemoveScheduledCancellationResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/remove_scheduled_cancellation", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRemoveScheduledCancellationResponse](req, s.config)
}

func (s *SubscriptionService) RemoveCoupons(id string, req *SubscriptionRemoveCouponsRequest) (*SubscriptionRemoveCouponsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/remove_coupons", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRemoveCouponsResponse](req, s.config)
}

func (s *SubscriptionService) Update(id string, req *SubscriptionUpdateRequest) (*SubscriptionUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionUpdateResponse](req, s.config)
}

func (s *SubscriptionService) UpdateForItems(id string, req *SubscriptionUpdateForItemsRequest) (*SubscriptionUpdateForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/update_for_items", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionUpdateForItemsResponse](req, s.config)
}

func (s *SubscriptionService) ChangeTermEnd(id string, req *SubscriptionChangeTermEndRequest) (*SubscriptionChangeTermEndResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/change_term_end", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionChangeTermEndResponse](req, s.config)
}

func (s *SubscriptionService) Reactivate(id string, req *SubscriptionReactivateRequest) (*SubscriptionReactivateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/reactivate", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionReactivateResponse](req, s.config)
}

func (s *SubscriptionService) AddChargeAtTermEnd(id string, req *SubscriptionAddChargeAtTermEndRequest) (*SubscriptionAddChargeAtTermEndResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/add_charge_at_term_end", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionAddChargeAtTermEndResponse](req, s.config)
}

func (s *SubscriptionService) ChargeAddonAtTermEnd(id string, req *SubscriptionChargeAddonAtTermEndRequest) (*SubscriptionChargeAddonAtTermEndResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/charge_addon_at_term_end", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionChargeAddonAtTermEndResponse](req, s.config)
}

func (s *SubscriptionService) ChargeFutureRenewals(id string, req *SubscriptionChargeFutureRenewalsRequest) (*SubscriptionChargeFutureRenewalsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/charge_future_renewals", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionChargeFutureRenewalsResponse](req, s.config)
}

func (s *SubscriptionService) EditAdvanceInvoiceSchedule(id string, req *SubscriptionEditAdvanceInvoiceScheduleRequest) (*SubscriptionEditAdvanceInvoiceScheduleResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/edit_advance_invoice_schedule", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionEditAdvanceInvoiceScheduleResponse](req, s.config)
}

func (s *SubscriptionService) RetrieveAdvanceInvoiceSchedule(id string) (*SubscriptionRetrieveAdvanceInvoiceScheduleResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/subscriptions/%v/retrieve_advance_invoice_schedule", url.PathEscape(id))
	return send[*SubscriptionRetrieveAdvanceInvoiceScheduleResponse](req, s.config)
}

func (s *SubscriptionService) RemoveAdvanceInvoiceSchedule(id string, req *SubscriptionRemoveAdvanceInvoiceScheduleRequest) (*SubscriptionRemoveAdvanceInvoiceScheduleResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/remove_advance_invoice_schedule", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRemoveAdvanceInvoiceScheduleResponse](req, s.config)
}

func (s *SubscriptionService) RegenerateInvoice(id string, req *SubscriptionRegenerateInvoiceRequest) (*SubscriptionRegenerateInvoiceResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/regenerate_invoice", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRegenerateInvoiceResponse](req, s.config)
}

func (s *SubscriptionService) ImportSubscription(req *SubscriptionImportSubscriptionRequest) (*SubscriptionImportSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/import_subscription")
	req.isIdempotent = true
	return send[*SubscriptionImportSubscriptionResponse](req, s.config)
}

func (s *SubscriptionService) ImportForCustomer(id string, req *SubscriptionImportForCustomerRequest) (*SubscriptionImportForCustomerResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/import_subscription", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionImportForCustomerResponse](req, s.config)
}

func (s *SubscriptionService) ImportContractTerm(id string, req *SubscriptionImportContractTermRequest) (*SubscriptionImportContractTermResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/import_contract_term", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionImportContractTermResponse](req, s.config)
}

func (s *SubscriptionService) ImportUnbilledCharges(id string, req *SubscriptionImportUnbilledChargesRequest) (*SubscriptionImportUnbilledChargesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/import_unbilled_charges", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionImportUnbilledChargesResponse](req, s.config)
}

func (s *SubscriptionService) ImportForItems(id string, req *SubscriptionImportForItemsRequest) (*SubscriptionImportForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/import_for_items", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionImportForItemsResponse](req, s.config)
}

func (s *SubscriptionService) OverrideBillingProfile(id string, req *SubscriptionOverrideBillingProfileRequest) (*SubscriptionOverrideBillingProfileResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/override_billing_profile", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionOverrideBillingProfileResponse](req, s.config)
}

func (s *SubscriptionService) Delete(id string) (*SubscriptionDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionDeleteResponse](req, s.config)
}

func (s *SubscriptionService) Pause(id string, req *SubscriptionPauseRequest) (*SubscriptionPauseResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/pause", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionPauseResponse](req, s.config)
}

func (s *SubscriptionService) Cancel(id string, req *SubscriptionCancelRequest) (*SubscriptionCancelResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/cancel", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionCancelResponse](req, s.config)
}

func (s *SubscriptionService) CancelForItems(id string, req *SubscriptionCancelForItemsRequest) (*SubscriptionCancelForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/cancel_for_items", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionCancelForItemsResponse](req, s.config)
}

func (s *SubscriptionService) Resume(id string, req *SubscriptionResumeRequest) (*SubscriptionResumeResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/resume", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionResumeResponse](req, s.config)
}

func (s *SubscriptionService) RemoveScheduledPause(id string) (*SubscriptionRemoveScheduledPauseResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/remove_scheduled_pause", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRemoveScheduledPauseResponse](req, s.config)
}

func (s *SubscriptionService) RemoveScheduledResumption(id string) (*SubscriptionRemoveScheduledResumptionResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/remove_scheduled_resumption", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionRemoveScheduledResumptionResponse](req, s.config)
}

func (s *SubscriptionService) Move(id string, req *SubscriptionMoveRequest) (*SubscriptionMoveResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/subscriptions/%v/move", url.PathEscape(id))
	req.isIdempotent = true
	return send[*SubscriptionMoveResponse](req, s.config)
}

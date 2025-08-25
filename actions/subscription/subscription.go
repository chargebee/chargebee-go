package subscription

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/subscription"
	"net/url"
)

func Create(params *subscription.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions"), params).SetIdempotency(true)
}
func CreateForCustomer(id string, params *subscription.CreateForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/subscriptions", url.PathEscape(id)), params).SetIdempotency(true)
}
func CreateWithItems(id string, params *subscription.CreateWithItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/subscription_for_items", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *subscription.ListRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions"), params)
}
func SubscriptionsForCustomer(id string, params *subscription.SubscriptionsForCustomerRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/subscriptions", url.PathEscape(id)), params)
}
func ContractTermsForSubscription(id string, params *subscription.ContractTermsForSubscriptionRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/contract_terms", url.PathEscape(id)), params)
}
func ListDiscounts(id string, params *subscription.ListDiscountsRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/discounts", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v", url.PathEscape(id)), nil)
}
func RetrieveWithScheduledChanges(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_with_scheduled_changes", url.PathEscape(id)), nil)
}
func RemoveScheduledChanges(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_changes", url.PathEscape(id)), nil).SetIdempotency(true)
}
func RemoveScheduledCancellation(id string, params *subscription.RemoveScheduledCancellationRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_cancellation", url.PathEscape(id)), params).SetIdempotency(true)
}
func RemoveCoupons(id string, params *subscription.RemoveCouponsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_coupons", url.PathEscape(id)), params).SetIdempotency(true)
}
func Update(id string, params *subscription.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func UpdateForItems(id string, params *subscription.UpdateForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/update_for_items", url.PathEscape(id)), params).SetIdempotency(true)
}
func ChangeTermEnd(id string, params *subscription.ChangeTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end", url.PathEscape(id)), params).SetIdempotency(true)
}
func Reactivate(id string, params *subscription.ReactivateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/reactivate", url.PathEscape(id)), params).SetIdempotency(true)
}
func AddChargeAtTermEnd(id string, params *subscription.AddChargeAtTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/add_charge_at_term_end", url.PathEscape(id)), params).SetIdempotency(true)
}
func ChargeAddonAtTermEnd(id string, params *subscription.ChargeAddonAtTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_addon_at_term_end", url.PathEscape(id)), params).SetIdempotency(true)
}
func ChargeFutureRenewals(id string, params *subscription.ChargeFutureRenewalsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_future_renewals", url.PathEscape(id)), params).SetIdempotency(true)
}
func EditAdvanceInvoiceSchedule(id string, params *subscription.EditAdvanceInvoiceScheduleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/edit_advance_invoice_schedule", url.PathEscape(id)), params).SetIdempotency(true)
}
func RetrieveAdvanceInvoiceSchedule(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_advance_invoice_schedule", url.PathEscape(id)), nil)
}
func RemoveAdvanceInvoiceSchedule(id string, params *subscription.RemoveAdvanceInvoiceScheduleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_advance_invoice_schedule", url.PathEscape(id)), params).SetIdempotency(true)
}
func RegenerateInvoice(id string, params *subscription.RegenerateInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/regenerate_invoice", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportSubscription(params *subscription.ImportSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/import_subscription"), params).SetIdempotency(true)
}
func ImportForCustomer(id string, params *subscription.ImportForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/import_subscription", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportContractTerm(id string, params *subscription.ImportContractTermRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/import_contract_term", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportUnbilledCharges(id string, params *subscription.ImportUnbilledChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/import_unbilled_charges", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportForItems(id string, params *subscription.ImportForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/import_for_items", url.PathEscape(id)), params).SetIdempotency(true)
}
func OverrideBillingProfile(id string, params *subscription.OverrideBillingProfileRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/override_billing_profile", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Pause(id string, params *subscription.PauseRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/pause", url.PathEscape(id)), params).SetIdempotency(true)
}
func Cancel(id string, params *subscription.CancelRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel", url.PathEscape(id)), params).SetIdempotency(true)
}
func CancelForItems(id string, params *subscription.CancelForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_for_items", url.PathEscape(id)), params).SetIdempotency(true)
}
func Resume(id string, params *subscription.ResumeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/resume", url.PathEscape(id)), params).SetIdempotency(true)
}
func RemoveScheduledPause(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_pause", url.PathEscape(id)), nil).SetIdempotency(true)
}
func RemoveScheduledResumption(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_resumption", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Move(id string, params *subscription.MoveRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/move", url.PathEscape(id)), params).SetIdempotency(true)
}

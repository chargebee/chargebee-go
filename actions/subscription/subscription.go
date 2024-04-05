package subscription

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/subscription"
)

func Create(params *subscription.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions"), params)
}
func CreateForCustomer(id string, params *subscription.CreateForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/subscriptions", chargebee.IDEscape(id)), params)
}
func CreateWithItems(id string, params *subscription.CreateWithItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/subscription_for_items", chargebee.IDEscape(id)), params)
}
func List(params *subscription.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions"), params)
}
func SubscriptionsForCustomer(id string, params *subscription.SubscriptionsForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/subscriptions", chargebee.IDEscape(id)), params)
}
func ContractTermsForSubscription(id string, params *subscription.ContractTermsForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/contract_terms", chargebee.IDEscape(id)), params)
}
func ListDiscounts(id string, params *subscription.ListDiscountsRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/discounts", chargebee.IDEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v", chargebee.IDEscape(id)), nil)
}
func RetrieveWithScheduledChanges(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_with_scheduled_changes", chargebee.IDEscape(id)), nil)
}
func RemoveScheduledChanges(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_changes", chargebee.IDEscape(id)), nil)
}
func RemoveScheduledCancellation(id string, params *subscription.RemoveScheduledCancellationRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_cancellation", chargebee.IDEscape(id)), params)
}
func RemoveCoupons(id string, params *subscription.RemoveCouponsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_coupons", chargebee.IDEscape(id)), params)
}
func Update(id string, params *subscription.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v", chargebee.IDEscape(id)), params)
}
func UpdateForItems(id string, params *subscription.UpdateForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/update_for_items", chargebee.IDEscape(id)), params)
}
func ChangeTermEnd(id string, params *subscription.ChangeTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end", chargebee.IDEscape(id)), params)
}
func Reactivate(id string, params *subscription.ReactivateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/reactivate", chargebee.IDEscape(id)), params)
}
func AddChargeAtTermEnd(id string, params *subscription.AddChargeAtTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/add_charge_at_term_end", chargebee.IDEscape(id)), params)
}
func ChargeAddonAtTermEnd(id string, params *subscription.ChargeAddonAtTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_addon_at_term_end", chargebee.IDEscape(id)), params)
}
func ChargeFutureRenewals(id string, params *subscription.ChargeFutureRenewalsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_future_renewals", chargebee.IDEscape(id)), params)
}
func EditAdvanceInvoiceSchedule(id string, params *subscription.EditAdvanceInvoiceScheduleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/edit_advance_invoice_schedule", chargebee.IDEscape(id)), params)
}
func RetrieveAdvanceInvoiceSchedule(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_advance_invoice_schedule", chargebee.IDEscape(id)), nil)
}
func RemoveAdvanceInvoiceSchedule(id string, params *subscription.RemoveAdvanceInvoiceScheduleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_advance_invoice_schedule", chargebee.IDEscape(id)), params)
}
func RegenerateInvoice(id string, params *subscription.RegenerateInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/regenerate_invoice", chargebee.IDEscape(id)), params)
}
func ImportSubscription(params *subscription.ImportSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/import_subscription"), params)
}
func ImportForCustomer(id string, params *subscription.ImportForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/import_subscription", chargebee.IDEscape(id)), params)
}
func ImportContractTerm(id string, params *subscription.ImportContractTermRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/import_contract_term", chargebee.IDEscape(id)), params)
}
func ImportUnbilledCharges(id string, params *subscription.ImportUnbilledChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/import_unbilled_charges", chargebee.IDEscape(id)), params)
}
func ImportForItems(id string, params *subscription.ImportForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/import_for_items", chargebee.IDEscape(id)), params)
}
func OverrideBillingProfile(id string, params *subscription.OverrideBillingProfileRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/override_billing_profile", chargebee.IDEscape(id)), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/delete", chargebee.IDEscape(id)), nil)
}
func Pause(id string, params *subscription.PauseRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/pause", chargebee.IDEscape(id)), params)
}
func Cancel(id string, params *subscription.CancelRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel", chargebee.IDEscape(id)), params)
}
func CancelForItems(id string, params *subscription.CancelForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_for_items", chargebee.IDEscape(id)), params)
}
func Resume(id string, params *subscription.ResumeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/resume", chargebee.IDEscape(id)), params)
}
func RemoveScheduledPause(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_pause", chargebee.IDEscape(id)), nil)
}
func RemoveScheduledResumption(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_resumption", chargebee.IDEscape(id)), nil)
}
func Move(id string, params *subscription.MoveRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/move", url.PathEscape(id)), params)
}

package estimate

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/estimate"
	"net/url"
)

func CreateSubscription(params *estimate.CreateSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_subscription"), params)
}
func CreateSubItemEstimate(params *estimate.CreateSubItemEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_subscription_for_items"), params)
}
func CreateSubForCustomerEstimate(id string, params *estimate.CreateSubForCustomerEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/create_subscription_estimate", url.PathEscape(id)), params)
}
func CreateSubItemForCustomerEstimate(id string, params *estimate.CreateSubItemForCustomerEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_for_items_estimate", url.PathEscape(id)), params)
}
func UpdateSubscription(params *estimate.UpdateSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/update_subscription"), params)
}
func UpdateSubscriptionForItems(params *estimate.UpdateSubscriptionForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/update_subscription_for_items"), params)
}
func RenewalEstimate(id string, params *estimate.RenewalEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/renewal_estimate", url.PathEscape(id)), params)
}
func AdvanceInvoiceEstimate(id string, params *estimate.AdvanceInvoiceEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/advance_invoice_estimate", url.PathEscape(id)), params)
}
func RegenerateInvoiceEstimate(id string, params *estimate.RegenerateInvoiceEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/regenerate_invoice_estimate", url.PathEscape(id)), params)
}
func UpcomingInvoicesEstimate(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/upcoming_invoices_estimate", url.PathEscape(id)), nil)
}
func ChangeTermEnd(id string, params *estimate.ChangeTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end_estimate", url.PathEscape(id)), params)
}
func CancelSubscription(id string, params *estimate.CancelSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_estimate", url.PathEscape(id)), params)
}
func CancelSubscriptionForItems(id string, params *estimate.CancelSubscriptionForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_for_items_estimate", url.PathEscape(id)), params)
}
func PauseSubscription(id string, params *estimate.PauseSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/pause_subscription_estimate", url.PathEscape(id)), params)
}
func ResumeSubscription(id string, params *estimate.ResumeSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/resume_subscription_estimate", url.PathEscape(id)), params)
}
func GiftSubscription(params *estimate.GiftSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/gift_subscription"), params)
}
func GiftSubscriptionForItems(params *estimate.GiftSubscriptionForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/gift_subscription_for_items"), params)
}
func CreateInvoice(params *estimate.CreateInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_invoice"), params)
}
func CreateInvoiceForItems(params *estimate.CreateInvoiceForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_invoice_for_items"), params)
}
func PaymentSchedules(params *estimate.PaymentSchedulesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/payment_schedules"), params)
}

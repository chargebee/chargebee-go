package estimate

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/estimate"
	"net/url"
)

func CreateSubscription(params *estimate.CreateSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_subscription"), params).WithTelemetryResource("estimate").WithTelemetryOperation("createSubscription")
}
func CreateSubItemEstimate(params *estimate.CreateSubItemEstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_subscription_for_items"), params).WithTelemetryResource("estimate").WithTelemetryOperation("createSubItemEstimate")
}
func CreateSubForCustomerEstimate(id string, params *estimate.CreateSubForCustomerEstimateRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/create_subscription_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("createSubForCustomerEstimate")
}
func CreateSubItemForCustomerEstimate(id string, params *estimate.CreateSubItemForCustomerEstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_for_items_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("createSubItemForCustomerEstimate")
}
func UpdateSubscription(params *estimate.UpdateSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/update_subscription"), params).WithTelemetryResource("estimate").WithTelemetryOperation("updateSubscription")
}
func UpdateSubscriptionForItems(params *estimate.UpdateSubscriptionForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/update_subscription_for_items"), params).WithTelemetryResource("estimate").WithTelemetryOperation("updateSubscriptionForItems")
}
func RenewalEstimate(id string, params *estimate.RenewalEstimateRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/renewal_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("renewalEstimate")
}
func AdvanceInvoiceEstimate(id string, params *estimate.AdvanceInvoiceEstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/advance_invoice_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("advanceInvoiceEstimate")
}
func RegenerateInvoiceEstimate(id string, params *estimate.RegenerateInvoiceEstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/regenerate_invoice_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("regenerateInvoiceEstimate")
}
func UpcomingInvoicesEstimate(id string, params *estimate.UpcomingInvoicesEstimateRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/upcoming_invoices_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("upcomingInvoicesEstimate")
}
func ChangeTermEnd(id string, params *estimate.ChangeTermEndRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("changeTermEnd")
}
func CancelSubscription(id string, params *estimate.CancelSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("cancelSubscription")
}
func CancelSubscriptionForItems(id string, params *estimate.CancelSubscriptionForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_for_items_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("cancelSubscriptionForItems")
}
func PauseSubscription(id string, params *estimate.PauseSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/pause_subscription_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("pauseSubscription")
}
func ResumeSubscription(id string, params *estimate.ResumeSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/resume_subscription_estimate", url.PathEscape(id)), params).WithTelemetryResource("estimate").WithTelemetryOperation("resumeSubscription")
}
func GiftSubscription(params *estimate.GiftSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/gift_subscription"), params).WithTelemetryResource("estimate").WithTelemetryOperation("giftSubscription")
}
func GiftSubscriptionForItems(params *estimate.GiftSubscriptionForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/gift_subscription_for_items"), params).WithTelemetryResource("estimate").WithTelemetryOperation("giftSubscriptionForItems")
}
func CreateInvoice(params *estimate.CreateInvoiceRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_invoice"), params).WithTelemetryResource("estimate").WithTelemetryOperation("createInvoice")
}
func CreateInvoiceForItems(params *estimate.CreateInvoiceForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_invoice_for_items"), params).WithTelemetryResource("estimate").WithTelemetryOperation("createInvoiceForItems")
}
func PaymentSchedules(params *estimate.PaymentSchedulesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/payment_schedules"), params).WithTelemetryResource("estimate").WithTelemetryOperation("paymentSchedules").SetIdempotency(true)
}

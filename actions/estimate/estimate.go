package estimate

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/estimate"
)

func CreateSubscription(params *estimate.CreateSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_subscription"), params)
}
func CreateSubForCustomerEstimate(id string, params *estimate.CreateSubForCustomerEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/create_subscription_estimate", id), params)
}
func UpdateSubscription(params *estimate.UpdateSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/update_subscription"), params)
}
func RenewalEstimate(id string, params *estimate.RenewalEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/renewal_estimate", id), params)
}
func UpcomingInvoicesEstimate(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/upcoming_invoices_estimate", id), nil)
}
func ChangeTermEnd(id string, params *estimate.ChangeTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end_estimate", id), params)
}
func CancelSubscription(id string, params *estimate.CancelSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel_subscription_estimate", id), params)
}
func PauseSubscription(id string, params *estimate.PauseSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/pause_subscription_estimate", id), params)
}
func ResumeSubscription(id string, params *estimate.ResumeSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/resume_subscription_estimate", id), params)
}
func GiftSubscription(params *estimate.GiftSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/gift_subscription"), params)
}
func CreateInvoice(params *estimate.CreateInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/estimates/create_invoice"), params)
}

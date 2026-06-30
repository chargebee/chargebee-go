package hostedpage

import (
	"encoding/json"
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/hostedpage"
	"net/url"
)

func CheckoutNew(params *hostedpage.CheckoutNewRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutNew").SetIdempotency(true)
}
func CheckoutOneTime(params *hostedpage.CheckoutOneTimeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_one_time"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutOneTime").SetIdempotency(true)
}
func CheckoutOneTimeForItems(params *hostedpage.CheckoutOneTimeForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_one_time_for_items"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutOneTimeForItems").SetIdempotency(true)
}
func CheckoutNewForItems(params *hostedpage.CheckoutNewForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new_for_items"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutNewForItems").SetIdempotency(true)
}
func CheckoutExisting(params *hostedpage.CheckoutExistingRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutExisting").SetIdempotency(true)
}
func CheckoutExistingForItems(params *hostedpage.CheckoutExistingForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing_for_items"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutExistingForItems").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func UpdateCard(params *hostedpage.UpdateCardRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/update_card"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("updateCard").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func UpdatePaymentMethod(params *hostedpage.UpdatePaymentMethodRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/update_payment_method"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("updatePaymentMethod").SetIdempotency(true)
}
func ManagePaymentSources(params *hostedpage.ManagePaymentSourcesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/manage_payment_sources"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("managePaymentSources").SetIdempotency(true)
}
func CollectNow(params *hostedpage.CollectNowRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/collect_now"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("collectNow").SetIdempotency(true)
}
func AcceptQuote(params *hostedpage.AcceptQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/accept_quote"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("acceptQuote").SetIdempotency(true)
}
func ExtendSubscription(params *hostedpage.ExtendSubscriptionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/extend_subscription"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("extendSubscription").SetIdempotency(true)
}
func CheckoutGift(params *hostedpage.CheckoutGiftRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_gift"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutGift").SetIdempotency(true)
}
func CheckoutGiftForItems(params *hostedpage.CheckoutGiftForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_gift_for_items"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("checkoutGiftForItems").SetIdempotency(true)
}
func ClaimGift(params *hostedpage.ClaimGiftRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/claim_gift"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("claimGift").SetIdempotency(true)
}
func RetrieveAgreementPdf(params *hostedpage.RetrieveAgreementPdfRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/retrieve_agreement_pdf"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("retrieveAgreementPdf").SetIdempotency(true)
}
func Acknowledge(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/%v/acknowledge", url.PathEscape(id)), nil).WithTelemetryResource("hostedPage").WithTelemetryOperation("acknowledge").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/hosted_pages/%v", url.PathEscape(id)), nil).WithTelemetryResource("hostedPage").WithTelemetryOperation("retrieve")
}
func List(params *hostedpage.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/hosted_pages"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("list")
}
func PreCancel(params *hostedpage.PreCancelRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/pre_cancel"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("preCancel").SetIdempotency(true)
}
func Events(params *hostedpage.EventsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/events"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("events").SetIdempotency(true)
}
func ViewVoucher(params *hostedpage.ViewVoucherRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/view_voucher"), params).WithTelemetryResource("hostedPage").WithTelemetryOperation("viewVoucher").SetIdempotency(true)
}
func Content(page hostedpage.HostedPage) *chargebee.Result {
	content := &chargebee.Result{}
	err1 := json.Unmarshal(page.Content, content)
	if err1 != nil {
		panic(err1)
	}
	return content
}

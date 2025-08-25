package hostedpage

import (
	"encoding/json"
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/hostedpage"
	"net/url"
)

func CheckoutNew(params *hostedpage.CheckoutNewRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new"), params).SetIdempotency(true)
}
func CheckoutOneTime(params *hostedpage.CheckoutOneTimeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_one_time"), params).SetIdempotency(true)
}
func CheckoutOneTimeForItems(params *hostedpage.CheckoutOneTimeForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_one_time_for_items"), params).SetIdempotency(true)
}
func CheckoutNewForItems(params *hostedpage.CheckoutNewForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new_for_items"), params).SetIdempotency(true)
}
func CheckoutExisting(params *hostedpage.CheckoutExistingRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing"), params).SetIdempotency(true)
}
func CheckoutExistingForItems(params *hostedpage.CheckoutExistingForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing_for_items"), params).SetIdempotency(true)
}
func UpdateCard(params *hostedpage.UpdateCardRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/update_card"), params).SetIdempotency(true)
}
func UpdatePaymentMethod(params *hostedpage.UpdatePaymentMethodRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/update_payment_method"), params).SetIdempotency(true)
}
func ManagePaymentSources(params *hostedpage.ManagePaymentSourcesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/manage_payment_sources"), params).SetIdempotency(true)
}
func CollectNow(params *hostedpage.CollectNowRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/collect_now"), params).SetIdempotency(true)
}
func AcceptQuote(params *hostedpage.AcceptQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/accept_quote"), params).SetIdempotency(true)
}
func ExtendSubscription(params *hostedpage.ExtendSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/extend_subscription"), params).SetIdempotency(true)
}
func CheckoutGift(params *hostedpage.CheckoutGiftRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_gift"), params).SetIdempotency(true)
}
func CheckoutGiftForItems(params *hostedpage.CheckoutGiftForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_gift_for_items"), params).SetIdempotency(true)
}
func ClaimGift(params *hostedpage.ClaimGiftRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/claim_gift"), params).SetIdempotency(true)
}
func RetrieveAgreementPdf(params *hostedpage.RetrieveAgreementPdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/retrieve_agreement_pdf"), params).SetIdempotency(true)
}
func Acknowledge(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/%v/acknowledge", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/hosted_pages/%v", url.PathEscape(id)), nil)
}
func List(params *hostedpage.ListRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/hosted_pages"), params)
}
func PreCancel(params *hostedpage.PreCancelRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/pre_cancel"), params).SetIdempotency(true)
}
func Events(params *hostedpage.EventsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/events"), params).SetIdempotency(true)
}
func ViewVoucher(params *hostedpage.ViewVoucherRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/view_voucher"), params).SetIdempotency(true)
}
func Content(page hostedpage.HostedPage) *chargebee.Result {
	content := &chargebee.Result{}
	err1 := json.Unmarshal(page.Content, content)
	if err1 != nil {
		panic(err1)
	}
	return content
}

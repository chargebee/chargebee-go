package chargebee

import (
	"encoding/json"
	"fmt"

	"net/url"
)

type HostedPageService struct {
	transport *Transport
}

func (s *HostedPageService) CheckoutNew(req *CheckoutNewRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutOneTime(req *CheckoutOneTimeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_one_time"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutOneTimeForItems(req *CheckoutOneTimeForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_one_time_for_items"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutNewForItems(req *CheckoutNewForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new_for_items"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutExisting(req *CheckoutExistingRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutExistingForItems(req *CheckoutExistingForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing_for_items"), req).SetIdempotency(true)
}

func (s *HostedPageService) UpdateCard(req *UpdateCardRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/update_card"), req).SetIdempotency(true)
}

func (s *HostedPageService) UpdatePaymentMethod(req *UpdatePaymentMethodRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/update_payment_method"), req).SetIdempotency(true)
}

func (s *HostedPageService) ManagePaymentSources(req *ManagePaymentSourcesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/manage_payment_sources"), req).SetIdempotency(true)
}

func (s *HostedPageService) CollectNow(req *CollectNowRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/collect_now"), req).SetIdempotency(true)
}

func (s *HostedPageService) AcceptQuote(req *AcceptQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/accept_quote"), req).SetIdempotency(true)
}

func (s *HostedPageService) ExtendSubscription(req *ExtendSubscriptionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/extend_subscription"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutGift(req *CheckoutGiftRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_gift"), req).SetIdempotency(true)
}

func (s *HostedPageService) CheckoutGiftForItems(req *CheckoutGiftForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/checkout_gift_for_items"), req).SetIdempotency(true)
}

func (s *HostedPageService) ClaimGift(req *ClaimGiftRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/claim_gift"), req).SetIdempotency(true)
}

func (s *HostedPageService) RetrieveAgreementPdf(req *RetrieveAgreementPdfRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/retrieve_agreement_pdf"), req).SetIdempotency(true)
}

func (s *HostedPageService) Acknowledge(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/%v/acknowledge", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *HostedPageService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/hosted_pages/%v", url.PathEscape(id)), nil)
}

func (s *HostedPageService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/hosted_pages"), req)
}

func (s *HostedPageService) PreCancel(req *PreCancelRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/pre_cancel"), req).SetIdempotency(true)
}

func (s *HostedPageService) Events(req *EventsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/events"), req).SetIdempotency(true)
}

func (s *HostedPageService) ViewVoucher(req *ViewVoucherRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/hosted_pages/view_voucher"), req).SetIdempotency(true)
}

func Content(page HostedPage) *chargebee.Result {
	content := &chargebee.Result{}
	err1 := json.Unmarshal(page.Content, content)
	if err1 != nil {
		panic(err1)
	}
	return content
}

package chargebee

import (
	"fmt"
	"net/url"
)

type HostedPageService struct {
	config *ClientConfig
}

func (s *HostedPageService) CheckoutNew(req *HostedPageCheckoutNewRequest) (*HostedPageCheckoutNewResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_new")
	req.isIdempotent = true
	return send[*HostedPageCheckoutNewResponse](req, s.config)
}

func (s *HostedPageService) CheckoutOneTime(req *HostedPageCheckoutOneTimeRequest) (*HostedPageCheckoutOneTimeResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_one_time")
	req.isIdempotent = true
	return send[*HostedPageCheckoutOneTimeResponse](req, s.config)
}

func (s *HostedPageService) CheckoutOneTimeForItems(req *HostedPageCheckoutOneTimeForItemsRequest) (*HostedPageCheckoutOneTimeForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_one_time_for_items")
	req.isIdempotent = true
	return send[*HostedPageCheckoutOneTimeForItemsResponse](req, s.config)
}

func (s *HostedPageService) CheckoutNewForItems(req *HostedPageCheckoutNewForItemsRequest) (*HostedPageCheckoutNewForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_new_for_items")
	req.isIdempotent = true
	return send[*HostedPageCheckoutNewForItemsResponse](req, s.config)
}

func (s *HostedPageService) CheckoutExisting(req *HostedPageCheckoutExistingRequest) (*HostedPageCheckoutExistingResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_existing")
	req.isIdempotent = true
	return send[*HostedPageCheckoutExistingResponse](req, s.config)
}

func (s *HostedPageService) CheckoutExistingForItems(req *HostedPageCheckoutExistingForItemsRequest) (*HostedPageCheckoutExistingForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_existing_for_items")
	req.isIdempotent = true
	return send[*HostedPageCheckoutExistingForItemsResponse](req, s.config)
}

func (s *HostedPageService) UpdateCard(req *HostedPageUpdateCardRequest) (*HostedPageUpdateCardResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/update_card")
	req.isIdempotent = true
	return send[*HostedPageUpdateCardResponse](req, s.config)
}

func (s *HostedPageService) UpdatePaymentMethod(req *HostedPageUpdatePaymentMethodRequest) (*HostedPageUpdatePaymentMethodResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/update_payment_method")
	req.isIdempotent = true
	return send[*HostedPageUpdatePaymentMethodResponse](req, s.config)
}

func (s *HostedPageService) ManagePaymentSources(req *HostedPageManagePaymentSourcesRequest) (*HostedPageManagePaymentSourcesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/manage_payment_sources")
	req.isIdempotent = true
	return send[*HostedPageManagePaymentSourcesResponse](req, s.config)
}

func (s *HostedPageService) CollectNow(req *HostedPageCollectNowRequest) (*HostedPageCollectNowResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/collect_now")
	req.isIdempotent = true
	return send[*HostedPageCollectNowResponse](req, s.config)
}

func (s *HostedPageService) AcceptQuote(req *HostedPageAcceptQuoteRequest) (*HostedPageAcceptQuoteResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/accept_quote")
	req.isIdempotent = true
	return send[*HostedPageAcceptQuoteResponse](req, s.config)
}

func (s *HostedPageService) ExtendSubscription(req *HostedPageExtendSubscriptionRequest) (*HostedPageExtendSubscriptionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/extend_subscription")
	req.isIdempotent = true
	return send[*HostedPageExtendSubscriptionResponse](req, s.config)
}

func (s *HostedPageService) CheckoutGift(req *HostedPageCheckoutGiftRequest) (*HostedPageCheckoutGiftResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_gift")
	req.isIdempotent = true
	return send[*HostedPageCheckoutGiftResponse](req, s.config)
}

func (s *HostedPageService) CheckoutGiftForItems(req *HostedPageCheckoutGiftForItemsRequest) (*HostedPageCheckoutGiftForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/checkout_gift_for_items")
	req.isIdempotent = true
	return send[*HostedPageCheckoutGiftForItemsResponse](req, s.config)
}

func (s *HostedPageService) ClaimGift(req *HostedPageClaimGiftRequest) (*HostedPageClaimGiftResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/claim_gift")
	req.isIdempotent = true
	return send[*HostedPageClaimGiftResponse](req, s.config)
}

func (s *HostedPageService) RetrieveAgreementPdf(req *HostedPageRetrieveAgreementPdfRequest) (*HostedPageRetrieveAgreementPdfResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/retrieve_agreement_pdf")
	req.isIdempotent = true
	return send[*HostedPageRetrieveAgreementPdfResponse](req, s.config)
}

func (s *HostedPageService) Acknowledge(id string) (*HostedPageAcknowledgeResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/%v/acknowledge", url.PathEscape(id))
	req.isIdempotent = true
	return send[*HostedPageAcknowledgeResponse](req, s.config)
}

func (s *HostedPageService) Retrieve(id string) (*HostedPageRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/hosted_pages/%v", url.PathEscape(id))
	return send[*HostedPageRetrieveResponse](req, s.config)
}

func (s *HostedPageService) List(req *HostedPageListRequest) (*HostedPageListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/hosted_pages")
	req.isListRequest = true
	return send[*HostedPageListResponse](req, s.config)
}

func (s *HostedPageService) PreCancel(req *HostedPagePreCancelRequest) (*HostedPagePreCancelResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/pre_cancel")
	req.isIdempotent = true
	return send[*HostedPagePreCancelResponse](req, s.config)
}

func (s *HostedPageService) Events(req *HostedPageEventsRequest) (*HostedPageEventsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/events")
	req.isIdempotent = true
	return send[*HostedPageEventsResponse](req, s.config)
}

func (s *HostedPageService) ViewVoucher(req *HostedPageViewVoucherRequest) (*HostedPageViewVoucherResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/hosted_pages/view_voucher")
	req.isIdempotent = true
	return send[*HostedPageViewVoucherResponse](req, s.config)
}

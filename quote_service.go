package chargebee

import (
	"fmt"
	"net/url"
)

type QuoteService struct {
	transport *Transport
}

func (s *QuoteService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/quotes/%v", url.PathEscape(id)), nil)
}

func (s *QuoteService) CreateSubForCustomerQuote(id string, req *CreateSubForCustomerQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) EditCreateSubForCustomerQuote(id string, req *EditCreateSubForCustomerQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) UpdateSubscriptionQuote(req *UpdateSubscriptionQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote"), req).SetIdempotency(true)
}

func (s *QuoteService) EditUpdateSubscriptionQuote(id string, req *EditUpdateSubscriptionQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) CreateForOnetimeCharges(req *CreateForOnetimeChargesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/create_for_onetime_charges"), req).SetIdempotency(true)
}

func (s *QuoteService) EditOneTimeQuote(id string, req *EditOneTimeQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/edit_one_time_quote", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) CreateSubItemsForCustomerQuote(id string, req *CreateSubItemsForCustomerQuoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) EditCreateSubCustomerQuoteForItems(id string, req *EditCreateSubCustomerQuoteForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) UpdateSubscriptionQuoteForItems(req *UpdateSubscriptionQuoteForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote_for_items"), req).SetIdempotency(true)
}

func (s *QuoteService) EditUpdateSubscriptionQuoteForItems(id string, req *EditUpdateSubscriptionQuoteForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) CreateForChargeItemsAndCharges(req *CreateForChargeItemsAndChargesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/create_for_charge_items_and_charges"), req).SetIdempotency(true)
}

func (s *QuoteService) EditForChargeItemsAndCharges(id string, req *EditForChargeItemsAndChargesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/edit_for_charge_items_and_charges", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/quotes"), req)
}

func (s *QuoteService) QuoteLineGroupsForQuote(id string, req *QuoteLineGroupsForQuoteRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/quotes/%v/quote_line_groups", url.PathEscape(id)), req)
}

func (s *QuoteService) Convert(id string, req *ConvertRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/convert", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) UpdateStatus(id string, req *UpdateStatusRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/update_status", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) ExtendExpiryDate(id string, req *ExtendExpiryDateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/extend_expiry_date", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/delete", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *QuoteService) Pdf(id string, req *PdfRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/quotes/%v/pdf", url.PathEscape(id)), req).SetIdempotency(true)
}

package chargebee

import (
	"fmt"
	"net/url"
)

type OrderService struct {
	transport *Transport
}

func (s *OrderService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders"), req).SetIdempotency(true)
}

func (s *OrderService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *OrderService) ImportOrder(req *ImportOrderRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/import_order"), req).SetIdempotency(true)
}

func (s *OrderService) AssignOrderNumber(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v/assign_order_number", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *OrderService) Cancel(id string, req *CancelRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v/cancel", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *OrderService) CreateRefundableCreditNote(id string, req *CreateRefundableCreditNoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v/create_refundable_credit_note", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *OrderService) Reopen(id string, req *ReopenRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v/reopen", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *OrderService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/orders/%v", url.PathEscape(id)), nil)
}

func (s *OrderService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *OrderService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/orders"), req)
}

func (s *OrderService) OrdersForInvoice(id string, req *OrdersForInvoiceRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/invoices/%v/orders", url.PathEscape(id)), req)
}

func (s *OrderService) Resend(id string, req *ResendRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/orders/%v/resend", url.PathEscape(id)), req).SetIdempotency(true)
}

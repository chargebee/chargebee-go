package order

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/order"
	"net/url"
)

func Create(params *order.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders"), params).SetIdempotency(true)
}
func Update(id string, params *order.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func ImportOrder(params *order.ImportOrderRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/import_order"), params).SetIdempotency(true)
}
func AssignOrderNumber(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/assign_order_number", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Cancel(id string, params *order.CancelRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/cancel", url.PathEscape(id)), params).SetIdempotency(true)
}
func CreateRefundableCreditNote(id string, params *order.CreateRefundableCreditNoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/create_refundable_credit_note", url.PathEscape(id)), params).SetIdempotency(true)
}
func Reopen(id string, params *order.ReopenRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/reopen", url.PathEscape(id)), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/orders/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func List(params *order.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/orders"), params)
}

// Deprecated: This function is deprecated.
func OrdersForInvoice(id string, params *order.OrdersForInvoiceRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/orders", url.PathEscape(id)), params)
}
func Resend(id string, params *order.ResendRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/resend", url.PathEscape(id)), params).SetIdempotency(true)
}

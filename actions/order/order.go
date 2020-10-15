package order

import (
	"fmt"

	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/order"
)

func Create(params *order.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders"), params)
}
func Update(id string, params *order.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v", id), params)
}
func ImportOrder(params *order.ImportOrderRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/import_order"), params)
}
func AssignOrderNumber(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/assign_order_number", id), nil)
}
func Cancel(id string, params *order.CancelRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/cancel", id), params)
}
func CreateRefundableCreditNote(id string, params *order.CreateRefundableCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/create_refundable_credit_note", id), params)
}
func Reopen(id string, params *order.ReopenRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/reopen", id), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/orders/%v", id), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/delete", id), nil)
}
func List(params *order.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/orders"), params)
}
func OrdersForInvoice(id string, params *order.OrdersForInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/orders", id), params)
}

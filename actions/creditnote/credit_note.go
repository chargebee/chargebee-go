package creditnote

import (
	"fmt"

	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/creditnote"
)

func Create(params *creditnote.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v", id), nil)
}
func Pdf(id string, params *creditnote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/pdf", id), params)
}
func Refund(id string, params *creditnote.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/refund", id), params)
}
func RecordRefund(id string, params *creditnote.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/record_refund", id), params)
}
func VoidCreditNote(id string, params *creditnote.VoidCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/void", id), params)
}
func List(params *creditnote.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/credit_notes"), params)
}
func CreditNotesForCustomer(id string, params *creditnote.CreditNotesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/credit_notes", id), params)
}
func Delete(id string, params *creditnote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/delete", id), params)
}

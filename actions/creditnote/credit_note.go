package creditnote

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/creditnote"
	"net/url"
)

func Create(params *creditnote.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v", url.PathEscape(id)), nil)
}
func Pdf(id string, params *creditnote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/pdf", url.PathEscape(id)), params)
}
func DownloadEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v/download_einvoice", url.PathEscape(id)), nil)
}
func Refund(id string, params *creditnote.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/refund", url.PathEscape(id)), params)
}
func RecordRefund(id string, params *creditnote.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/record_refund", url.PathEscape(id)), params)
}
func VoidCreditNote(id string, params *creditnote.VoidCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/void", url.PathEscape(id)), params)
}
func List(params *creditnote.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/credit_notes"), params)
}
func CreditNotesForCustomer(id string, params *creditnote.CreditNotesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/credit_notes", url.PathEscape(id)), params)
}
func Delete(id string, params *creditnote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/delete", url.PathEscape(id)), params)
}

package creditnote

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/creditnote"
	"net/url"
)

func Create(params *creditnote.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes"), params).SetIdempotency(true)
}
func Retrieve(id string, params *creditnote.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v", url.PathEscape(id)), params)
}
func Pdf(id string, params *creditnote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/pdf", url.PathEscape(id)), params).SetIdempotency(true)
}
func DownloadEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v/download_einvoice", url.PathEscape(id)), nil)
}
func Refund(id string, params *creditnote.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/refund", url.PathEscape(id)), params).SetIdempotency(true)
}
func RecordRefund(id string, params *creditnote.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/record_refund", url.PathEscape(id)), params).SetIdempotency(true)
}
func VoidCreditNote(id string, params *creditnote.VoidCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/void", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *creditnote.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/credit_notes"), params)
}
func CreditNotesForCustomer(id string, params *creditnote.CreditNotesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/credit_notes", url.PathEscape(id)), params)
}
func Delete(id string, params *creditnote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/delete", url.PathEscape(id)), params).SetIdempotency(true)
}
func RemoveTaxWithheldRefund(id string, params *creditnote.RemoveTaxWithheldRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/remove_tax_withheld_refund", url.PathEscape(id)), params).SetIdempotency(true)
}
func ResendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/resend_einvoice", url.PathEscape(id)), nil).SetIdempotency(true)
}
func SendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/send_einvoice", url.PathEscape(id)), nil).SetIdempotency(true)
}
func ImportCreditNote(params *creditnote.ImportCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/import_credit_note"), params).SetIdempotency(true)
}

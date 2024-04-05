package creditnote

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/creditnote"
)

func Create(params *creditnote.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v", chargebee.IDEscape(id)), nil)
}
func Pdf(id string, params *creditnote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/pdf", chargebee.IDEscape(id)), params)
}
func DownloadEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/credit_notes/%v/download_einvoice", chargebee.IDEscape(id)), nil)
}
func Refund(id string, params *creditnote.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/refund", chargebee.IDEscape(id)), params)
}
func RecordRefund(id string, params *creditnote.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/record_refund", chargebee.IDEscape(id)), params)
}
func VoidCreditNote(id string, params *creditnote.VoidCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/void", chargebee.IDEscape(id)), params)
}
func List(params *creditnote.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/credit_notes"), params)
}
func CreditNotesForCustomer(id string, params *creditnote.CreditNotesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/credit_notes", chargebee.IDEscape(id)), params)
}
func Delete(id string, params *creditnote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/delete", chargebee.IDEscape(id)), params)
}
func RemoveTaxWithheldRefund(id string, params *creditnote.RemoveTaxWithheldRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/remove_tax_withheld_refund", chargebee.IDEscape(id)), params)
}
func ResendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/resend_einvoice", chargebee.IDEscape(id)), nil)
}
func SendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/%v/send_einvoice", chargebee.IDEscape(id)), nil)
}
func ImportCreditNote(params *creditnote.ImportCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/credit_notes/import_credit_note"), params)
}

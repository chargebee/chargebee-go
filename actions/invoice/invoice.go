package invoice

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/invoice"
)

func Create(params *invoice.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices"), params)
}
func CreateForChargeItemsAndCharges(params *invoice.CreateForChargeItemsAndChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/create_for_charge_items_and_charges"), params)
}
func Charge(params *invoice.ChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/charge"), params)
}
func ChargeAddon(params *invoice.ChargeAddonRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/charge_addon"), params)
}
func CreateForChargeItem(params *invoice.CreateForChargeItemRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/create_for_charge_item"), params)
}
func StopDunning(id string, params *invoice.StopDunningRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/stop_dunning", chargebee.IDEscape(id)), params)
}
func ImportInvoice(params *invoice.ImportInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/import_invoice"), params)
}
func ApplyPayments(id string, params *invoice.ApplyPaymentsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_payments", chargebee.IDEscape(id)), params)
}
func SyncUsages(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/sync_usages", chargebee.IDEscape(id)), nil)
}
func DeleteLineItems(id string, params *invoice.DeleteLineItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete_line_items", chargebee.IDEscape(id)), params)
}
func ApplyCredits(id string, params *invoice.ApplyCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_credits", chargebee.IDEscape(id)), params)
}
func List(params *invoice.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices"), params)
}
func InvoicesForCustomer(id string, params *invoice.InvoicesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/invoices", chargebee.IDEscape(id)), params)
}
func InvoicesForSubscription(id string, params *invoice.InvoicesForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/invoices", chargebee.IDEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v", chargebee.IDEscape(id)), nil)
}
func Pdf(id string, params *invoice.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/pdf", chargebee.IDEscape(id)), params)
}
func DownloadEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v/download_einvoice", chargebee.IDEscape(id)), nil)
}
func ListPaymentReferenceNumbers(params *invoice.ListPaymentReferenceNumbersRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/payment_reference_numbers"), params)
}
func AddCharge(id string, params *invoice.AddChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge", chargebee.IDEscape(id)), params)
}
func AddAddonCharge(id string, params *invoice.AddAddonChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_addon_charge", chargebee.IDEscape(id)), params)
}
func AddChargeItem(id string, params *invoice.AddChargeItemRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge_item", chargebee.IDEscape(id)), params)
}
func Close(id string, params *invoice.CloseRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/close", chargebee.IDEscape(id)), params)
}
func CollectPayment(id string, params *invoice.CollectPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/collect_payment", chargebee.IDEscape(id)), params)
}
func RecordPayment(id string, params *invoice.RecordPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_payment", chargebee.IDEscape(id)), params)
}
func RecordTaxWithheld(id string, params *invoice.RecordTaxWithheldRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_tax_withheld", chargebee.IDEscape(id)), params)
}
func RemoveTaxWithheld(id string, params *invoice.RemoveTaxWithheldRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_tax_withheld", chargebee.IDEscape(id)), params)
}
func Refund(id string, params *invoice.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/refund", chargebee.IDEscape(id)), params)
}
func RecordRefund(id string, params *invoice.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_refund", chargebee.IDEscape(id)), params)
}
func RemovePayment(id string, params *invoice.RemovePaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_payment", chargebee.IDEscape(id)), params)
}
func RemoveCreditNote(id string, params *invoice.RemoveCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_credit_note", chargebee.IDEscape(id)), params)
}
func VoidInvoice(id string, params *invoice.VoidInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/void", chargebee.IDEscape(id)), params)
}
func WriteOff(id string, params *invoice.WriteOffRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/write_off", chargebee.IDEscape(id)), params)
}
func Delete(id string, params *invoice.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete", chargebee.IDEscape(id)), params)
}
func UpdateDetails(id string, params *invoice.UpdateDetailsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/update_details", chargebee.IDEscape(id)), params)
}
func ResendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/resend_einvoice", chargebee.IDEscape(id)), nil)
}
func SendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/send_einvoice", chargebee.IDEscape(id)), nil)
}

package invoice

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/invoice"
	"net/url"
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
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/stop_dunning", url.PathEscape(id)), params)
}
func ImportInvoice(params *invoice.ImportInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/import_invoice"), params)
}
func ApplyPayments(id string, params *invoice.ApplyPaymentsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_payments", url.PathEscape(id)), params)
}
func SyncUsages(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/sync_usages", url.PathEscape(id)), nil)
}
func DeleteLineItems(id string, params *invoice.DeleteLineItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete_line_items", url.PathEscape(id)), params)
}
func ApplyCredits(id string, params *invoice.ApplyCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_credits", url.PathEscape(id)), params)
}
func List(params *invoice.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices"), params)
}
func InvoicesForCustomer(id string, params *invoice.InvoicesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/invoices", url.PathEscape(id)), params)
}
func InvoicesForSubscription(id string, params *invoice.InvoicesForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/invoices", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v", url.PathEscape(id)), nil)
}
func Pdf(id string, params *invoice.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/pdf", url.PathEscape(id)), params)
}
func DownloadEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v/download_einvoice", url.PathEscape(id)), nil)
}
func AddCharge(id string, params *invoice.AddChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge", url.PathEscape(id)), params)
}
func AddAddonCharge(id string, params *invoice.AddAddonChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_addon_charge", url.PathEscape(id)), params)
}
func AddChargeItem(id string, params *invoice.AddChargeItemRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge_item", url.PathEscape(id)), params)
}
func Close(id string, params *invoice.CloseRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/close", url.PathEscape(id)), params)
}
func CollectPayment(id string, params *invoice.CollectPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/collect_payment", url.PathEscape(id)), params)
}
func RecordPayment(id string, params *invoice.RecordPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_payment", url.PathEscape(id)), params)
}
func Refund(id string, params *invoice.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/refund", url.PathEscape(id)), params)
}
func RecordRefund(id string, params *invoice.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_refund", url.PathEscape(id)), params)
}
func RemovePayment(id string, params *invoice.RemovePaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_payment", url.PathEscape(id)), params)
}
func RemoveCreditNote(id string, params *invoice.RemoveCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_credit_note", url.PathEscape(id)), params)
}
func VoidInvoice(id string, params *invoice.VoidInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/void", url.PathEscape(id)), params)
}
func WriteOff(id string, params *invoice.WriteOffRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/write_off", url.PathEscape(id)), params)
}
func Delete(id string, params *invoice.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete", url.PathEscape(id)), params)
}
func UpdateDetails(id string, params *invoice.UpdateDetailsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/update_details", url.PathEscape(id)), params)
}
func ResendEinvoice(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/resend_einvoice", url.PathEscape(id)), nil)
}

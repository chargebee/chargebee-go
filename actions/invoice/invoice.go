package invoice

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/invoice"
)

func Create(params *invoice.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices"), params)
}
func Charge(params *invoice.ChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/charge"), params)
}
func ChargeAddon(params *invoice.ChargeAddonRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/charge_addon"), params)
}
func StopDunning(id string, params *invoice.StopDunningRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/stop_dunning", id), params)
}
func ImportInvoice(params *invoice.ImportInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/import_invoice"), params)
}
func ApplyPayments(id string, params *invoice.ApplyPaymentsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_payments", id), params)
}
func ApplyCredits(id string, params *invoice.ApplyCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_credits", id), params)
}
func List(params *invoice.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices"), params)
}
func InvoicesForCustomer(id string, params *invoice.InvoicesForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/invoices", id), params)
}
func InvoicesForSubscription(id string, params *invoice.InvoicesForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/invoices", id), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v", id), nil)
}
func Pdf(id string, params *invoice.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/pdf", id), params)
}
func AddCharge(id string, params *invoice.AddChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge", id), params)
}
func AddAddonCharge(id string, params *invoice.AddAddonChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_addon_charge", id), params)
}
func Close(id string, params *invoice.CloseRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/close", id), params)
}
func CollectPayment(id string, params *invoice.CollectPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/collect_payment", id), params)
}
func RecordPayment(id string, params *invoice.RecordPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_payment", id), params)
}
func Refund(id string, params *invoice.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/refund", id), params)
}
func RecordRefund(id string, params *invoice.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_refund", id), params)
}
func RemovePayment(id string, params *invoice.RemovePaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_payment", id), params)
}
func RemoveCreditNote(id string, params *invoice.RemoveCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_credit_note", id), params)
}
func VoidInvoice(id string, params *invoice.VoidInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/void", id), params)
}
func WriteOff(id string, params *invoice.WriteOffRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/write_off", id), params)
}
func Delete(id string, params *invoice.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete", id), params)
}
func UpdateDetails(id string, params *invoice.UpdateDetailsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/update_details", id), params)
}

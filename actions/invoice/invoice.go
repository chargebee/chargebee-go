package invoice

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/invoice"
	"net/url"
)

func Create(params *invoice.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices"), params).WithTelemetryResource("invoice").WithTelemetryOperation("create").SetIdempotency(true)
}
func CreateForChargeItemsAndCharges(params *invoice.CreateForChargeItemsAndChargesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/create_for_charge_items_and_charges"), params).WithTelemetryResource("invoice").WithTelemetryOperation("createForChargeItemsAndCharges").SetIdempotency(true)
}
func Charge(params *invoice.ChargeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/charge"), params).WithTelemetryResource("invoice").WithTelemetryOperation("charge").SetIdempotency(true)
}
func ChargeAddon(params *invoice.ChargeAddonRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/charge_addon"), params).WithTelemetryResource("invoice").WithTelemetryOperation("chargeAddon").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func CreateForChargeItem(params *invoice.CreateForChargeItemRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/create_for_charge_item"), params).WithTelemetryResource("invoice").WithTelemetryOperation("createForChargeItem").SetIdempotency(true)
}
func StopDunning(id string, params *invoice.StopDunningRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/stop_dunning", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("stopDunning").SetIdempotency(true)
}
func PauseDunning(id string, params *invoice.PauseDunningRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/pause_dunning", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("pauseDunning").SetIdempotency(true)
}
func ResumeDunning(id string, params *invoice.ResumeDunningRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/resume_dunning", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("resumeDunning").SetIdempotency(true)
}
func ImportInvoice(params *invoice.ImportInvoiceRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/import_invoice"), params).WithTelemetryResource("invoice").WithTelemetryOperation("importInvoice").SetIdempotency(true)
}
func ApplyPayments(id string, params *invoice.ApplyPaymentsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_payments", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("applyPayments").SetIdempotency(true)
}
func SyncUsages(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/sync_usages", url.PathEscape(id)), nil).WithTelemetryResource("invoice").WithTelemetryOperation("syncUsages").SetIdempotency(true)
}
func DeleteLineItems(id string, params *invoice.DeleteLineItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete_line_items", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("deleteLineItems").SetIdempotency(true)
}
func ApplyCredits(id string, params *invoice.ApplyCreditsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_credits", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("applyCredits").SetIdempotency(true)
}
func List(params *invoice.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices"), params).WithTelemetryResource("invoice").WithTelemetryOperation("list")
}

// Deprecated: This function is deprecated.
func InvoicesForCustomer(id string, params *invoice.InvoicesForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/invoices", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("invoicesForCustomer")
}

// Deprecated: This function is deprecated.
func InvoicesForSubscription(id string, params *invoice.InvoicesForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/invoices", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("invoicesForSubscription")
}
func Retrieve(id string, params *invoice.RetrieveRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("retrieve")
}
func Pdf(id string, params *invoice.PdfRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/pdf", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("pdf").SetIdempotency(true)
}
func DownloadEinvoice(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v/download_einvoice", url.PathEscape(id)), nil).WithTelemetryResource("invoice").WithTelemetryOperation("downloadEinvoice")
}
func ListPaymentReferenceNumbers(params *invoice.ListPaymentReferenceNumbersRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/payment_reference_numbers"), params).WithTelemetryResource("invoice").WithTelemetryOperation("listPaymentReferenceNumbers")
}
func AddCharge(id string, params *invoice.AddChargeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("addCharge").SetIdempotency(true)
}
func AddAddonCharge(id string, params *invoice.AddAddonChargeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_addon_charge", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("addAddonCharge").SetIdempotency(true)
}
func AddChargeItem(id string, params *invoice.AddChargeItemRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/add_charge_item", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("addChargeItem").SetIdempotency(true)
}
func Close(id string, params *invoice.CloseRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/close", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("close").SetIdempotency(true)
}
func CollectPayment(id string, params *invoice.CollectPaymentRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/collect_payment", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("collectPayment").SetIdempotency(true)
}
func RecordPayment(id string, params *invoice.RecordPaymentRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_payment", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("recordPayment").SetIdempotency(true)
}
func RecordTaxWithheld(id string, params *invoice.RecordTaxWithheldRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_tax_withheld", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("recordTaxWithheld").SetIdempotency(true)
}
func RemoveTaxWithheld(id string, params *invoice.RemoveTaxWithheldRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_tax_withheld", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("removeTaxWithheld").SetIdempotency(true)
}
func Refund(id string, params *invoice.RefundRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/refund", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("refund").SetIdempotency(true)
}
func RecordRefund(id string, params *invoice.RecordRefundRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/record_refund", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("recordRefund").SetIdempotency(true)
}
func RemovePayment(id string, params *invoice.RemovePaymentRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_payment", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("removePayment").SetIdempotency(true)
}
func RemoveCreditNote(id string, params *invoice.RemoveCreditNoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/remove_credit_note", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("removeCreditNote").SetIdempotency(true)
}
func VoidInvoice(id string, params *invoice.VoidInvoiceRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/void", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("voidInvoice").SetIdempotency(true)
}
func WriteOff(id string, params *invoice.WriteOffRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/write_off", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("writeOff").SetIdempotency(true)
}
func Delete(id string, params *invoice.DeleteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/delete", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("delete").SetIdempotency(true)
}
func UpdateDetails(id string, params *invoice.UpdateDetailsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/update_details", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("updateDetails").SetIdempotency(true)
}
func ApplyPaymentScheduleScheme(id string, params *invoice.ApplyPaymentScheduleSchemeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/apply_payment_schedule_scheme", url.PathEscape(id)), params).WithTelemetryResource("invoice").WithTelemetryOperation("applyPaymentScheduleScheme").SetIdempotency(true)
}
func PaymentSchedules(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/invoices/%v/payment_schedules", url.PathEscape(id)), nil).WithTelemetryResource("invoice").WithTelemetryOperation("paymentSchedules")
}
func ResendEinvoice(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/resend_einvoice", url.PathEscape(id)), nil).WithTelemetryResource("invoice").WithTelemetryOperation("resendEinvoice").SetIdempotency(true)
}
func SendEinvoice(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/invoices/%v/send_einvoice", url.PathEscape(id)), nil).WithTelemetryResource("invoice").WithTelemetryOperation("sendEinvoice").SetIdempotency(true)
}

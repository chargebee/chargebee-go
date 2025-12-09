package chargebee

import (
	"fmt"
	"net/url"
)

type InvoiceService struct {
	transport *Transport
}

func (s *InvoiceService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices"), req).SetIdempotency(true)
}

func (s *InvoiceService) CreateForChargeItemsAndCharges(req *CreateForChargeItemsAndChargesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/create_for_charge_items_and_charges"), req).SetIdempotency(true)
}

func (s *InvoiceService) Charge(req *ChargeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/charge"), req).SetIdempotency(true)
}

func (s *InvoiceService) ChargeAddon(req *ChargeAddonRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/charge_addon"), req).SetIdempotency(true)
}

func (s *InvoiceService) CreateForChargeItem(req *CreateForChargeItemRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/create_for_charge_item"), req).SetIdempotency(true)
}

func (s *InvoiceService) StopDunning(id string, req *StopDunningRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/stop_dunning", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) PauseDunning(id string, req *PauseDunningRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/pause_dunning", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) ResumeDunning(id string, req *ResumeDunningRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/resume_dunning", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) ImportInvoice(req *ImportInvoiceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/import_invoice"), req).SetIdempotency(true)
}

func (s *InvoiceService) ApplyPayments(id string, req *ApplyPaymentsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/apply_payments", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) SyncUsages(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/sync_usages", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *InvoiceService) DeleteLineItems(id string, req *DeleteLineItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/delete_line_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) ApplyCredits(id string, req *ApplyCreditsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/apply_credits", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/invoices"), req)
}

func (s *InvoiceService) InvoicesForCustomer(id string, req *InvoicesForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/invoices", url.PathEscape(id)), req)
}

func (s *InvoiceService) InvoicesForSubscription(id string, req *InvoicesForSubscriptionRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions/%v/invoices", url.PathEscape(id)), req)
}

func (s *InvoiceService) Retrieve(id string, req *RetrieveRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/invoices/%v", url.PathEscape(id)), req)
}

func (s *InvoiceService) Pdf(id string, req *PdfRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/pdf", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) DownloadEinvoice(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/invoices/%v/download_einvoice", url.PathEscape(id)), nil)
}

func (s *InvoiceService) ListPaymentReferenceNumbers(req *ListPaymentReferenceNumbersRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/invoices/payment_reference_numbers"), req)
}

func (s *InvoiceService) AddCharge(id string, req *AddChargeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/add_charge", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) AddAddonCharge(id string, req *AddAddonChargeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/add_addon_charge", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) AddChargeItem(id string, req *AddChargeItemRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/add_charge_item", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) Close(id string, req *CloseRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/close", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) CollectPayment(id string, req *CollectPaymentRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/collect_payment", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) RecordPayment(id string, req *RecordPaymentRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/record_payment", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) RecordTaxWithheld(id string, req *RecordTaxWithheldRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/record_tax_withheld", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) RemoveTaxWithheld(id string, req *RemoveTaxWithheldRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/remove_tax_withheld", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) Refund(id string, req *RefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) RecordRefund(id string, req *RecordRefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/record_refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) RemovePayment(id string, req *RemovePaymentRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/remove_payment", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) RemoveCreditNote(id string, req *RemoveCreditNoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/remove_credit_note", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) VoidInvoice(id string, req *VoidInvoiceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/void", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) WriteOff(id string, req *WriteOffRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/write_off", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/delete", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) UpdateDetails(id string, req *UpdateDetailsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/update_details", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) ApplyPaymentScheduleScheme(id string, req *ApplyPaymentScheduleSchemeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/apply_payment_schedule_scheme", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *InvoiceService) PaymentSchedules(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/invoices/%v/payment_schedules", url.PathEscape(id)), nil)
}

func (s *InvoiceService) ResendEinvoice(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/resend_einvoice", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *InvoiceService) SendEinvoice(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/invoices/%v/send_einvoice", url.PathEscape(id)), nil).SetIdempotency(true)
}

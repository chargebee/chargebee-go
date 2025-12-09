package chargebee

import (
	"fmt"
	"net/url"
)

type CreditNoteService struct {
	transport *Transport
}

func (s *CreditNoteService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes"), req).SetIdempotency(true)
}

func (s *CreditNoteService) Retrieve(id string, req *RetrieveRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/credit_notes/%v", url.PathEscape(id)), req)
}

func (s *CreditNoteService) Pdf(id string, req *PdfRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/pdf", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CreditNoteService) DownloadEinvoice(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/credit_notes/%v/download_einvoice", url.PathEscape(id)), nil)
}

func (s *CreditNoteService) Refund(id string, req *RefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CreditNoteService) RecordRefund(id string, req *RecordRefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/record_refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CreditNoteService) VoidCreditNote(id string, req *VoidCreditNoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/void", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CreditNoteService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/credit_notes"), req)
}

func (s *CreditNoteService) CreditNotesForCustomer(id string, req *CreditNotesForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/credit_notes", url.PathEscape(id)), req)
}

func (s *CreditNoteService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/delete", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CreditNoteService) RemoveTaxWithheldRefund(id string, req *RemoveTaxWithheldRefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/remove_tax_withheld_refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CreditNoteService) ResendEinvoice(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/resend_einvoice", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *CreditNoteService) SendEinvoice(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/%v/send_einvoice", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *CreditNoteService) ImportCreditNote(req *ImportCreditNoteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/credit_notes/import_credit_note"), req).SetIdempotency(true)
}

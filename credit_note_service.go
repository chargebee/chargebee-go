package chargebee

import (
	"fmt"
	"net/url"
)

type CreditNoteService struct {
	config *ClientConfig
}

func (s *CreditNoteService) Create(req *CreditNoteCreateRequest) (*CreditNoteCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes")
	req.isIdempotent = true
	return send[*CreditNoteCreateResponse](req, s.config)
}

func (s *CreditNoteService) Retrieve(id string, req *CreditNoteRetrieveRequest) (*CreditNoteRetrieveResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/credit_notes/%v", url.PathEscape(id))
	return send[*CreditNoteRetrieveResponse](req, s.config)
}

func (s *CreditNoteService) Pdf(id string, req *CreditNotePdfRequest) (*CreditNotePdfResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/pdf", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNotePdfResponse](req, s.config)
}

func (s *CreditNoteService) DownloadEinvoice(id string) (*CreditNoteDownloadEinvoiceResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/credit_notes/%v/download_einvoice", url.PathEscape(id))
	return send[*CreditNoteDownloadEinvoiceResponse](req, s.config)
}

func (s *CreditNoteService) Refund(id string, req *CreditNoteRefundRequest) (*CreditNoteRefundResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/refund", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteRefundResponse](req, s.config)
}

func (s *CreditNoteService) RecordRefund(id string, req *CreditNoteRecordRefundRequest) (*CreditNoteRecordRefundResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/record_refund", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteRecordRefundResponse](req, s.config)
}

func (s *CreditNoteService) VoidCreditNote(id string, req *CreditNoteVoidCreditNoteRequest) (*CreditNoteVoidCreditNoteResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/void", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteVoidCreditNoteResponse](req, s.config)
}

func (s *CreditNoteService) List(req *CreditNoteListRequest) (*CreditNoteListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/credit_notes")
	req.isListRequest = true
	return send[*CreditNoteListResponse](req, s.config)
}

func (s *CreditNoteService) CreditNotesForCustomer(id string, req *CreditNoteCreditNotesForCustomerRequest) (*CreditNoteCreditNotesForCustomerResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/customers/%v/credit_notes", url.PathEscape(id))
	req.isListRequest = true
	return send[*CreditNoteCreditNotesForCustomerResponse](req, s.config)
}

func (s *CreditNoteService) Delete(id string, req *CreditNoteDeleteRequest) (*CreditNoteDeleteResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteDeleteResponse](req, s.config)
}

func (s *CreditNoteService) RemoveTaxWithheldRefund(id string, req *CreditNoteRemoveTaxWithheldRefundRequest) (*CreditNoteRemoveTaxWithheldRefundResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/remove_tax_withheld_refund", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteRemoveTaxWithheldRefundResponse](req, s.config)
}

func (s *CreditNoteService) ResendEinvoice(id string) (*CreditNoteResendEinvoiceResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/resend_einvoice", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteResendEinvoiceResponse](req, s.config)
}

func (s *CreditNoteService) SendEinvoice(id string) (*CreditNoteSendEinvoiceResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/%v/send_einvoice", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CreditNoteSendEinvoiceResponse](req, s.config)
}

func (s *CreditNoteService) ImportCreditNote(req *CreditNoteImportCreditNoteRequest) (*CreditNoteImportCreditNoteResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/credit_notes/import_credit_note")
	req.isIdempotent = true
	return send[*CreditNoteImportCreditNoteResponse](req, s.config)
}

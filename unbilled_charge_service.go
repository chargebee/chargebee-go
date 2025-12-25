package chargebee

import (
	"fmt"
	"net/url"
)

type UnbilledChargeService struct {
	config *ClientConfig
}

func (s *UnbilledChargeService) CreateUnbilledCharge(req *UnbilledChargeCreateUnbilledChargeRequest) (*UnbilledChargeCreateUnbilledChargeResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/unbilled_charges/create")
	req.isIdempotent = true
	return send[*UnbilledChargeCreateUnbilledChargeResponse](req, s.config)
}

func (s *UnbilledChargeService) Create(req *UnbilledChargeCreateRequest) (*UnbilledChargeCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/unbilled_charges")
	req.isIdempotent = true
	return send[*UnbilledChargeCreateResponse](req, s.config)
}

func (s *UnbilledChargeService) InvoiceUnbilledCharges(req *UnbilledChargeInvoiceUnbilledChargesRequest) (*UnbilledChargeInvoiceUnbilledChargesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/unbilled_charges/invoice_unbilled_charges")
	req.isIdempotent = true
	return send[*UnbilledChargeInvoiceUnbilledChargesResponse](req, s.config)
}

func (s *UnbilledChargeService) Delete(id string) (*UnbilledChargeDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/unbilled_charges/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*UnbilledChargeDeleteResponse](req, s.config)
}

func (s *UnbilledChargeService) List(req *UnbilledChargeListRequest) (*UnbilledChargeListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/unbilled_charges")
	req.isListRequest = true
	return send[*UnbilledChargeListResponse](req, s.config)
}

func (s *UnbilledChargeService) InvoiceNowEstimate(req *UnbilledChargeInvoiceNowEstimateRequest) (*UnbilledChargeInvoiceNowEstimateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/unbilled_charges/invoice_now_estimate")
	return send[*UnbilledChargeInvoiceNowEstimateResponse](req, s.config)
}

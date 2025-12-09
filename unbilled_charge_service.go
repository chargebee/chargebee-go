package chargebee

import (
	"fmt"
	"net/url"
)

type UnbilledChargeService struct {
	transport *Transport
}

func (s *UnbilledChargeService) CreateUnbilledCharge(req *CreateUnbilledChargeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/unbilled_charges/create"), req).SetIdempotency(true)
}

func (s *UnbilledChargeService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/unbilled_charges"), req).SetIdempotency(true)
}

func (s *UnbilledChargeService) InvoiceUnbilledCharges(req *InvoiceUnbilledChargesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_unbilled_charges"), req).SetIdempotency(true)
}

func (s *UnbilledChargeService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/unbilled_charges/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *UnbilledChargeService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/unbilled_charges"), req)
}

func (s *UnbilledChargeService) InvoiceNowEstimate(req *InvoiceNowEstimateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_now_estimate"), req)
}

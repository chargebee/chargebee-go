package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentScheduleSchemeService struct {
	config *ClientConfig
}

func (s *PaymentScheduleSchemeService) Create(req *PaymentScheduleSchemeCreateRequest) (*PaymentScheduleSchemeCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_schedule_schemes")
	req.isIdempotent = true
	return send[*PaymentScheduleSchemeCreateResponse](req, s.config)
}

func (s *PaymentScheduleSchemeService) Retrieve(id string) (*PaymentScheduleSchemeRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/payment_schedule_schemes/%v", url.PathEscape(id))
	return send[*PaymentScheduleSchemeRetrieveResponse](req, s.config)
}

func (s *PaymentScheduleSchemeService) Delete(id string) (*PaymentScheduleSchemeDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_schedule_schemes/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentScheduleSchemeDeleteResponse](req, s.config)
}

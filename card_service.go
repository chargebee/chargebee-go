package chargebee

import (
	"fmt"
	"net/url"
)

type CardService struct {
	config *ClientConfig
}

func (s *CardService) Retrieve(id string) (*CardRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/cards/%v", url.PathEscape(id))
	return send[*CardRetrieveResponse](req, s.config)
}

func (s *CardService) UpdateCardForCustomer(id string, req *CardUpdateCardForCustomerRequest) (*CardUpdateCardForCustomerResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/credit_card", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CardUpdateCardForCustomerResponse](req, s.config)
}

func (s *CardService) SwitchGatewayForCustomer(id string, req *CardSwitchGatewayForCustomerRequest) (*CardSwitchGatewayForCustomerResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/switch_gateway", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CardSwitchGatewayForCustomerResponse](req, s.config)
}

func (s *CardService) CopyCardForCustomer(id string, req *CardCopyCardForCustomerRequest) (*CardCopyCardForCustomerResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/copy_card", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CardCopyCardForCustomerResponse](req, s.config)
}

func (s *CardService) DeleteCardForCustomer(id string) (*CardDeleteCardForCustomerResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/customers/%v/delete_card", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CardDeleteCardForCustomerResponse](req, s.config)
}

package chargebee

import (
	"fmt"
	"net/url"
)

type DifferentialPriceService struct {
	config *ClientConfig
}

func (s *DifferentialPriceService) Create(id string, req *DifferentialPriceCreateRequest) (*DifferentialPriceCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/item_prices/%v/differential_prices", url.PathEscape(id))
	req.isIdempotent = true
	return send[*DifferentialPriceCreateResponse](req, s.config)
}

func (s *DifferentialPriceService) Retrieve(id string, req *DifferentialPriceRetrieveRequest) (*DifferentialPriceRetrieveResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/differential_prices/%v", url.PathEscape(id))
	return send[*DifferentialPriceRetrieveResponse](req, s.config)
}

func (s *DifferentialPriceService) Update(id string, req *DifferentialPriceUpdateRequest) (*DifferentialPriceUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/differential_prices/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*DifferentialPriceUpdateResponse](req, s.config)
}

func (s *DifferentialPriceService) Delete(id string, req *DifferentialPriceDeleteRequest) (*DifferentialPriceDeleteResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/differential_prices/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*DifferentialPriceDeleteResponse](req, s.config)
}

func (s *DifferentialPriceService) List(req *DifferentialPriceListRequest) (*DifferentialPriceListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/differential_prices")
	req.isListRequest = true
	return send[*DifferentialPriceListResponse](req, s.config)
}

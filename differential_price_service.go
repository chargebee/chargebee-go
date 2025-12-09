package chargebee

import (
	"fmt"
	"net/url"
)

type DifferentialPriceService struct {
	transport *Transport
}

func (s *DifferentialPriceService) Create(id string, req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_prices/%v/differential_prices", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *DifferentialPriceService) Retrieve(id string, req *RetrieveRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/differential_prices/%v", url.PathEscape(id)), req)
}

func (s *DifferentialPriceService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/differential_prices/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *DifferentialPriceService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/differential_prices/%v/delete", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *DifferentialPriceService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/differential_prices"), req)
}

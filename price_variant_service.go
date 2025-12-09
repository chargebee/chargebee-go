package chargebee

import (
	"fmt"
	"net/url"
)

type PriceVariantService struct {
	transport *Transport
}

func (s *PriceVariantService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/price_variants"), req).SetIdempotency(true)
}

func (s *PriceVariantService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/price_variants/%v", url.PathEscape(id)), nil)
}

func (s *PriceVariantService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/price_variants/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PriceVariantService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/price_variants/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *PriceVariantService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/price_variants"), req)
}

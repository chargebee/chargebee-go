package chargebee

import (
	"fmt"
	"net/url"
)

type PriceVariantService struct {
	config *ClientConfig
}

func (s *PriceVariantService) Create(req *PriceVariantCreateRequest) (*PriceVariantCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/price_variants")
	req.isIdempotent = true
	return send[*PriceVariantCreateResponse](req, s.config)
}

func (s *PriceVariantService) Retrieve(id string) (*PriceVariantRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/price_variants/%v", url.PathEscape(id))
	return send[*PriceVariantRetrieveResponse](req, s.config)
}

func (s *PriceVariantService) Update(id string, req *PriceVariantUpdateRequest) (*PriceVariantUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/price_variants/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PriceVariantUpdateResponse](req, s.config)
}

func (s *PriceVariantService) Delete(id string) (*PriceVariantDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/price_variants/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PriceVariantDeleteResponse](req, s.config)
}

func (s *PriceVariantService) List(req *PriceVariantListRequest) (*PriceVariantListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/price_variants")
	req.isListRequest = true
	return send[*PriceVariantListResponse](req, s.config)
}

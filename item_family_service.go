package chargebee

import (
	"fmt"
	"net/url"
)

type ItemFamilyService struct {
	config *ClientConfig
}

func (s *ItemFamilyService) Create(req *ItemFamilyCreateRequest) (*ItemFamilyCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/item_families")
	req.isIdempotent = true
	return send[*ItemFamilyCreateResponse](req, s.config)
}

func (s *ItemFamilyService) Retrieve(id string) (*ItemFamilyRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/item_families/%v", url.PathEscape(id))
	return send[*ItemFamilyRetrieveResponse](req, s.config)
}

func (s *ItemFamilyService) List(req *ItemFamilyListRequest) (*ItemFamilyListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/item_families")
	req.isListRequest = true
	return send[*ItemFamilyListResponse](req, s.config)
}

func (s *ItemFamilyService) Update(id string, req *ItemFamilyUpdateRequest) (*ItemFamilyUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/item_families/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*ItemFamilyUpdateResponse](req, s.config)
}

func (s *ItemFamilyService) Delete(id string) (*ItemFamilyDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/item_families/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*ItemFamilyDeleteResponse](req, s.config)
}

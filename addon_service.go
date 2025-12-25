package chargebee

import (
	"fmt"
	"net/url"
)

type AddonService struct {
	config *ClientConfig
}

func (s *AddonService) Create(req *AddonCreateRequest) (*AddonCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/addons")
	req.isIdempotent = true
	return send[*AddonCreateResponse](req, s.config)
}

func (s *AddonService) Update(id string, req *AddonUpdateRequest) (*AddonUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/addons/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*AddonUpdateResponse](req, s.config)
}

func (s *AddonService) List(req *AddonListRequest) (*AddonListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/addons")
	req.isListRequest = true
	return send[*AddonListResponse](req, s.config)
}

func (s *AddonService) Retrieve(id string) (*AddonRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/addons/%v", url.PathEscape(id))
	return send[*AddonRetrieveResponse](req, s.config)
}

func (s *AddonService) Delete(id string) (*AddonDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/addons/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*AddonDeleteResponse](req, s.config)
}

func (s *AddonService) Copy(req *AddonCopyRequest) (*AddonCopyResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/addons/copy")
	req.isIdempotent = true
	return send[*AddonCopyResponse](req, s.config)
}

func (s *AddonService) Unarchive(id string) (*AddonUnarchiveResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/addons/%v/unarchive", url.PathEscape(id))
	req.isIdempotent = true
	return send[*AddonUnarchiveResponse](req, s.config)
}

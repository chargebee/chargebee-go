package chargebee

import (
	"fmt"
	"net/url"
)

type ItemService struct {
	config *ClientConfig
}

func (s *ItemService) Create(req *ItemCreateRequest) (*ItemCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/items")
	req.isIdempotent = true
	return send[*ItemCreateResponse](req, s.config)
}

func (s *ItemService) Retrieve(id string) (*ItemRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/items/%v", url.PathEscape(id))
	return send[*ItemRetrieveResponse](req, s.config)
}

func (s *ItemService) Update(id string, req *ItemUpdateRequest) (*ItemUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/items/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*ItemUpdateResponse](req, s.config)
}

func (s *ItemService) List(req *ItemListRequest) (*ItemListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/items")
	req.isListRequest = true
	return send[*ItemListResponse](req, s.config)
}

func (s *ItemService) Delete(id string) (*ItemDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/items/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*ItemDeleteResponse](req, s.config)
}

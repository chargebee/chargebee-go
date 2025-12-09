package chargebee

import (
	"fmt"
	"net/url"
)

type AttachedItemService struct {
	transport *Transport
}

func (s *AttachedItemService) Create(id string, req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/items/%v/attached_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *AttachedItemService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/attached_items/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *AttachedItemService) Retrieve(id string, req *RetrieveRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/attached_items/%v", url.PathEscape(id)), req)
}

func (s *AttachedItemService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/attached_items/%v/delete", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *AttachedItemService) List(id string, req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/items/%v/attached_items", url.PathEscape(id)), req)
}

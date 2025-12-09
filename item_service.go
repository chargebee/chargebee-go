package chargebee

import (
	"fmt"
	"net/url"
)

type ItemService struct {
	transport *Transport
}

func (s *ItemService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/items"), req).SetIdempotency(true)
}

func (s *ItemService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/items/%v", url.PathEscape(id)), nil)
}

func (s *ItemService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/items/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *ItemService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/items"), req)
}

func (s *ItemService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/items/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

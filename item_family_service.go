package chargebee

import (
	"fmt"
	"net/url"
)

type ItemFamilyService struct {
	transport *Transport
}

func (s *ItemFamilyService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_families"), req).SetIdempotency(true)
}

func (s *ItemFamilyService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/item_families/%v", url.PathEscape(id)), nil)
}

func (s *ItemFamilyService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/item_families"), req)
}

func (s *ItemFamilyService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_families/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *ItemFamilyService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_families/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

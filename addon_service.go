package chargebee

import (
	"fmt"
	"net/url"
)

type AddonService struct {
	transport *Transport
}

func (s *AddonService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/addons"), req).SetIdempotency(true)
}

func (s *AddonService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/addons/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *AddonService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/addons"), req)
}

func (s *AddonService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/addons/%v", url.PathEscape(id)), nil)
}

func (s *AddonService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/addons/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *AddonService) Copy(req *CopyRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/addons/copy"), req).SetIdempotency(true)
}

func (s *AddonService) Unarchive(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/addons/%v/unarchive", url.PathEscape(id)), nil).SetIdempotency(true)
}

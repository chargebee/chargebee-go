package chargebee

import (
	"fmt"
	"net/url"
)

type WebhookEndpointService struct {
	transport *Transport
}

func (s *WebhookEndpointService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/webhook_endpoints"), req).SetIdempotency(true)
}

func (s *WebhookEndpointService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/webhook_endpoints/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *WebhookEndpointService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/webhook_endpoints/%v", url.PathEscape(id)), nil)
}

func (s *WebhookEndpointService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/webhook_endpoints/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *WebhookEndpointService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/webhook_endpoints"), req)
}

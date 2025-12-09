package chargebee

import (
	"fmt"
	"net/url"
)

type UsageService struct {
	transport *Transport
}

func (s *UsageService) Create(id string, req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/usages", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *UsageService) Retrieve(id string, req *RetrieveRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/subscriptions/%v/usages", url.PathEscape(id)), req)
}

func (s *UsageService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/subscriptions/%v/delete_usage", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *UsageService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/usages"), req)
}

func (s *UsageService) Pdf(req *PdfRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/usages/pdf"), req).SetIdempotency(true)
}

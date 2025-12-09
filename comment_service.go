package chargebee

import (
	"fmt"
	"net/url"
)

type CommentService struct {
	transport *Transport
}

func (s *CommentService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/comments"), req).SetIdempotency(true)
}

func (s *CommentService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/comments/%v", url.PathEscape(id)), nil)
}

func (s *CommentService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/comments"), req)
}

func (s *CommentService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/comments/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

package chargebee

import (
	"fmt"
	"net/url"
)

type CommentService struct {
	config *ClientConfig
}

func (s *CommentService) Create(req *CommentCreateRequest) (*CommentCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/comments")
	req.isIdempotent = true
	return send[*CommentCreateResponse](req, s.config)
}

func (s *CommentService) Retrieve(id string) (*CommentRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/comments/%v", url.PathEscape(id))
	return send[*CommentRetrieveResponse](req, s.config)
}

func (s *CommentService) List(req *CommentListRequest) (*CommentListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/comments")
	req.isListRequest = true
	return send[*CommentListResponse](req, s.config)
}

func (s *CommentService) Delete(id string) (*CommentDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/comments/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CommentDeleteResponse](req, s.config)
}

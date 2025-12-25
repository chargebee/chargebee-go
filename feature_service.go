package chargebee

import (
	"fmt"
	"net/url"
)

type FeatureService struct {
	config *ClientConfig
}

func (s *FeatureService) List(req *FeatureListRequest) (*FeatureListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/features")
	req.isListRequest = true
	return send[*FeatureListResponse](req, s.config)
}

func (s *FeatureService) Create(req *FeatureCreateRequest) (*FeatureCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/features")
	req.isIdempotent = true
	return send[*FeatureCreateResponse](req, s.config)
}

func (s *FeatureService) Update(id string, req *FeatureUpdateRequest) (*FeatureUpdateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/features/%v", url.PathEscape(id))
	req.isIdempotent = true
	return send[*FeatureUpdateResponse](req, s.config)
}

func (s *FeatureService) Retrieve(id string) (*FeatureRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/features/%v", url.PathEscape(id))
	return send[*FeatureRetrieveResponse](req, s.config)
}

func (s *FeatureService) Delete(id string) (*FeatureDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/features/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*FeatureDeleteResponse](req, s.config)
}

func (s *FeatureService) Activate(id string) (*FeatureActivateResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/features/%v/activate_command", url.PathEscape(id))
	req.isIdempotent = true
	return send[*FeatureActivateResponse](req, s.config)
}

func (s *FeatureService) Archive(id string) (*FeatureArchiveResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/features/%v/archive_command", url.PathEscape(id))
	req.isIdempotent = true
	return send[*FeatureArchiveResponse](req, s.config)
}

func (s *FeatureService) Reactivate(id string) (*FeatureReactivateResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/features/%v/reactivate_command", url.PathEscape(id))
	req.isIdempotent = true
	return send[*FeatureReactivateResponse](req, s.config)
}

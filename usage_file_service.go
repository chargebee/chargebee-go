package chargebee

import (
	"fmt"
	"net/url"
)

type UsageFileService struct {
	config *ClientConfig
}

func (s *UsageFileService) UploadUrl(req *UsageFileUploadUrlRequest) (*UsageFileUploadUrlResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/usage_files/upload_url")
	return send[*UsageFileUploadUrlResponse](req, s.config)
}

func (s *UsageFileService) ProcessingStatus(id string) (*UsageFileProcessingStatusResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/usage_files/%v/processing_status", url.PathEscape(id))
	return send[*UsageFileProcessingStatusResponse](req, s.config)
}

package chargebee

import (
	"fmt"
	"net/url"
)

type UsageFileService struct {
	transport *Transport
}

func (s *UsageFileService) UploadUrl(req *UploadUrlRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/usage_files/upload_url"), req).SetSubDomain("file-ingest")
}

func (s *UsageFileService) ProcessingStatus(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/usage_files/%v/processing_status", url.PathEscape(id)), nil).SetSubDomain("file-ingest")
}

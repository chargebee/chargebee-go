package chargebee

import (
	"fmt"
)

type UsageEventService struct {
	transport *Transport
}

func (s *UsageEventService) Create(req *CreateRequest) Request {
	return s.transport.SendJsonRequest("POST", fmt.Sprintf("/usage_events"), req).SetSubDomain("ingest")
}

func (s *UsageEventService) BatchIngest(req *BatchIngestRequest) Request {
	return s.transport.SendJsonRequest("POST", fmt.Sprintf("/batch/usage_events"), req).SetSubDomain("ingest")
}

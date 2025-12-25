package chargebee

import (
	"encoding/json"
	"fmt"

	"net/url"

	"errors"
	"strings"
)

type EventService struct {
	config *ClientConfig
}

func (s *EventService) List(req *EventListRequest) (*EventListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/events")
	req.isListRequest = true
	return send[*EventListResponse](req, s.config)
}

func (s *EventService) Retrieve(id string) (*EventRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/events/%v", url.PathEscape(id))
	return send[*EventRetrieveResponse](req, s.config)
}

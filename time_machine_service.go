package chargebee

import (
	"encoding/json"
	"fmt"

	"errors"
	"time"

	"net/url"
)

type TimeMachineService struct {
	config *ClientConfig
}

func (s *TimeMachineService) Retrieve(id string) (*TimeMachineRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/time_machines/%v", url.PathEscape(id))
	return send[*TimeMachineRetrieveResponse](req, s.config)
}

func (s *TimeMachineService) StartAfresh(id string, req *TimeMachineStartAfreshRequest) (*TimeMachineStartAfreshResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/time_machines/%v/start_afresh", url.PathEscape(id))
	req.isIdempotent = true
	return send[*TimeMachineStartAfreshResponse](req, s.config)
}

func (s *TimeMachineService) TravelForward(id string, req *TimeMachineTravelForwardRequest) (*TimeMachineTravelForwardResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/time_machines/%v/travel_forward", url.PathEscape(id))
	req.isIdempotent = true
	return send[*TimeMachineTravelForwardResponse](req, s.config)
}

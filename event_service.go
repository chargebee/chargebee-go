package chargebee

import (
	"encoding/json"
	"fmt"

	"net/url"

	"errors"
	"strings"
)

type EventService struct {
	transport *Transport
}

func (s *EventService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/events"), req)
}

func (s *EventService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/events/%v", url.PathEscape(id)), nil)
}

func Content(event Event) *chargebee.Result {
	content := &chargebee.Result{}
	err1 := json.Unmarshal(Content, content)
	if err1 != nil {
		panic(err1)
	}
	return content
}

func Deserialize(jsonObj string) *Event {
	event := &Event{}
	err := json.Unmarshal([]byte(jsonObj), event)
	if err != nil {
		panic(errors.New("Response not in JSON format. Might not be a ChargeBee webhook call"))
	}
	apiVersion := ApiVersion
	envVersion := chargebee.APIVersion
	if apiVersion != "" && strings.ToUpper(string(apiVersion)) != strings.ToUpper(envVersion) {
		panic(errors.New("API version [" + strings.ToUpper(string(apiVersion)) + "] in responce does not match with client library API version [" + strings.ToUpper(envVersion) + "]"))
	}
	return event
}

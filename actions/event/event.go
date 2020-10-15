package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/event"
)

func List(params *event.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/events"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/events/%v", id), nil)
}
func Content(event event.Event) *chargebee.Result {
	content := &chargebee.Result{}
	err1 := json.Unmarshal(event.Content, content)
	if err1 != nil {
		panic(err1)
	}
	return content
}

func Deserialize(jsonObj string) *event.Event {
	event := &event.Event{}
	err := json.Unmarshal([]byte(jsonObj), event)
	if err != nil {
		panic(errors.New("Response not in JSON format. Might not be a ChargeBee webhook call"))
	}
	apiVersion := event.ApiVersion
	envVersion := chargebee.APIVersion
	if apiVersion != "" && strings.ToUpper(string(apiVersion)) != strings.ToUpper(envVersion) {
		panic(errors.New("API version [" + strings.ToUpper(string(apiVersion)) + "] in responce does not match with client library API version [" + strings.ToUpper(envVersion) + "]"))
	}
	return event
}

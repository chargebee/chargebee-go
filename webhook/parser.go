package webhook

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/chargebee/chargebee-go/v3"
)

// ParseEventType reads only the event_type (and validates api_version) from the webhook payload
func ParseEventType(body []byte) (chargebee.EventType, error) {
	var envelope struct {
		EventType  chargebee.EventType `json:"event_type"`
		ApiVersion string              `json:"api_version"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return "", err
	}
	envVersion := chargebee.APIVersion
	if envelope.ApiVersion != "" && strings.ToUpper(envelope.ApiVersion) != strings.ToUpper(envVersion) {
		return "", errors.New("API version [" + strings.ToUpper(envelope.ApiVersion) + "] in response does not match with client library API version [" + strings.ToUpper(envVersion) + "]")
	}
	return envelope.EventType, nil
}

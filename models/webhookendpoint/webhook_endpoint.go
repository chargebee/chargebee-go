package webhookendpoint

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	webhookEndpointEnum "github.com/chargebee/chargebee-go/v3/models/webhookendpoint/enum"
)

type WebhookEndpoint struct {
	Id                          string                           `json:"id"`
	Name                        string                           `json:"name"`
	Url                         string                           `json:"url"`
	SendCardResource            bool                             `json:"send_card_resource"`
	Disabled                    bool                             `json:"disabled"`
	PrimaryUrl                  bool                             `json:"primary_url"`
	ApiVersion                  webhookEndpointEnum.ApiVersion   `json:"api_version"`
	ChargebeeResponseSchemaType enum.ChargebeeResponseSchemaType `json:"chargebee_response_schema_type"`
	EnabledEvents               []enum.EventType                 `json:"enabled_events"`
	Object                      string                           `json:"object"`
}
type CreateRequestParams struct {
	Name                        string                           `json:"name"`
	ApiVersion                  webhookEndpointEnum.ApiVersion   `json:"api_version,omitempty"`
	Url                         string                           `json:"url"`
	PrimaryUrl                  *bool                            `json:"primary_url,omitempty"`
	Disabled                    *bool                            `json:"disabled,omitempty"`
	BasicAuthPassword           string                           `json:"basic_auth_password,omitempty"`
	BasicAuthUsername           string                           `json:"basic_auth_username,omitempty"`
	SendCardResource            *bool                            `json:"send_card_resource,omitempty"`
	ChargebeeResponseSchemaType enum.ChargebeeResponseSchemaType `json:"chargebee_response_schema_type,omitempty"`
	EnabledEvents               []enum.EventType                 `json:"enabled_events,omitempty"`
}
type UpdateRequestParams struct {
	Name              string                         `json:"name,omitempty"`
	ApiVersion        webhookEndpointEnum.ApiVersion `json:"api_version,omitempty"`
	Url               string                         `json:"url,omitempty"`
	PrimaryUrl        *bool                          `json:"primary_url,omitempty"`
	SendCardResource  *bool                          `json:"send_card_resource,omitempty"`
	BasicAuthPassword string                         `json:"basic_auth_password,omitempty"`
	BasicAuthUsername string                         `json:"basic_auth_username,omitempty"`
	Disabled          *bool                          `json:"disabled,omitempty"`
	EnabledEvents     []enum.EventType               `json:"enabled_events,omitempty"`
}
type ListRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

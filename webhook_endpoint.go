package chargebee

type ApiVersion string

const (
	ApiVersionV1 ApiVersion = "v1"
	ApiVersionV2 ApiVersion = "v2"
)

type WebhookEndpoint struct {
	Id                          string                           `json:"id"`
	Name                        string                           `json:"name"`
	Url                         string                           `json:"url"`
	SendCardResource            bool                             `json:"send_card_resource"`
	Disabled                    bool                             `json:"disabled"`
	PrimaryUrl                  bool                             `json:"primary_url"`
	ApiVersion                  ApiVersion                       `json:"api_version"`
	ChargebeeResponseSchemaType enum.ChargebeeResponseSchemaType `json:"chargebee_response_schema_type"`
	EnabledEvents               []enum.EventType                 `json:"enabled_events"`
	Object                      string                           `json:"object"`
}
type CreateRequest struct {
	Name                        string                           `json:"name"`
	ApiVersion                  ApiVersion                       `json:"api_version,omitempty"`
	Url                         string                           `json:"url"`
	PrimaryUrl                  *bool                            `json:"primary_url,omitempty"`
	Disabled                    *bool                            `json:"disabled,omitempty"`
	BasicAuthPassword           string                           `json:"basic_auth_password,omitempty"`
	BasicAuthUsername           string                           `json:"basic_auth_username,omitempty"`
	SendCardResource            *bool                            `json:"send_card_resource,omitempty"`
	ChargebeeResponseSchemaType enum.ChargebeeResponseSchemaType `json:"chargebee_response_schema_type,omitempty"`
	EnabledEvents               []enum.EventType                 `json:"enabled_events,omitempty"`
}
type UpdateRequest struct {
	Name              string           `json:"name,omitempty"`
	ApiVersion        ApiVersion       `json:"api_version,omitempty"`
	Url               string           `json:"url,omitempty"`
	PrimaryUrl        *bool            `json:"primary_url,omitempty"`
	SendCardResource  *bool            `json:"send_card_resource,omitempty"`
	BasicAuthPassword string           `json:"basic_auth_password,omitempty"`
	BasicAuthUsername string           `json:"basic_auth_username,omitempty"`
	Disabled          *bool            `json:"disabled,omitempty"`
	EnabledEvents     []enum.EventType `json:"enabled_events,omitempty"`
}
type ListRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type CreateResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

type UpdateResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

type RetrieveResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

type DeleteResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

type ListWebhookEndpointResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

type ListResponse struct {
	List       []*ListWebhookEndpointResponse `json:"list,omitempty"`
	NextOffset string                         `json:"next_offset,omitempty"`
}

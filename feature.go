package chargebee

type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
	StatusDraft    Status = "draft"
)

type Type string

const (
	TypeSwitch   Type = "switch"
	TypeCustom   Type = "custom"
	TypeQuantity Type = "quantity"
	TypeRange    Type = "range"
)

type Feature struct {
	Id              string                 `json:"id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Status          Status                 `json:"status"`
	Type            Type                   `json:"type"`
	Unit            string                 `json:"unit"`
	ResourceVersion int64                  `json:"resource_version"`
	UpdatedAt       int64                  `json:"updated_at"`
	CreatedAt       int64                  `json:"created_at"`
	Levels          []*Level               `json:"levels"`
	CustomField     map[string]interface{} `json:"custom_field"`
	Object          string                 `json:"object"`
}
type Level struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Level       int32  `json:"level"`
	IsUnlimited bool   `json:"is_unlimited"`
	Object      string `json:"object"`
}
type ListRequest struct {
	Limit  *int32               `json:"limit,omitempty"`
	Offset string               `json:"offset,omitempty"`
	Name   *filter.StringFilter `json:"name,omitempty"`
	Id     *filter.StringFilter `json:"id,omitempty"`
	Status *filter.EnumFilter   `json:"status,omitempty"`
	Type   *filter.EnumFilter   `json:"type,omitempty"`
}
type CreateRequest struct {
	Id          string         `json:"id,omitempty"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Type        Type           `json:"type,omitempty"`
	Unit        string         `json:"unit,omitempty"`
	Levels      []*CreateLevel `json:"levels,omitempty"`
}
type CreateLevel struct {
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	IsUnlimited *bool  `json:"is_unlimited,omitempty"`
	Level       *int32 `json:"level,omitempty"`
}
type UpdateRequest struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Unit        string         `json:"unit,omitempty"`
	Levels      []*UpdateLevel `json:"levels,omitempty"`
}
type UpdateLevel struct {
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	IsUnlimited *bool  `json:"is_unlimited,omitempty"`
	Level       *int32 `json:"level,omitempty"`
}

type ListFeatureResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type ListResponse struct {
	List       []*ListFeatureResponse `json:"list,omitempty"`
	NextOffset string                 `json:"next_offset,omitempty"`
}

type CreateResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type UpdateResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type RetrieveResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type DeleteResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type ActivateResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type ArchiveResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

type ReactivateResponse struct {
	Feature *Feature `json:"feature,omitempty"`
}

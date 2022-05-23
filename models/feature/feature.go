package feature

import (
	"github.com/chargebee/chargebee-go/filter"
	featureEnum "github.com/chargebee/chargebee-go/models/feature/enum"
)

type Feature struct {
	Id              string             `json:"id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Status          featureEnum.Status `json:"status"`
	Type            featureEnum.Type   `json:"type"`
	Unit            string             `json:"unit"`
	ResourceVersion int64              `json:"resource_version"`
	UpdatedAt       int64              `json:"updated_at"`
	CreatedAt       int64              `json:"created_at"`
	Levels          []*Level           `json:"levels"`
	Object          string             `json:"object"`
}
type Level struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Level       int32  `json:"level"`
	IsUnlimited bool   `json:"is_unlimited"`
	Object      string `json:"object"`
}
type ListRequestParams struct {
	Limit  *int32               `json:"limit,omitempty"`
	Offset string               `json:"offset,omitempty"`
	Name   *filter.StringFilter `json:"name,omitempty"`
	Id     *filter.StringFilter `json:"id,omitempty"`
	Status *filter.EnumFilter   `json:"status,omitempty"`
	Type   *filter.EnumFilter   `json:"type,omitempty"`
}
type CreateRequestParams struct {
	Id          string               `json:"id,omitempty"`
	Name        string               `json:"name"`
	Description string               `json:"description,omitempty"`
	Type        featureEnum.Type     `json:"type,omitempty"`
	Status      featureEnum.Status   `json:"status,omitempty"`
	Unit        string               `json:"unit,omitempty"`
	Levels      []*CreateLevelParams `json:"levels,omitempty"`
}
type CreateLevelParams struct {
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	IsUnlimited *bool  `json:"is_unlimited,omitempty"`
	Level       *int32 `json:"level,omitempty"`
}
type UpdateRequestParams struct {
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Status      featureEnum.Status   `json:"status,omitempty"`
	Unit        string               `json:"unit,omitempty"`
	Levels      []*UpdateLevelParams `json:"levels,omitempty"`
}
type UpdateLevelParams struct {
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	IsUnlimited *bool  `json:"is_unlimited,omitempty"`
	Level       *int32 `json:"level,omitempty"`
}

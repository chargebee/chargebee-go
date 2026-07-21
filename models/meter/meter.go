package meter

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	"github.com/chargebee/chargebee-go/v3/models/columndefinition"
	"github.com/chargebee/chargebee-go/v3/models/feature"
	meterEnum "github.com/chargebee/chargebee-go/v3/models/meter/enum"
)

type Meter struct {
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Description       string                               `json:"description"`
	Type              meterEnum.Type                       `json:"type"`
	Status            meterEnum.Status                     `json:"status"`
	Query             string                               `json:"query"`
	CreatedAt         int64                                `json:"created_at"`
	UpdatedAt         int64                                `json:"updated_at"`
	ColumnDefinitions []*columndefinition.ColumnDefinition `json:"column_definitions"`
	Features          []*feature.Feature                   `json:"features"`
	Object            string                               `json:"object"`
}
type ListRequestParams struct {
	Limit  *int32               `json:"limit,omitempty"`
	Offset string               `json:"offset,omitempty"`
	Name   *filter.StringFilter `json:"name,omitempty"`
	SortBy *filter.SortFilter   `json:"sort_by,omitempty"`
}

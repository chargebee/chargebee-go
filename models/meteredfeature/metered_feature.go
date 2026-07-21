package meteredfeature

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/models/columndefinition"
	"github.com/chargebee/chargebee-go/v3/models/feature"
	meteredFeatureEnum "github.com/chargebee/chargebee-go/v3/models/meteredfeature/enum"
)

type MeteredFeature struct {
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Description       string                               `json:"description"`
	Type              enum.Type                            `json:"type"`
	Status            enum.Status                          `json:"status"`
	Query             string                               `json:"query"`
	ColumnDefinitions []*columndefinition.ColumnDefinition `json:"column_definitions"`
	Features          []*feature.Feature                   `json:"features"`
	Object            string                               `json:"object"`
}
type CreateRequestParams struct {
	Name              string                          `json:"name"`
	Description       string                          `json:"description,omitempty"`
	FeatureUnit       string                          `json:"feature_unit"`
	Query             string                          `json:"query"`
	ColumnDefinitions []*CreateColumnDefinitionParams `json:"column_definitions,omitempty"`
}
type CreateColumnDefinitionParams struct {
	ColumnName string                                      `json:"column_name"`
	DataType   meteredFeatureEnum.ColumnDefinitionDataType `json:"data_type"`
}

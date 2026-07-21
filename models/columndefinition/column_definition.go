package columndefinition

import (
	columnDefinitionEnum "github.com/chargebee/chargebee-go/v3/models/columndefinition/enum"
)

type ColumnDefinition struct {
	ColumnName string                        `json:"column_name"`
	DataType   columnDefinitionEnum.DataType `json:"data_type"`
	Object     string                        `json:"object"`
}

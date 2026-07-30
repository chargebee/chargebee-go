package creditunit

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	creditUnitEnum "github.com/chargebee/chargebee-go/v3/models/creditunit/enum"
)

type CreditUnit struct {
	Id              string                `json:"id"`
	Name            string                `json:"name"`
	ExternalName    string                `json:"external_name"`
	Status          creditUnitEnum.Status `json:"status"`
	ResourceVersion int64                 `json:"resource_version"`
	UpdatedAt       int64                 `json:"updated_at"`
	CreatedAt       int64                 `json:"created_at"`
	CreatedBy       string                `json:"created_by"`
	UpdatedBy       string                `json:"updated_by"`
	IsUnlimited     bool                  `json:"is_unlimited"`
	OverdraftAmount string                `json:"overdraft_amount"`
	Object          string                `json:"object"`
}
type ListRequestParams struct {
	Limit  *int32               `json:"limit,omitempty"`
	Offset string               `json:"offset,omitempty"`
	Status *filter.EnumFilter   `json:"status,omitempty"`
	Id     *filter.StringFilter `json:"id,omitempty"`
}
type CreateRequestParams struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	IsUnlimited     *bool  `json:"is_unlimited"`
	OverdraftAmount string `json:"overdraft_amount,omitempty"`
	ExternalName    string `json:"external_name,omitempty"`
}
type UpdateRequestParams struct {
	Name         string `json:"name,omitempty"`
	ExternalName string `json:"external_name,omitempty"`
}

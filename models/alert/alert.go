package alert

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	alertEnum "github.com/chargebee/chargebee-go/v3/models/alert/enum"
)

type Alert struct {
	Id               string             `json:"id"`
	Type             enum.Type          `json:"type"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	MeteredFeatureId string             `json:"metered_feature_id"`
	CurrencyCode     string             `json:"currency_code"`
	SubscriptionId   string             `json:"subscription_id"`
	Status           alertEnum.Status   `json:"status"`
	Meta             string             `json:"meta"`
	CreatedAt        int64              `json:"created_at"`
	UpdatedAt        int64              `json:"updated_at"`
	Threshold        []*Threshold       `json:"threshold"`
	FilterConditions []*FilterCondition `json:"filter_conditions"`
	Object           string             `json:"object"`
}
type Threshold struct {
	Mode   enum.Mode `json:"mode"`
	Value  float64   `json:"value"`
	Object string    `json:"object"`
}
type FilterCondition struct {
	Field    alertEnum.FilterConditionField    `json:"field"`
	Operator alertEnum.FilterConditionOperator `json:"operator"`
	Value    string                            `json:"value"`
	Object   string                            `json:"object"`
}
type CreateRequestParams struct {
	Type             enum.Type                      `json:"type"`
	Name             string                         `json:"name"`
	Description      string                         `json:"description,omitempty"`
	MeteredFeatureId string                         `json:"metered_feature_id,omitempty"`
	CurrencyCode     string                         `json:"currency_code,omitempty"`
	SubscriptionId   string                         `json:"subscription_id,omitempty"`
	Threshold        *CreateThresholdParams         `json:"threshold,omitempty"`
	Meta             string                         `json:"meta,omitempty"`
	FilterConditions []*CreateFilterConditionParams `json:"filter_conditions,omitempty"`
}
type CreateThresholdParams struct {
	Mode  enum.Mode `json:"mode,omitempty"`
	Value *float64  `json:"value"`
}
type CreateFilterConditionParams struct {
	Field    alertEnum.FilterConditionField    `json:"field,omitempty"`
	Operator alertEnum.FilterConditionOperator `json:"operator,omitempty"`
	Value    string                            `json:"value,omitempty"`
}
type ListRequestParams struct {
	Limit          *int32               `json:"limit,omitempty"`
	Offset         string               `json:"offset,omitempty"`
	Id             *filter.StringFilter `json:"id,omitempty"`
	Type           *filter.EnumFilter   `json:"type,omitempty"`
	SubscriptionId *filter.StringFilter `json:"subscription_id,omitempty"`
	Status         *filter.EnumFilter   `json:"status,omitempty"`
}
type UpdateRequestParams struct {
	Threshold *UpdateThresholdParams `json:"threshold,omitempty"`
	Status    alertEnum.Status       `json:"status,omitempty"`
}
type UpdateThresholdParams struct {
	Mode  enum.Mode `json:"mode,omitempty"`
	Value *float64  `json:"value,omitempty"`
}
type ApplicationAlertsForSubscriptionRequestParams struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	Status *filter.EnumFilter `json:"status,omitempty"`
	Type   *filter.EnumFilter `json:"type,omitempty"`
}

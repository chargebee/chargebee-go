package alertstatus

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
)

type AlertStatus struct {
	AlertId          string           `json:"alert_id"`
	SubscriptionId   string           `json:"subscription_id"`
	AlarmStatus      enum.AlarmStatus `json:"alarm_status"`
	AlarmTriggeredAt int64            `json:"alarm_triggered_at"`
	Object           string           `json:"object"`
}
type AlertStatusesForSubscriptionRequestParams struct {
	Limit       *int32               `json:"limit,omitempty"`
	Offset      string               `json:"offset,omitempty"`
	AlarmStatus *filter.EnumFilter   `json:"alarm_status,omitempty"`
	AlertId     *filter.StringFilter `json:"alert_id,omitempty"`
}
type AlertStatusesForAlertRequestParams struct {
	Limit       *int32             `json:"limit,omitempty"`
	Offset      string             `json:"offset,omitempty"`
	AlarmStatus *filter.EnumFilter `json:"alarm_status,omitempty"`
}

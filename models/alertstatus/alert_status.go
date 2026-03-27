package alertstatus

import (
	"github.com/chargebee/chargebee-go/v3/enum"
)

type AlertStatus struct {
	AlertId          string           `json:"alert_id"`
	SubscriptionId   string           `json:"subscription_id"`
	AlertStatus      enum.AlertStatus `json:"alert_status"`
	AlarmTriggeredAt int64            `json:"alarm_triggered_at"`
	Object           string           `json:"object"`
}

package alert

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	alertEnum "github.com/chargebee/chargebee-go/v3/models/alert/enum"
)

type Alert struct {
	Id               string           `json:"id"`
	Type             enum.Type        `json:"type"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	MeteredFeatureId string           `json:"metered_feature_id"`
	SubscriptionId   string           `json:"subscription_id"`
	Status           alertEnum.Status `json:"status"`
	AlarmTriggeredAt int64            `json:"alarm_triggered_at"`
	Scope            alertEnum.Scope  `json:"scope"`
	Meta             string           `json:"meta"`
	CreatedAt        int64            `json:"created_at"`
	UpdatedAt        int64            `json:"updated_at"`
	Object           string           `json:"object"`
}

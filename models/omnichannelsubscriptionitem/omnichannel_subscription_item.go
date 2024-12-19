package omnichannelsubscriptionitem

import (
	omnichannelSubscriptionItemEnum "github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem/enum"
)

type OmnichannelSubscriptionItem struct {
	Id                   string                                             `json:"id"`
	ItemIdAtSource       string                                             `json:"item_id_at_source"`
	Status               omnichannelSubscriptionItemEnum.Status             `json:"status"`
	CurrentTermStart     int64                                              `json:"current_term_start"`
	CurrentTermEnd       int64                                              `json:"current_term_end"`
	ExpiredAt            int64                                              `json:"expired_at"`
	ExpirationReason     omnichannelSubscriptionItemEnum.ExpirationReason   `json:"expiration_reason"`
	CancelledAt          int64                                              `json:"cancelled_at"`
	CancellationReason   omnichannelSubscriptionItemEnum.CancellationReason `json:"cancellation_reason"`
	GracePeriodExpiresAt int64                                              `json:"grace_period_expires_at"`
	ResourceVersion      int64                                              `json:"resource_version"`
	Object               string                                             `json:"object"`
}

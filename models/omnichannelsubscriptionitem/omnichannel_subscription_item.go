package omnichannelsubscriptionitem

import (
	omnichannelSubscriptionItemEnum "github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem/enum"
)

type OmnichannelSubscriptionItem struct {
	Id                   string                                             `json:"id"`
	ItemIdAtSource       string                                             `json:"item_id_at_source"`
	ItemParentIdAtSource string                                             `json:"item_parent_id_at_source"`
	Status               omnichannelSubscriptionItemEnum.Status             `json:"status"`
	AutoRenewStatus      omnichannelSubscriptionItemEnum.AutoRenewStatus    `json:"auto_renew_status"`
	CurrentTermStart     int64                                              `json:"current_term_start"`
	CurrentTermEnd       int64                                              `json:"current_term_end"`
	ExpiredAt            int64                                              `json:"expired_at"`
	ExpirationReason     omnichannelSubscriptionItemEnum.ExpirationReason   `json:"expiration_reason"`
	CancelledAt          int64                                              `json:"cancelled_at"`
	CancellationReason   omnichannelSubscriptionItemEnum.CancellationReason `json:"cancellation_reason"`
	GracePeriodExpiresAt int64                                              `json:"grace_period_expires_at"`
	HasScheduledChanges  bool                                               `json:"has_scheduled_changes"`
	ResourceVersion      int64                                              `json:"resource_version"`
	UpcomingRenewal      *UpcomingRenewal                                   `json:"upcoming_renewal"`
	Object               string                                             `json:"object"`
}
type UpcomingRenewal struct {
	PriceCurrency string `json:"price_currency"`
	PriceUnits    int64  `json:"price_units"`
	PriceNanos    int64  `json:"price_nanos"`
	Object        string `json:"object"`
}
type ListOmniSubItemScheduleChangesRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

package omnichannelonetimeorderitem

import (
	omnichannelOneTimeOrderItemEnum "github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorderitem/enum"
)

type OmnichannelOneTimeOrderItem struct {
	Id                 string                                             `json:"id"`
	ItemIdAtSource     string                                             `json:"item_id_at_source"`
	ItemTypeAtSource   string                                             `json:"item_type_at_source"`
	Quantity           int32                                              `json:"quantity"`
	CancelledAt        int64                                              `json:"cancelled_at"`
	CancellationReason omnichannelOneTimeOrderItemEnum.CancellationReason `json:"cancellation_reason"`
	CreatedAt          int64                                              `json:"created_at"`
	ResourceVersion    int64                                              `json:"resource_version"`
	Object             string                                             `json:"object"`
}

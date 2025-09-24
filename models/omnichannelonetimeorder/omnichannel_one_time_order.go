package omnichannelonetimeorder

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	omnichannelOneTimeOrderEnum "github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorder/enum"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorderitem"
	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"
)

type OmnichannelOneTimeOrder struct {
	Id                           string                                                     `json:"id"`
	AppId                        string                                                     `json:"app_id"`
	CustomerId                   string                                                     `json:"customer_id"`
	IdAtSource                   string                                                     `json:"id_at_source"`
	Origin                       string                                                     `json:"origin"`
	Source                       omnichannelOneTimeOrderEnum.Source                         `json:"source"`
	CreatedAt                    int64                                                      `json:"created_at"`
	ResourceVersion              int64                                                      `json:"resource_version"`
	OmnichannelOneTimeOrderItems []*omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_items"`
	PurchaseTransaction          *omnichanneltransaction.OmnichannelTransaction             `json:"purchase_transaction"`
	Object                       string                                                     `json:"object"`
}
type ListRequestParams struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	Source     *filter.EnumFilter   `json:"source,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}

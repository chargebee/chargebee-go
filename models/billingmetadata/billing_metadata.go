package billingmetadata

import (
	"github.com/chargebee/chargebee-go/v3/models/billinglineitem"
)

type BillingMetadata struct {
	ItemPriceId string                             `json:"item_price_id"`
	LineItems   []*billinglineitem.BillingLineItem `json:"line_items"`
	Object      string                             `json:"object"`
}

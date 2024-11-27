package omnichanneltransaction

import (
	omnichannelTransactionEnum "github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction/enum"
)

type OmnichannelTransaction struct {
	Id              string                          `json:"id"`
	IdAtSource      string                          `json:"id_at_source"`
	AppId           string                          `json:"app_id"`
	PriceCurrency   string                          `json:"price_currency"`
	PriceUnits      int64                           `json:"price_units"`
	PriceNanos      int64                           `json:"price_nanos"`
	Type            omnichannelTransactionEnum.Type `json:"type"`
	TransactedAt    int64                           `json:"transacted_at"`
	CreatedAt       int64                           `json:"created_at"`
	ResourceVersion int64                           `json:"resource_version"`
	Object          string                          `json:"object"`
}

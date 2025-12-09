package chargebee

type Type string

const (
	TypePurchase Type = "purchase"
	TypeRenewal  Type = "renewal"
)

type OmnichannelTransaction struct {
	Id                             string                           `json:"id"`
	IdAtSource                     string                           `json:"id_at_source"`
	AppId                          string                           `json:"app_id"`
	PriceCurrency                  string                           `json:"price_currency"`
	PriceUnits                     int64                            `json:"price_units"`
	PriceNanos                     int64                            `json:"price_nanos"`
	Type                           Type                             `json:"type"`
	TransactedAt                   int64                            `json:"transacted_at"`
	CreatedAt                      int64                            `json:"created_at"`
	ResourceVersion                int64                            `json:"resource_version"`
	LinkedOmnichannelSubscriptions []*LinkedOmnichannelSubscription `json:"linked_omnichannel_subscriptions"`
	LinkedOmnichannelOneTimeOrders []*LinkedOmnichannelOneTimeOrder `json:"linked_omnichannel_one_time_orders"`
	Object                         string                           `json:"object"`
}
type LinkedOmnichannelSubscription struct {
	OmnichannelSubscriptionId string `json:"omnichannel_subscription_id"`
	Object                    string `json:"object"`
}
type LinkedOmnichannelOneTimeOrder struct {
	OmnichannelOneTimeOrderId string `json:"omnichannel_one_time_order_id"`
	Object                    string `json:"object"`
}

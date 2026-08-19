package omnichannelsubscriptionitemmetric

type OmnichannelSubscriptionItemMetric struct {
	CustomerId                    string `json:"customer_id"`
	OmnichannelSubscriptionId     string `json:"omnichannel_subscription_id"`
	OmnichannelSubscriptionItemId string `json:"omnichannel_subscription_item_id"`
	ItemIdAtSource                string `json:"item_id_at_source"`
	MrrCurrency                   string `json:"mrr_currency"`
	MrrUnits                      int64  `json:"mrr_units"`
	MrrNanos                      int64  `json:"mrr_nanos"`
	EffectiveFrom                 int64  `json:"effective_from"`
	CalculatedAt                  int64  `json:"calculated_at"`
	CreatedAt                     int64  `json:"created_at"`
	ResourceVersion               int64  `json:"resource_version"`
	Object                        string `json:"object"`
}

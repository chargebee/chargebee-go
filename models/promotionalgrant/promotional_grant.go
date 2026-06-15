package promotionalgrant

type PromotionalGrant struct {
	SubscriptionId string `json:"subscription_id"`
	UnitId         string `json:"unit_id"`
	Amount         string `json:"amount"`
	ExpiresAt      int64  `json:"expires_at"`
	Metadata       string `json:"metadata"`
	Object         string `json:"object"`
}
type PromotionalGrantsRequestParams struct {
	SubscriptionId string                 `json:"subscription_id"`
	UnitId         string                 `json:"unit_id"`
	Amount         string                 `json:"amount"`
	ExpiresAt      *int64                 `json:"expires_at"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

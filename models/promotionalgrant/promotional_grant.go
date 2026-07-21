package promotionalgrant

import (
	"encoding/json"
)

type PromotionalGrant struct {
	SubscriptionId string          `json:"subscription_id"`
	UnitId         string          `json:"unit_id"`
	Amount         string          `json:"amount"`
	ExpiresAt      int64           `json:"expires_at"`
	Metadata       json.RawMessage `json:"metadata"`
	Object         string          `json:"object"`
}
type PromotionalGrantsRequestParams struct {
	SubscriptionId string                 `json:"subscription_id"`
	UnitId         string                 `json:"unit_id"`
	Amount         string                 `json:"amount"`
	ExpiresAt      *int64                 `json:"expires_at"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

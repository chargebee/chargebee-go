package chargebee

type SubscriptionEntitlementsCreatedDetail struct {
	SubscriptionId string `json:"subscription_id"`
	HasNext        bool   `json:"has_next"`
	Object         string `json:"object"`
}

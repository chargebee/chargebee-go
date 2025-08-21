package customerentitlement

type CustomerEntitlement struct {
	CustomerId     string `json:"customer_id"`
	SubscriptionId string `json:"subscription_id"`
	FeatureId      string `json:"feature_id"`
	Value          string `json:"value"`
	Name           string `json:"name"`
	IsEnabled      bool   `json:"is_enabled"`
	Object         string `json:"object"`
}
type EntitlementsForCustomerRequestParams struct {
	Limit                   *int32 `json:"limit,omitempty"`
	Offset                  string `json:"offset,omitempty"`
	ConsolidateEntitlements *bool  `json:"consolidate_entitlements,omitempty"`
}

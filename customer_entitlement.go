package chargebee

type CustomerEntitlement struct {
	CustomerId     string `json:"customer_id"`
	SubscriptionId string `json:"subscription_id"`
	FeatureId      string `json:"feature_id"`
	Value          string `json:"value"`
	Name           string `json:"name"`
	IsEnabled      bool   `json:"is_enabled"`
	Object         string `json:"object"`
}
type EntitlementsForCustomerRequest struct {
	Limit                   *int32 `json:"limit,omitempty"`
	Offset                  string `json:"offset,omitempty"`
	ConsolidateEntitlements *bool  `json:"consolidate_entitlements,omitempty"`
}

type EntitlementsForCustomerCustomerEntitlementResponse struct {
	CustomerEntitlement *CustomerEntitlement `json:"customer_entitlement,omitempty"`
}

type EntitlementsForCustomerResponse struct {
	List       []*EntitlementsForCustomerCustomerEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                                `json:"next_offset,omitempty"`
}

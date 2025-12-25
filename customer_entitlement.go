package chargebee

// just struct
type CustomerEntitlement struct {
	CustomerId     string `json:"customer_id"`
	SubscriptionId string `json:"subscription_id"`
	FeatureId      string `json:"feature_id"`
	Value          string `json:"value"`
	Name           string `json:"name"`
	IsEnabled      bool   `json:"is_enabled"`
	Object         string `json:"object"`
}

// sub resources
// operations
// input params
type CustomerEntitlementEntitlementsForCustomerRequest struct {
	Limit                   *int32 `json:"limit,omitempty"`
	Offset                  string `json:"offset,omitempty"`
	ConsolidateEntitlements *bool  `json:"consolidate_entitlements,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *CustomerEntitlementEntitlementsForCustomerRequest) payload() any { return r }

// operation sub response
type CustomerEntitlementEntitlementsForCustomerCustomerEntitlementResponse struct {
	CustomerEntitlement *CustomerEntitlement `json:"customer_entitlement,omitempty"`
}

// operation response
type CustomerEntitlementEntitlementsForCustomerResponse struct {
	List       []*CustomerEntitlementEntitlementsForCustomerCustomerEntitlementResponse `json:"list,omitempty"`
	NextOffset string                                                                   `json:"next_offset,omitempty"`
	apiResponse
}

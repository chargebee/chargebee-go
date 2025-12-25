package chargebee

// just struct
type QuotedDeltaRamp struct {
	LineItems []*QuotedDeltaRampLineItem `json:"line_items"`
	Object    string                     `json:"object"`
}

// sub resources
type QuotedDeltaRampLineItem struct {
	ItemLevelDiscountPerBillingCycleInDecimal string `json:"item_level_discount_per_billing_cycle_in_decimal"`
	Object                                    string `json:"object"`
}

// operations
// input params

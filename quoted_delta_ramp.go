package chargebee

type QuotedDeltaRamp struct {
	LineItems []*LineItem `json:"line_items"`
	Object    string      `json:"object"`
}
type LineItem struct {
	ItemLevelDiscountPerBillingCycleInDecimal string `json:"item_level_discount_per_billing_cycle_in_decimal"`
	Object                                    string `json:"object"`
}

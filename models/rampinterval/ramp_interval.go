package rampinterval

type RampInterval struct {
	StartingBillingCycle int32  `json:"starting_billing_cycle"`
	EndingBillingCycle   int32  `json:"ending_billing_cycle"`
	ApplyCatalogPrice    bool   `json:"apply_catalog_price"`
	Price                int64  `json:"price"`
	PriceInDecimal       string `json:"price_in_decimal"`
	Object               string `json:"object"`
}

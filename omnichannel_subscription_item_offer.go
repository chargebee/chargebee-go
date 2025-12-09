package chargebee

type OmnichannelSubscriptionItemOffer struct {
	Id               string            `json:"id"`
	OfferIdAtSource  string            `json:"offer_id_at_source"`
	Category         enum.Category     `json:"category"`
	CategoryAtSource string            `json:"category_at_source"`
	Type             enum.Type         `json:"type"`
	TypeAtSource     string            `json:"type_at_source"`
	DiscountType     enum.DiscountType `json:"discount_type"`
	Duration         string            `json:"duration"`
	Percentage       float64           `json:"percentage"`
	PriceCurrency    string            `json:"price_currency"`
	PriceUnits       int64             `json:"price_units"`
	PriceNanos       int64             `json:"price_nanos"`
	OfferTermStart   int64             `json:"offer_term_start"`
	OfferTermEnd     int64             `json:"offer_term_end"`
	ResourceVersion  int64             `json:"resource_version"`
	Object           string            `json:"object"`
}

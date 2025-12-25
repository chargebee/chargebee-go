package chargebee

type OmnichannelSubscriptionItemOfferCategory string

const (
	OmnichannelSubscriptionItemOfferCategoryIntroductory        OmnichannelSubscriptionItemOfferCategory = "introductory"
	OmnichannelSubscriptionItemOfferCategoryPromotional         OmnichannelSubscriptionItemOfferCategory = "promotional"
	OmnichannelSubscriptionItemOfferCategoryDeveloperDetermined OmnichannelSubscriptionItemOfferCategory = "developer_determined"
)

type OmnichannelSubscriptionItemOfferType string

const (
	OmnichannelSubscriptionItemOfferTypeFreeTrial  OmnichannelSubscriptionItemOfferType = "free_trial"
	OmnichannelSubscriptionItemOfferTypePayUpFront OmnichannelSubscriptionItemOfferType = "pay_up_front"
	OmnichannelSubscriptionItemOfferTypePayAsYouGo OmnichannelSubscriptionItemOfferType = "pay_as_you_go"
)

type OmnichannelSubscriptionItemOfferDiscountType string

const (
	OmnichannelSubscriptionItemOfferDiscountTypeFixedAmount OmnichannelSubscriptionItemOfferDiscountType = "fixed_amount"
	OmnichannelSubscriptionItemOfferDiscountTypePercentage  OmnichannelSubscriptionItemOfferDiscountType = "percentage"
	OmnichannelSubscriptionItemOfferDiscountTypePrice       OmnichannelSubscriptionItemOfferDiscountType = "price"
)

// just struct
type OmnichannelSubscriptionItemOffer struct {
	Id               string                                       `json:"id"`
	OfferIdAtSource  string                                       `json:"offer_id_at_source"`
	Category         OmnichannelSubscriptionItemOfferCategory     `json:"category"`
	CategoryAtSource string                                       `json:"category_at_source"`
	Type             OmnichannelSubscriptionItemOfferType         `json:"type"`
	TypeAtSource     string                                       `json:"type_at_source"`
	DiscountType     OmnichannelSubscriptionItemOfferDiscountType `json:"discount_type"`
	Duration         string                                       `json:"duration"`
	Percentage       float64                                      `json:"percentage"`
	PriceCurrency    string                                       `json:"price_currency"`
	PriceUnits       int64                                        `json:"price_units"`
	PriceNanos       int64                                        `json:"price_nanos"`
	OfferTermStart   int64                                        `json:"offer_term_start"`
	OfferTermEnd     int64                                        `json:"offer_term_end"`
	ResourceVersion  int64                                        `json:"resource_version"`
	Object           string                                       `json:"object"`
}

// sub resources
// operations
// input params

package enum

type PricingModel string

const (
	PricingModelFlatFee   PricingModel = "flat_fee"
	PricingModelPerUnit   PricingModel = "per_unit"
	PricingModelTiered    PricingModel = "tiered"
	PricingModelVolume    PricingModel = "volume"
	PricingModelStairstep PricingModel = "stairstep"
)

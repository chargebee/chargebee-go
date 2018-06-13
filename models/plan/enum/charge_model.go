package enum

type ChargeModel string

const (
	ChargeModelFlatFee ChargeModel = "flat_fee"
	ChargeModelPerUnit ChargeModel = "per_unit"
)

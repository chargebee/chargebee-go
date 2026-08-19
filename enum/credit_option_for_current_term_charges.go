package enum

type CreditOptionForCurrentTermCharges string

const (
	CreditOptionForCurrentTermChargesNone             CreditOptionForCurrentTermCharges = "none"
	CreditOptionForCurrentTermChargesProrate          CreditOptionForCurrentTermCharges = "prorate"
	CreditOptionForCurrentTermChargesFull             CreditOptionForCurrentTermCharges = "full"
	CreditOptionForCurrentTermChargesConsumptionBased CreditOptionForCurrentTermCharges = "consumption_based"
)

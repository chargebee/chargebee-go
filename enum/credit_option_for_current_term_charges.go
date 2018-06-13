package enum

type CreditOptionForCurrentTermCharges string

const (
	CreditOptionForCurrentTermChargesNone    CreditOptionForCurrentTermCharges = "none"
	CreditOptionForCurrentTermChargesProrate CreditOptionForCurrentTermCharges = "prorate"
	CreditOptionForCurrentTermChargesFull    CreditOptionForCurrentTermCharges = "full"
)

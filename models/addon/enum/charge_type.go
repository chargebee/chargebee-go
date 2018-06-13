package enum

type ChargeType string

const (
	ChargeTypeRecurring    ChargeType = "recurring"
	ChargeTypeNonRecurring ChargeType = "non_recurring"
)

package enum

type ChargeOn string

const (
	ChargeOnImmediately ChargeOn = "immediately"
	ChargeOnOnEvent     ChargeOn = "on_event"
)

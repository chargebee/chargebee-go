package enum

type ChargeOnOption string

const (
	ChargeOnOptionImmediately ChargeOnOption = "immediately"
	ChargeOnOptionOnEvent     ChargeOnOption = "on_event"
)

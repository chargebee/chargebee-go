package enum

type ActionType string

const (
	ActionTypeUpsert ActionType = "upsert"
	ActionTypeRemove ActionType = "remove"
)

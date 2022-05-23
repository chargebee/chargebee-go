package enum

type Action string

const (
	ActionUpsert Action = "upsert"
	ActionRemove Action = "remove"
)

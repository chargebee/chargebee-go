package enum

type Type string

const (
	TypeSwitch   Type = "switch"
	TypeCustom   Type = "custom"
	TypeQuantity Type = "quantity"
	TypeRange    Type = "range"
)

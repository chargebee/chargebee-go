package enum

type OrderType string

const (
	OrderTypeManual          OrderType = "manual"
	OrderTypeSystemGenerated OrderType = "system_generated"
)

package enum

type LinkedOrderOrderType string

const (
	LinkedOrderOrderTypeManual          LinkedOrderOrderType = "manual"
	LinkedOrderOrderTypeSystemGenerated LinkedOrderOrderType = "system_generated"
)

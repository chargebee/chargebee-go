package enum

type ExcludeTaxType string

const (
	ExcludeTaxTypeExclusive ExcludeTaxType = "exclusive"
	ExcludeTaxTypeNone      ExcludeTaxType = "none"
)

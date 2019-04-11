package enum

type AvalaraSaleType string

const (
	AvalaraSaleTypeWholesale AvalaraSaleType = "wholesale"
	AvalaraSaleTypeRetail    AvalaraSaleType = "retail"
	AvalaraSaleTypeConsumed  AvalaraSaleType = "consumed"
	AvalaraSaleTypeVendorUse AvalaraSaleType = "vendor_use"
)

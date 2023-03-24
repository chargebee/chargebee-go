package enum

type EinvoicingMethod string

const (
	EinvoicingMethodAutomatic   EinvoicingMethod = "automatic"
	EinvoicingMethodManual      EinvoicingMethod = "manual"
	EinvoicingMethodSiteDefault EinvoicingMethod = "site_default"
)

package enum

type PreferredScheme string

const (
	PreferredSchemeCartesBancaires PreferredScheme = "cartes_bancaires"
	PreferredSchemeMastercard      PreferredScheme = "mastercard"
	PreferredSchemeVisa            PreferredScheme = "visa"
)

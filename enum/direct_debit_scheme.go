package enum

type DirectDebitScheme string

const (
	DirectDebitSchemeAch           DirectDebitScheme = "ach"
	DirectDebitSchemeBacs          DirectDebitScheme = "bacs"
	DirectDebitSchemeSepaCore      DirectDebitScheme = "sepa_core"
	DirectDebitSchemeAutogiro      DirectDebitScheme = "autogiro"
	DirectDebitSchemeBecs          DirectDebitScheme = "becs"
	DirectDebitSchemeBecsNz        DirectDebitScheme = "becs_nz"
	DirectDebitSchemePad           DirectDebitScheme = "pad"
	DirectDebitSchemeNotApplicable DirectDebitScheme = "not_applicable"
)

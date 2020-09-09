package enum

type PoweredBy string

const (
	PoweredByIdeal         PoweredBy = "ideal"
	PoweredBySofort        PoweredBy = "sofort"
	PoweredByBancontact    PoweredBy = "bancontact"
	PoweredByGiropay       PoweredBy = "giropay"
	PoweredByNotApplicable PoweredBy = "not_applicable"
)

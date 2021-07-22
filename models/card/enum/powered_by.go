package enum

type PoweredBy string

const (
	PoweredByIdeal          PoweredBy = "ideal"
	PoweredBySofort         PoweredBy = "sofort"
	PoweredByBancontact     PoweredBy = "bancontact"
	PoweredByGiropay        PoweredBy = "giropay"
	PoweredByCard           PoweredBy = "card"
	PoweredByLatamLocalCard PoweredBy = "latam_local_card"
	PoweredByNotApplicable  PoweredBy = "not_applicable"
)

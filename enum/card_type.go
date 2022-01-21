package enum

type CardType string

const (
	CardTypeVisa            CardType = "visa"
	CardTypeMastercard      CardType = "mastercard"
	CardTypeAmericanExpress CardType = "american_express"
	CardTypeDiscover        CardType = "discover"
	CardTypeJcb             CardType = "jcb"
	CardTypeDinersClub      CardType = "diners_club"
	CardTypeBancontact      CardType = "bancontact"
	CardTypeOther           CardType = "other"
	CardTypeNotApplicable   CardType = "not_applicable"
)

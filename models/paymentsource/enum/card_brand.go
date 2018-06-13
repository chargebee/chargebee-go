package enum

type CardBrand string

const (
	CardBrandVisa            CardBrand = "visa"
	CardBrandMastercard      CardBrand = "mastercard"
	CardBrandAmericanExpress CardBrand = "american_express"
	CardBrandDiscover        CardBrand = "discover"
	CardBrandJcb             CardBrand = "jcb"
	CardBrandDinersClub      CardBrand = "diners_club"
	CardBrandOther           CardBrand = "other"
)

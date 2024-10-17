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
	CardTypeCmrFalabella    CardType = "cmr_falabella"
	CardTypeTarjetaNaranja  CardType = "tarjeta_naranja"
	CardTypeNativa          CardType = "nativa"
	CardTypeCencosud        CardType = "cencosud"
	CardTypeCabal           CardType = "cabal"
	CardTypeArgencard       CardType = "argencard"
	CardTypeElo             CardType = "elo"
	CardTypeHipercard       CardType = "hipercard"
	CardTypeCarnet          CardType = "carnet"
	CardTypeRupay           CardType = "rupay"
	CardTypeMaestro         CardType = "maestro"
	CardTypeDankort         CardType = "dankort"
	CardTypeCartesBancaires CardType = "cartes_bancaires"
	CardTypeOther           CardType = "other"
	CardTypeNotApplicable   CardType = "not_applicable"
)

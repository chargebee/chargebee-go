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
	CardBrandBancontact      CardBrand = "bancontact"
	CardBrandCmrFalabella    CardBrand = "cmr_falabella"
	CardBrandTarjetaNaranja  CardBrand = "tarjeta_naranja"
	CardBrandNativa          CardBrand = "nativa"
	CardBrandCencosud        CardBrand = "cencosud"
	CardBrandCabal           CardBrand = "cabal"
	CardBrandArgencard       CardBrand = "argencard"
	CardBrandElo             CardBrand = "elo"
	CardBrandHipercard       CardBrand = "hipercard"
	CardBrandCarnet          CardBrand = "carnet"
	CardBrandRupay           CardBrand = "rupay"
	CardBrandMaestro         CardBrand = "maestro"
	CardBrandDankort         CardBrand = "dankort"
	CardBrandCartesBancaires CardBrand = "cartes_bancaires"
	CardBrandNotApplicable   CardBrand = "not_applicable"
)

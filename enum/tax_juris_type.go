package enum

type TaxJurisType string

const (
	TaxJurisTypeCountry TaxJurisType = "country"
	TaxJurisTypeState   TaxJurisType = "state"
	TaxJurisTypeCounty  TaxJurisType = "county"
	TaxJurisTypeCity    TaxJurisType = "city"
	TaxJurisTypeSpecial TaxJurisType = "special"
	TaxJurisTypeOther   TaxJurisType = "other"
)

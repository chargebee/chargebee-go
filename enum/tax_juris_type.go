package enum

type TaxJurisType string

const (
	TaxJurisTypeCountry        TaxJurisType = "country"
	TaxJurisTypeFederal        TaxJurisType = "federal"
	TaxJurisTypeState          TaxJurisType = "state"
	TaxJurisTypeCounty         TaxJurisType = "county"
	TaxJurisTypeCity           TaxJurisType = "city"
	TaxJurisTypeSpecial        TaxJurisType = "special"
	TaxJurisTypeUnincorporated TaxJurisType = "unincorporated"
	TaxJurisTypeOther          TaxJurisType = "other"
)

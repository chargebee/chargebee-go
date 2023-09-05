package enum

type Tax3JurisType string

const (
	Tax3JurisTypeCountry        Tax3JurisType = "country"
	Tax3JurisTypeFederal        Tax3JurisType = "federal"
	Tax3JurisTypeState          Tax3JurisType = "state"
	Tax3JurisTypeCounty         Tax3JurisType = "county"
	Tax3JurisTypeCity           Tax3JurisType = "city"
	Tax3JurisTypeSpecial        Tax3JurisType = "special"
	Tax3JurisTypeUnincorporated Tax3JurisType = "unincorporated"
	Tax3JurisTypeOther          Tax3JurisType = "other"
)

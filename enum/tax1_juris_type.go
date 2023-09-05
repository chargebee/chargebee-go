package enum

type Tax1JurisType string

const (
	Tax1JurisTypeCountry        Tax1JurisType = "country"
	Tax1JurisTypeFederal        Tax1JurisType = "federal"
	Tax1JurisTypeState          Tax1JurisType = "state"
	Tax1JurisTypeCounty         Tax1JurisType = "county"
	Tax1JurisTypeCity           Tax1JurisType = "city"
	Tax1JurisTypeSpecial        Tax1JurisType = "special"
	Tax1JurisTypeUnincorporated Tax1JurisType = "unincorporated"
	Tax1JurisTypeOther          Tax1JurisType = "other"
)

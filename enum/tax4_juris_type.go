package enum

type Tax4JurisType string

const (
	Tax4JurisTypeCountry        Tax4JurisType = "country"
	Tax4JurisTypeFederal        Tax4JurisType = "federal"
	Tax4JurisTypeState          Tax4JurisType = "state"
	Tax4JurisTypeCounty         Tax4JurisType = "county"
	Tax4JurisTypeCity           Tax4JurisType = "city"
	Tax4JurisTypeSpecial        Tax4JurisType = "special"
	Tax4JurisTypeUnincorporated Tax4JurisType = "unincorporated"
	Tax4JurisTypeOther          Tax4JurisType = "other"
)

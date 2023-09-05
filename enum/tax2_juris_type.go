package enum

type Tax2JurisType string

const (
	Tax2JurisTypeCountry        Tax2JurisType = "country"
	Tax2JurisTypeFederal        Tax2JurisType = "federal"
	Tax2JurisTypeState          Tax2JurisType = "state"
	Tax2JurisTypeCounty         Tax2JurisType = "county"
	Tax2JurisTypeCity           Tax2JurisType = "city"
	Tax2JurisTypeSpecial        Tax2JurisType = "special"
	Tax2JurisTypeUnincorporated Tax2JurisType = "unincorporated"
	Tax2JurisTypeOther          Tax2JurisType = "other"
)

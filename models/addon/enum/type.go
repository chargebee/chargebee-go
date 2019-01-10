package enum

type Type string

const (
	TypeOnOff     Type = "on_off"
	TypeQuantity  Type = "quantity"
	TypeTiered    Type = "tiered"
	TypeVolume    Type = "volume"
	TypeStairstep Type = "stairstep"
)

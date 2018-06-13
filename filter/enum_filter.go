package filter

type EnumFilter struct {
	Is        interface{}   `json:"is,omitempty"`
	IsNot     interface{}   `json:"is_not,omitempty"`
	IsPresent *bool         `json:"is_present,omitempty"`
	In        []interface{} `json:"in,omitempty"`
	NotIn     []interface{} `json:"not_in,omitempty"`
}

package filter

type BooleanFilter struct {
	Is        *bool `json:"is,omitempty"`
	IsPresent *bool `json:"is_present,omitempty"`
}

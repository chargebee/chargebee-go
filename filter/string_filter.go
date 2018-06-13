package filter

type StringFilter struct {
	Is         string   `json:"is,omitempty"`
	IsNot      string   `json:"is_not,omitempty"`
	StartsWith string   `json:"starts_with,omitempty"`
	In         []string `json:"in,omitempty"`
	NotIn      []string `json:"not_in,omitempty"`
	IsPresent  *bool    `json:"is_present,omitempty"`
}

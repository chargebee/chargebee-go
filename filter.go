package chargebee

type BooleanFilter struct {
	Is        *bool `json:"is,omitempty"`
	IsPresent *bool `json:"is_present,omitempty"`
}

type EnumFilter struct {
	Is        interface{}   `json:"is,omitempty"`
	IsNot     interface{}   `json:"is_not,omitempty"`
	IsPresent *bool         `json:"is_present,omitempty"`
	In        []interface{} `json:"in,omitempty"`
	NotIn     []interface{} `json:"not_in,omitempty"`
}

type NumberFilter struct {
	Is        int64   `json:"is,omitempty"`
	IsNot     int64   `json:"is_not,omitempty"`
	Lt        int64   `json:"lt,omitempty"`
	Lte       int64   `json:"lte,omitempty"`
	Gt        int64   `json:"gt,omitempty"`
	Gte       int64   `json:"gte,omitempty"`
	Between   []int64 `json:"between,omitempty"`
	IsPresent *bool   `json:"is_present,omitempty"`
}

type SortFilter struct {
	Asc  string `json:"asc,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type StringFilter struct {
	Is         string   `json:"is,omitempty"`
	IsNot      string   `json:"is_not,omitempty"`
	StartsWith string   `json:"starts_with,omitempty"`
	In         []string `json:"in,omitempty"`
	NotIn      []string `json:"not_in,omitempty"`
	IsPresent  *bool    `json:"is_present,omitempty"`
}

type TimestampFilter struct {
	After     int64   `json:"after,omitempty"`
	Before    int64   `json:"before,omitempty"`
	On        int64   `json:"on,omitempty"`
	Between   []int64 `json:"between,omitempty"`
	IsPresent *bool   `json:"is_present,omitempty"`
}

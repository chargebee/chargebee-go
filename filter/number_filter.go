package filter

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

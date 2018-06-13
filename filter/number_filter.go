package filter

type NumberFilter struct {
	Is        int32   `json:"is,omitempty"`
	IsNot     int32   `json:"is_not,omitempty"`
	Lt        int32   `json:"lt,omitempty"`
	Lte       int32   `json:"lte,omitempty"`
	Gt        int32   `json:"gt,omitempty"`
	Gte       int32   `json:"gte,omitempty"`
	Between   []int32 `json:"between,omitempty"`
	IsPresent *bool   `json:"is_present,omitempty"`
}

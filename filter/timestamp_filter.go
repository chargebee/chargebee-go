package filter

type TimestampFilter struct {
	After     int64   `json:"after,omitempty"`
	Before    int64   `json:"before,omitempty"`
	On        int64   `json:"on,omitempty"`
	Between   []int64 `json:"between,omitempty"`
	IsPresent *bool   `json:"is_present,omitempty"`
}

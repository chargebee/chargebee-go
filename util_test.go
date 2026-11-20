package chargebee

import (
	"testing"

	"github.com/chargebee/chargebee-go/v3/filter"
)

type topLevelListParams struct {
	Limit          *int32 `json:"limit,omitempty"`
	Offset         string `json:"offset,omitempty"`
	IncludeDeleted *bool  `json:"include_deleted,omitempty"`
}

type oneLevelListParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Status    *filter.EnumFilter      `json:"status,omitempty"`
	Total     *filter.NumberFilter    `json:"total,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
	Deleted   *filter.BooleanFilter   `json:"deleted,omitempty"`
	SortBy    *filter.SortFilter      `json:"sort_by,omitempty"`
}

type nestedListInvoiceParams struct {
	DunningStatus *filter.EnumFilter   `json:"dunning_status,omitempty"`
	Id            *filter.StringFilter `json:"id,omitempty"`
}

type nestedListRequestParams struct {
	Limit          *int32                   `json:"limit,omitempty"`
	Invoice        *nestedListInvoiceParams `json:"invoice,omitempty"`
	RetrySegmentId *filter.StringFilter     `json:"retry_segment_id,omitempty"`
}

func TestSerializeListParams(t *testing.T) {
	tests := []struct {
		name   string
		params interface{}
		want   map[string]string
	}{
		{
			name: "top-level scalars only",
			params: &topLevelListParams{
				Limit:          Int32(10),
				Offset:         "abc",
				IncludeDeleted: Bool(true),
			},
			want: map[string]string{
				"limit":           "10",
				"offset":          "abc",
				"include_deleted": "true",
			},
		},
		{
			name: "one-level filters of every type",
			params: &oneLevelListParams{
				Limit:     Int32(5),
				Id:        &filter.StringFilter{StartsWith: "cus_"},
				Status:    &filter.EnumFilter{Is: "active"},
				Total:     &filter.NumberFilter{Gte: 100},
				UpdatedAt: &filter.TimestampFilter{After: 1700000000},
				Deleted:   &filter.BooleanFilter{Is: Bool(false)},
				SortBy:    &filter.SortFilter{Asc: "created_at"},
			},
			want: map[string]string{
				"limit":             "5",
				"id[starts_with]":   "cus_",
				"status[is]":        "active",
				"total[gte]":        "100",
				"updated_at[after]": "1700000000",
				"deleted[is]":       "false",
				"sort_by[asc]":      "created_at",
			},
		},
		{
			name: "one-level filters with array values",
			params: &oneLevelListParams{
				Id:        &filter.StringFilter{In: []string{"a", "b"}},
				Status:    &filter.EnumFilter{NotIn: []interface{}{"cancelled", "paused"}},
				UpdatedAt: &filter.TimestampFilter{Between: []int64{1, 2}},
			},
			want: map[string]string{
				"id[in]":              `["a","b"]`,
				"status[not_in]":      `["cancelled","paused"]`,
				"updated_at[between]": `["1","2"]`,
			},
		},
		{
			name: "one-level filter with multiple operators",
			params: &oneLevelListParams{
				Total: &filter.NumberFilter{Gt: 10, Lt: 20},
			},
			want: map[string]string{
				"total[gt]": "10",
				"total[lt]": "20",
			},
		},
		{
			name: "two-level nested filter",
			params: &nestedListRequestParams{
				Invoice: &nestedListInvoiceParams{
					DunningStatus: &filter.EnumFilter{Is: "in_progress"},
				},
			},
			want: map[string]string{
				"invoice[dunning_status][is]": "in_progress",
			},
		},
		{
			name: "nested and one-level filters together",
			params: &nestedListRequestParams{
				Limit: Int32(1),
				Invoice: &nestedListInvoiceParams{
					DunningStatus: &filter.EnumFilter{In: []interface{}{"in_progress", "exhausted"}},
					Id:            &filter.StringFilter{Is: "inv_1"},
				},
				RetrySegmentId: &filter.StringFilter{Is: "seg_1"},
			},
			want: map[string]string{
				"limit":                       "1",
				"invoice[dunning_status][in]": `["in_progress","exhausted"]`,
				"invoice[id][is]":             "inv_1",
				"retry_segment_id[is]":        "seg_1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SerializeListParams(tt.params)
			if len(*got) != len(tt.want) {
				t.Fatalf("expected %d params, got %d: %v", len(tt.want), len(*got), got.Encode())
			}
			for k, v := range tt.want {
				if got.Get(k) != v {
					t.Errorf("param %q: expected %q, got %q (all: %v)", k, v, got.Get(k), got.Encode())
				}
			}
		})
	}
}

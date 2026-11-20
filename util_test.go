package chargebee

import (
	"testing"

	"github.com/chargebee/chargebee-go/v3/filter"
)

type nestedListInvoiceParams struct {
	DunningStatus *filter.EnumFilter `json:"dunning_status,omitempty"`
}

type nestedListRequestParams struct {
	Limit          *int32                   `json:"limit,omitempty"`
	Invoice        *nestedListInvoiceParams `json:"invoice,omitempty"`
	RetrySegmentId *filter.StringFilter     `json:"retry_segment_id,omitempty"`
}

func TestSerializeListParamsNested(t *testing.T) {
	params := &nestedListRequestParams{
		Limit: Int32(1),
		Invoice: &nestedListInvoiceParams{
			DunningStatus: &filter.EnumFilter{Is: "in_progress"},
		},
		RetrySegmentId: &filter.StringFilter{Is: "seg_1"},
	}

	got := SerializeListParams(params)

	want := map[string]string{
		"limit":                       "1",
		"invoice[dunning_status][is]": "in_progress",
		"retry_segment_id[is]":        "seg_1",
	}
	if len(*got) != len(want) {
		t.Fatalf("expected %d params, got %d: %v", len(want), len(*got), got.Encode())
	}
	for k, v := range want {
		if got.Get(k) != v {
			t.Errorf("param %q: expected %q, got %q (all: %v)", k, v, got.Get(k), got.Encode())
		}
	}
}

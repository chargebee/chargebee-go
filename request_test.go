package chargebee

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareJsonRequest(t *testing.T) {
	type testRequest struct {
		ID         string `json:"id"`
		apiRequest `json:"-"`
	}

	req := testRequest{
		ID: "123",
	}
	req.method = "POST"
	req.isJsonRequest = true

	err := req.prepare(req)
	assert.NoError(t, err)
	assert.Nil(t, req.urlParams)
	assert.Equal(t, `{"id":"123"}`, req.jsonBody)
}

func TestPrepareListRequest(t *testing.T) {
	type testRequest struct {
		Limit      *int32        `json:"limit,omitempty"`
		Offset     string        `json:"offset,omitempty"`
		Name       *StringFilter `json:"name,omitempty"`
		apiRequest `json:"-"`
	}

	req := &testRequest{
		Limit:  Int32(10),
		Offset: "100",
		Name: &StringFilter{
			StartsWith: "test",
		},
	}
	req.method = "GET"
	req.isListRequest = true

	err := req.prepare(req)
	assert.NoError(t, err)
	assert.Equal(t, &url.Values{
		"limit":             {"10"},
		"name[starts_with]": {"test"},
		"offset":            {"100"},
	}, req.urlParams)
}

func TestPrepareRequest(t *testing.T) {
	type testRequest struct {
		Limit      *int32        `json:"limit,omitempty"`
		Offset     string        `json:"offset,omitempty"`
		Name       *StringFilter `json:"name,omitempty"`
		apiRequest `json:"-"`
	}

	req := &testRequest{
		Limit:  Int32(10),
		Offset: "100",
		Name: &StringFilter{
			StartsWith: "test",
		},
	}
	req.method = "GET"

	err := req.prepare(req)
	assert.NoError(t, err)
	assert.Equal(t, &url.Values{
		"limit":             {"10"},
		"name[starts_with]": {"test"},
		"offset":            {"100"},
	}, req.urlParams)
}

func TestPrepareNilRequest(t *testing.T) {
	type testRequest struct {
		Limit      *int32        `json:"limit,omitempty"`
		Offset     string        `json:"offset,omitempty"`
		Name       *StringFilter `json:"name,omitempty"`
		apiRequest `json:"-"`
	}

	req := &testRequest{}
	req.method = "GET"

	err := req.prepare(nil)
	assert.NoError(t, err)
	assert.Nil(t, req.urlParams)
	assert.Equal(t, "", req.jsonBody)
}

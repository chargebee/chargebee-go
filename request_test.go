package chargebee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRequestWithPayload struct {
	ID                                                string `json:"id"`
	apiRequest[testRequestWithPayload, *testResponse] `json:"-"`
}

type testResponse struct {
	ID          string
	Name        string
	Description string
	apiResponse `json:"-"`
}

func (r testRequestWithPayload) payload() testRequestWithPayload { return r }

func TestPrepareRequestParams(t *testing.T) {
	req := testRequestWithPayload{
		ID: "123",
	}
	req.isJsonRequest = true
	req.method = "POST"
	req.Bind(req)
	t.Logf("req: %+v", req.payload())

	err := req.prepare()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, req.requestParams)
	assert.Equal(t, `{"id":"123"}`, req.jsonBody)
}

func TestPrepareBlankRequestParams(t *testing.T) {
	req := NewBlankRequest[*testResponse]()

	err := req.prepare()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, req.requestParams)
}

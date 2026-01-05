package chargebee

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiResponse_IsIdempotencyReplayed(t *testing.T) {
	response := &apiResponse{
		Headers: http.Header{
			IdempotencyReplayHeader: {"true"},
		},
	}
	assert.True(t, response.IsIdempotencyReplayed())
}

func TestApiResponse_IsIdempotencyReplayed_False(t *testing.T) {
	response := &apiResponse{
		Headers: http.Header{},
	}
	assert.False(t, response.IsIdempotencyReplayed())
}

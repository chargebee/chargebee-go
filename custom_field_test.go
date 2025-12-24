package chargebee

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomFieldExtractionFromResponse(t *testing.T) {
	data := []byte(`{
		"subscription": {
			"id": "123",
			"custom_field": {
				"cf_gender": "Female",
				"foo": "bar"
			}
		},
		"customer": {
			"id": "123",
			"custom_field": {
				"cf_DOB": "12-10-1994"
			}
		}
	}`)
	var response struct {
		Subscription struct {
			CustomField CustomField `json:"custom_field"`
		} `json:"subscription"`
		Customer struct {
			CustomField CustomField `json:"custom_field"`
		} `json:"customer"`
	}
	err := json.Unmarshal(data, &response)
	assert.NoError(t, err)
	assert.Equal(t, "Female", response.Subscription.CustomField["cf_gender"])
	assert.Equal(t, "12-10-1994", response.Customer.CustomField["cf_DOB"])
}

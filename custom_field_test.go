package chargebee

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObjectWithCustomField struct {
	Id           string        `json:"id"`
	CustomFields *customFields `json:"-"`
}

func (t *testObjectWithCustomField) setCustomFields(cf *customFields) {
	t.CustomFields = cf
}

func TestUnmarshalWithCustomField(t *testing.T) {
	data := []byte(`{
		"id": "123",
		"cf_age": "25",
		"nested": {
			"cf_app": "google"
		}
	}`)

	var obj = new(testObjectWithCustomField)
	err := unmarshalObjectWithCustomField(data, obj, obj)
	fmt.Printf("decoded:+%v\n", obj)
	assert.NoError(t, err)
	assert.Equal(t, "123", obj.Id)
	assert.Equal(t, "25", obj.CustomFields.GetCustomField("cf_age"))
	assert.Equal(t, "google", obj.CustomFields.GetCustomField("app"))
}

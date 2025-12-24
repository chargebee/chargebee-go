package chargebee

import (
	"encoding/json"
	"strings"
)

type CustomField map[string]interface{}

func (c *CustomField) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if *c == nil {
		*c = make(CustomField)
	}
	for k, v := range m {
		if strings.HasPrefix(k, "cf_") {
			(*c)[k] = v
		}
	}

	return nil
}

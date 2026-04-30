package chargebee

import (
	"bytes"
	"encoding/json"
	"strings"
)

type customFields map[string]any

type hasCustomField interface {
	setCustomFields(*customFields)
}

func (c *customFields) set(key string, value any) {
	if *c == nil {
		*c = make(customFields)
	}
	(*c)[key] = value
}

func (c *customFields) Get(key string) any {
	if !strings.HasPrefix(key, "cf_") {
		key = "cf_" + key
	}
	if value, ok := (*c)[key]; ok {
		return value.(string)
	}
	return ""
}

func unmarshalObjectWithCustomField[T hasCustomField](data []byte, out T, alias any) error {
	if err := json.Unmarshal(data, alias); err != nil {
		return err
	}
	cf := new(customFields)

	var tree any
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&tree); err == nil {
		extractCustomFields(tree, cf)
	}

	out.setCustomFields(cf)
	return nil
}

func extractCustomFields(node any, cf *customFields) {
	switch x := node.(type) {
	case map[string]any:
		for k, v := range x {
			if strings.HasPrefix(k, "cf_") {
				cf.set(k, v)
			}
			extractCustomFields(v, cf)
		}
	case []any:
		for _, v := range x {
			extractCustomFields(v, cf)
		}
	}
}

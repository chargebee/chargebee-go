package chargebee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type CustomFields map[string]string

type hasCustomField interface {
	setCustomFields(*CustomFields)
}

func (c *CustomFields) set(key string, value string) {
	if *c == nil {
		*c = make(CustomFields)
	}
	(*c)[key] = value
}

// Get returns the value of the custom field as a string.
// Custom fields are stored and returned as strings, so
// any conversion to other types should be done by the caller.
func (c *CustomFields) Get(key string) string {
	if !strings.HasPrefix(key, "cf_") {
		key = "cf_" + key
	}
	if value, ok := (*c)[key]; ok {
		return value
	}
	return ""
}

func unmarshalObjectWithCustomField[T hasCustomField](data []byte, out T, alias any) error {
	if err := json.Unmarshal(data, alias); err != nil {
		return err
	}
	cf := new(CustomFields)

	var tree any
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&tree); err == nil {
		extractCustomFields(tree, cf)
	}

	out.setCustomFields(cf)
	return nil
}

func extractCustomFields(node any, cf *CustomFields) {
	switch x := node.(type) {
	case map[string]any:
		for k, v := range x {
			if strings.HasPrefix(k, "cf_") {
				cf.set(k, fmt.Sprintf("%v", v))
			}
			extractCustomFields(v, cf)
		}
	case []any:
		for _, v := range x {
			extractCustomFields(v, cf)
		}
	}
}

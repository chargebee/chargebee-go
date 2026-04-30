package chargebee

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// SerializeParams is to used to serialize the inputParams request .
// Eg : Customer : { FirstName : "John" } is serialized as "customer[first_name]" : "John".
func SerializeParams(params interface{}) *url.Values {
	values := &url.Values{}
	serialize(values, params, "", false)
	return values
}

// SerializeListParams is to used to serialize the inputParams of list request.
func SerializeListParams(params interface{}) *url.Values {
	values := &url.Values{}
	serialize(values, params, "", true)
	return values
}

func serialize(values *url.Values, params interface{}, prefix string, isList bool) {
	if params == nil {
		return
	}
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		serializeStruct(values, v, prefix, isList)
	}
}

func serializeStruct(values *url.Values, v reflect.Value, prefix string, isList bool) {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		// Special handling for CustomFields: Flatten them
		if field.Name == "CustomFields" {
			val := v.Field(i)
			if val.Kind() == reflect.Ptr {
				if val.IsNil() {
					continue
				}
				val = val.Elem()
			}
			if val.Kind() == reflect.Map {
				// Pass prefix directly to serializeMap so keys are merged at current level
				// e.g. "cf_app_name" or "customer[cf_app_name]"
				serializeMap(values, val, prefix, isList)
			}
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, opts, _ := strings.Cut(tag, ",")
		if name == "" {
			name = field.Name
		}

		key := name
		if prefix != "" {
			key = prefix + "[" + name + "]"
		}

		value := v.Field(i)
		if strings.Contains(opts, "omitempty") && isEmptyValue(value) {
			continue
		}

		// Handle pointer
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}
			value = value.Elem()
		}

		serializeValue(values, value, key, isList, field.Type)
	}
}

func serializeValue(values *url.Values, v reflect.Value, key string, isList bool, t reflect.Type) {
	// Special cases: map[string]interface{} or []map[string]interface{}
	// These are typically "bag of attributes" (like MetaData) and should be serialized as a JSON string.
	if isMapStringInterface(t) || isSliceMapStringInterface(t) {
		b, err := json.Marshal(v.Interface())
		if err == nil {
			values.Set(key, string(b))
		}
		return
	}

	switch v.Kind() {
	case reflect.Struct:
		serializeStruct(values, v, key, isList)
	case reflect.Map:
		serializeMap(values, v, key, isList)
	case reflect.Slice, reflect.Array:
		serializeSlice(values, v, key, isList)
	default:
		values.Set(key, fmt.Sprintf("%v", v.Interface()))
	}
}

func isMapStringInterface(t reflect.Type) bool {
	return t.Kind() == reflect.Map && t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Interface
}

func isSliceMapStringInterface(t reflect.Type) bool {
	return t.Kind() == reflect.Slice && isMapStringInterface(t.Elem())
}

func serializeMap(values *url.Values, v reflect.Value, prefix string, isList bool) {
	iter := v.MapRange()
	for iter.Next() {
		k := iter.Key().String()
		val := iter.Value()

		var newKey string
		if prefix != "" {
			newKey = prefix + "[" + k + "]"
		} else {
			newKey = k
		}

		// Handle interface value in map
		if val.Kind() == reflect.Interface {
			val = val.Elem()
		}

		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				continue
			}
			val = val.Elem()
		}

		serializeValue(values, val, newKey, isList, val.Type())
	}
}

func serializeSlice(values *url.Values, v reflect.Value, prefix string, isList bool) {
	if isList {
		// List Params: Serialize as JSON array of strings
		var items []string
		for i := 0; i < v.Len(); i++ {
			val := v.Index(i)
			if val.Kind() == reflect.Ptr && !val.IsNil() {
				val = val.Elem()
			}
			items = append(items, fmt.Sprintf("%v", val.Interface()))
		}
		if len(items) > 0 {
			b, _ := json.Marshal(items) // ["a","b"]
			values.Set(prefix, string(b))
		}
		return
	}

	// Standard Params
	// Check if slice of structs
	elemType := v.Type().Elem()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	if elemType.Kind() == reflect.Struct {
		// Inverted Indexing for arrays of structs: prefix[field][index]
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if item.Kind() == reflect.Ptr {
				if item.IsNil() {
					continue
				}
				item = item.Elem()
			}
			serializeStructInverted(values, item, prefix, i)
		}
	} else {
		// Normal Indexing: prefix[index]
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			key := prefix + "[" + strconv.Itoa(i) + "]"

			if item.Kind() == reflect.Ptr {
				if item.IsNil() {
					continue
				}
				item = item.Elem()
			}
			serializeValue(values, item, key, isList, item.Type())
		}
	}
}

func serializeStructInverted(values *url.Values, v reflect.Value, prefix string, index int) {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, opts, _ := strings.Cut(tag, ",")
		if name == "" {
			name = field.Name
		}

		val := v.Field(i)
		if strings.Contains(opts, "omitempty") && isEmptyValue(val) {
			continue
		}

		key := prefix + "[" + name + "][" + strconv.Itoa(index) + "]"

		// If the value is complex (slice/map), encode as JSON string (per original behavior)
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				continue
			}
			val = val.Elem()
		}

		if val.Kind() == reflect.Map || val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			b, err := json.Marshal(val.Interface())
			if err == nil {
				values.Set(key, string(b))
			}
		} else {
			values.Set(key, fmt.Sprintf("%v", val.Interface()))
		}
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

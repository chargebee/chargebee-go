package chargebee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func getParamTypes(params interface{}) map[string]string {
	v := reflect.TypeOf(params).Elem()
	m := map[string]string{}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		m[f.Name] = f.Type.String()
	}
	return m
}

//SerializeParams is to used to serialize the inputParams request .
// Eg : Customer : { FirstName : "John" } is serialized as "customer[first_name]" : "John".
func SerializeParams(params interface{}) *url.Values {
	queryParams, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	data := json.NewDecoder(strings.NewReader(string(queryParams)))
	data.UseNumber()
	var m map[string]interface{}
	if err := data.Decode(&m); err != nil {
		panic(err)
	}
	serParams := make(map[string]interface{})
	parseMap(m, serParams, "", "", getParamTypes(params))
	body := &url.Values{}
	for k, v := range serParams {
		body.Set(k, fmt.Sprintf("%v", v))
	}
	return body
}

func parseMap(aMap, serParams map[string]interface{}, prefix string, idx string, paramTypes map[string]string) {
	for key, val := range aMap {
		switch value := val.(type) {
		case map[string]interface{}:
			if paramTypes[camelCase(key)] == "map[string]interface {}" {
				buf := new(bytes.Buffer)
				enc := json.NewEncoder(buf)
				err := enc.Encode(value)
				if err != nil {
					fmt.Println(err.Error())
				}
				if prefix != "" && idx != "" {
					key = prefix + "[" + key + "]" + "[" + idx + "]"
				} else if prefix != "" {
					key = prefix + "[" + key + "]"
				}
				serParams[key] = buf.String()
			} else {
				if prefix != "" {
					key = prefix + "[" + key + "]"
				}
				parseMap(val.(map[string]interface{}), serParams, key, "", paramTypes)
			}
		case []interface{}:
			parseArray(val.([]interface{}), serParams, key, "", paramTypes)
		default:
			if prefix != "" && idx != "" {
				key = prefix + "[" + key + "]" + "[" + idx + "]"
			} else if prefix != "" {
				key = prefix + "[" + key + "]"
			}
			serParams[key] = value
		}
	}
}
func parseArray(anArray []interface{}, serParams map[string]interface{}, prefix string, idx string, paramTypes map[string]string) {
	if paramTypes[camelCase(prefix)] == "[]map[string]interface {}" {
		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		err := enc.Encode(anArray)
		if err != nil {
			fmt.Println(err.Error())
		}
		serParams[prefix] = buf.String()
		return
	}
	for i, val := range anArray {
		switch value := val.(type) {
		case map[string]interface{}:
			for mk, mv := range val.(map[string]interface{}) {
				k := prefix + "[" + mk + "]" + "[" + strconv.Itoa(i) + "]"
				if mvArray, ok := mv.([]interface{}); ok {
					if out, err := json.Marshal(mvArray); err == nil {
						serParams[k] = string(out)
					}
				} else if mvSlice, ok := mv.(map[string]interface{}); ok {
					if out, err := json.Marshal(mvSlice); err == nil {
						serParams[k] = string(out)
					}
				} else {
					serParams[k] = mv
				}
			}

		default:
			k := prefix + "[" + strconv.Itoa(i) + "]"
			serParams[k] = value
		}
	}
}

//SerializeListParams is to used to serialize the inputParams of list request.
func SerializeListParams(params interface{}) *url.Values {
	queryParams, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	data := json.NewDecoder(strings.NewReader(string(queryParams)))
	data.UseNumber()
	var m map[string]interface{}
	if err := data.Decode(&m); err != nil {
		panic(err)
	}

	serListParams := make(map[string]interface{})
	parseMapListParams(m, serListParams, "")
	body := &url.Values{}
	for k, v := range serListParams {
		switch val := v.(type) {
		case []interface{}:
			value := "[\""
			str := []string{}
			for _, element := range val {
				str = append(str, fmt.Sprintf("%v", element))
			}
			value = value + strings.Join(str, "\",\"")
			value = value + "\"]"
			body.Set(k, value)
		default:
			body.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return body

}
func parseMapListParams(aMap, serListParams map[string]interface{}, prefix string) {
	for key, val := range aMap {
		switch value := val.(type) {
		case map[string]interface{}:
			parseMapListParams(val.(map[string]interface{}), serListParams, key)
		default:
			if prefix != "" {
				k := prefix + "[" + key + "]"
				serListParams[k] = value
			} else {
				serListParams[key] = value
			}

		}
	}
}

func camelCase(str string) string {
	var val []string
	res := ""
	val = strings.SplitAfter(str, "_")
	for _, element := range val {
		if element != "_" {
			res = res + strings.Title(strings.Trim(element, "_"))
		}
	}
	return res
}

//customFieldExtraction is used extract the customFields from response and append it in Result struct
func customFieldExtraction(v interface{}, resJSON []byte) {
	switch v.(type) {
	case *Result:
		prepareResultCF(resJSON, v)

	case *ResultList:
		prepareResultListCF(resJSON, v)

	}
}

func customMapping(val interface{}) map[string]interface{} {
	custom := make(map[string]interface{})
	if rec, ok := val.(map[string]interface{}); ok {
		for k1, v1 := range rec {
			if strings.HasPrefix(k1, "cf_") {
				custom[k1] = v1
			}
		}
	}
	return custom
}

func prepareResultCF(resbody []byte, v interface{}) {
	data := json.NewDecoder(strings.NewReader(string(resbody)))
	data.UseNumber()
	var m map[string]interface{}
	if err := data.Decode(&m); err != nil {
		panic(err)
	}

	for key, val := range m {
		if key == "customer" {
			v.(*Result).Customer.CustomField = customMapping(val)
		} else if key == "subscription" {
			v.(*Result).Subscription.CustomField = customMapping(val)
		} else if key == "addon" {
			v.(*Result).Addon.CustomField = customMapping(val)
		} else if key == "plan" {
			v.(*Result).Plan.CustomField = customMapping(val)
		} else if key == "item" {
			v.(*Result).Item.CustomField = customMapping(val)
		} else if key == "item_price" {
			v.(*Result).ItemPrice.CustomField = customMapping(val)
		} else if key == "item_family" {
			v.(*Result).ItemFamily.CustomField = customMapping(val)
		}

	}

}

func prepareResultListCF(resbody []byte, v interface{}) {
	data := json.NewDecoder(strings.NewReader(string(resbody)))
	data.UseNumber()
	var m map[string]interface{}
	if err := data.Decode(&m); err != nil {
		panic(err)
	}
	for index, val := range m["list"].([]interface{}) {
		for key, value := range val.(map[string]interface{}) {
			if key == "customer" {
				v.(*ResultList).List[index].Customer.CustomField = customMapping(value)
			} else if key == "subscription" {
				v.(*ResultList).List[index].Subscription.CustomField = customMapping(value)
			} else if key == "addon" {
				v.(*ResultList).List[index].Addon.CustomField = customMapping(value)
			} else if key == "plan" {
				v.(*ResultList).List[index].Plan.CustomField = customMapping(value)
			} else if key == "item" {
				v.(*ResultList).List[index].Item.CustomField = customMapping(value)
			} else if key == "item_price" {
				v.(*ResultList).List[index].ItemPrice.CustomField = customMapping(value)
			} else if key == "item_family" {
				v.(*ResultList).List[index].ItemFamily.CustomField = customMapping(value)
			}

		}
	}
}

//Bool returns a pointer to the bool value passed.
func Bool(val bool) *bool {
	return &val
}

//BoolValue returns the value of the bool pointer passed or false if the pointer is nil.
func BoolValue(val *bool) bool {
	if val != nil {
		return *val
	}
	return false
}

//Int32 returns a pointer to the int32 value passed.
func Int32(val int32) *int32 {
	return &val
}

//Int32Value returns the value of the int32 pointer passed or 0 if the pointer is nil.
func Int32Value(val *int32) int32 {
	if val != nil {
		return *val
	}
	return 0
}

//Int64 returns a pointer to the int64 value passed.
func Int64(val int64) *int64 {
	return &val
}

//Int64Value returns the value of the int64 pointer passed or 0 if the pointer is nil.
func Int64Value(val *int64) int64 {
	if val != nil {
		return *val
	}
	return 0
}

//Float64 returns a pointer to the float64 value passed.
func Float64(val float64) *float64 {
	return &val
}

//Float64Value returns the value of the float64 pointer passed or 0 if the pointer is nil.
func Float64Value(val *float64) float64 {
	if val != nil {
		return *val
	}
	return 0
}

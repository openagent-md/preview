package types

import (
	"encoding/json"
	"fmt"

	"github.com/zclconf/go-cty/cty"
)

func CtyValueStringDefault(def string, val cty.Value) string {
	str, err := CtyValueString(val)
	if err != nil {
		return def
	}
	return str
}

// CtyValueString converts a cty.Value to a string.
// It supports only primitive types - bool, number, and string.
// As a special case, it also supports map[string]interface{} with key "value".
func CtyValueString(val cty.Value) (string, error) {
	switch {
	case val.Type().IsListType():
		vals := val.AsValueSlice()
		strs := make([]string, 0, len(vals))
		for _, ele := range vals {
			str, err := CtyValueString(ele)
			if err != nil {
				return "", err
			}
			strs = append(strs, str)
		}
		d, _ := json.Marshal(strs)
		return string(d), nil
	case val.Type().IsMapType():
		output := make(map[string]string)
		for k, v := range val.AsValueMap() {
			str, err := CtyValueString(v)
			if err != nil {
				return "", err
			}
			output[k] = str
		}

		d, _ := json.Marshal(output)
		return string(d), nil
	}

	switch val.Type() {
	case cty.Bool:
		if val.True() {
			return "true", nil
		}
		return "false", nil
	case cty.Number:
		return val.AsBigFloat().String(), nil
	case cty.String:
		//nolint:gocritic // string type asserted above
		return val.AsString(), nil
	// We may also have a map[string]interface{} with key "value".
	case cty.Map(cty.String):
		valval, ok := val.AsValueMap()["value"]
		if !ok {
			return "", fmt.Errorf("map does not have key 'value'")
		}
		return CtyValueString(valval)
	default:
		return "", fmt.Errorf("only primitive types are supported - bool, number, and string. Found %s", val.Type().GoString())
	}
}

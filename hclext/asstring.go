package hclext

import "github.com/zclconf/go-cty/cty"

func AsString(v cty.Value) (string, bool) {
	if v.IsNull() || !v.IsKnown() {
		return "", false
	}

	switch {
	case v.Type().Equals(cty.String):
		//nolint:gocritic // string type asserted
		return v.AsString(), true
	case v.Type().Equals(cty.Number):
		// TODO: Float vs Int?
		return v.AsBigFloat().String(), true
	case v.Type().Equals(cty.Bool):
		if v.True() {
			return "true", true
		}
		return "false", true

	}

	return "", false
}

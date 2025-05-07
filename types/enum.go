package types

import (
	"fmt"
	"strings"

	"github.com/coder/terraform-provider-coder/v2/provider"
)

// TODO: Just use the provider type directly.
type ParameterType provider.OptionType

const (
	ParameterTypeString     ParameterType = "string"
	ParameterTypeNumber     ParameterType = "number"
	ParameterTypeBoolean    ParameterType = "bool"
	ParameterTypeListString ParameterType = "list(string)"
)

func (t ParameterType) Valid() error {
	switch t {
	case ParameterTypeString, ParameterTypeNumber, ParameterTypeBoolean, ParameterTypeListString:
		return nil
	default:
		return fmt.Errorf("invalid parameter type %q, expected one of [%s]", t,
			strings.Join([]string{
				string(ParameterTypeString),
				string(ParameterTypeNumber),
				string(ParameterTypeBoolean),
				string(ParameterTypeListString),
			}, ", "))
	}
}

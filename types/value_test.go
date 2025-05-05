package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"

	"github.com/coder/preview/types"
)

func TestSafeHCLString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input types.HCLString

		asString string
		known    bool
		valid    bool
	}{
		{
			name:     "empty",
			asString: types.UnknownStringValue,
		},
		{
			name:     "string_literal",
			input:    types.StringLiteral("hello world"),
			asString: "hello world",
			known:    true,
			valid:    true,
		},
		{
			name: "number",
			input: types.HCLString{
				Value: cty.NumberIntVal(1),
			},
			asString: "1",
			known:    true,
			valid:    true,
		},
		{
			name: "bool",
			input: types.HCLString{
				Value: cty.BoolVal(true),
			},
			asString: "true",
			known:    true,
			valid:    true,
		},
		// Crazy ideas
		{
			name: "null",
			input: types.HCLString{
				Value: cty.NullVal(cty.NilType),
			},
			asString: types.UnknownStringValue,
			known:    false,
			valid:    false,
		},
		{
			name: "empty_string_list",
			input: types.HCLString{
				Value: cty.ListValEmpty(cty.String),
			},
			asString: types.UnknownStringValue,
			known:    false,
			valid:    false,
		},
		{
			name: "dynamic",
			input: types.HCLString{
				Value: cty.DynamicVal,
			},
			asString: types.UnknownStringValue,
			known:    false,
			valid:    false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.asString, tc.input.AsString())
			require.Equal(t, tc.known, tc.input.IsKnown(), "known")
			require.Equal(t, tc.valid, tc.input.Valid(), "valid")

			_, err := json.Marshal(tc.input)
			require.NoError(t, err)
		})
	}
}

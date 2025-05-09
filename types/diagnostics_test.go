package types_test

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/require"

	"github.com/coder/preview/types"
)

func TestDiagnosticsJSON(t *testing.T) {

	diags := types.Diagnostics{
		{
			Severity: hcl.DiagWarning,
			Summary:  "Some summary",
			Detail:   "Some detail",
		},
		{
			Severity: hcl.DiagError,
			Summary:  "Some summary",
			Detail:   "Some detail",
		},
	}

	data, err := json.Marshal(diags)
	require.NoError(t, err, "marshal")

	var newDiags types.Diagnostics
	err = json.Unmarshal(data, &newDiags)
	require.NoError(t, err, "unmarshal")

	require.Equal(t, diags, newDiags)
}

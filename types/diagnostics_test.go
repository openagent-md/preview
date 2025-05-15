package types_test

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/require"

	"github.com/coder/preview/types"
)

func TestDiagnosticExtra(t *testing.T) {
	diag := &hcl.Diagnostic{
		Severity: hcl.DiagWarning,
		Summary:  "Some summary",
		Detail:   "Some detail",
	}

	extra := types.ExtractDiagnosticExtra(diag)
	require.Empty(t, extra.Code)

	// Set it
	extra.Code = "foobar"
	types.SetDiagnosticExtra(diag, extra)

	extra = types.ExtractDiagnosticExtra(diag)
	require.Equal(t, "foobar", extra.Code)

	// Set it again
	extra.Code = "bazz"
	types.SetDiagnosticExtra(diag, extra)

	extra = types.ExtractDiagnosticExtra(diag)
	require.Equal(t, "bazz", extra.Code)
}

func TestDiagnosticsJSON(t *testing.T) {
	diags := types.Diagnostics{
		{
			Severity: hcl.DiagWarning,
			Summary:  "Some summary",
			Detail:   "Some detail",
			Extra: types.DiagnosticExtra{
				Code: "foobar",
			},
		},
		{
			Severity: hcl.DiagError,
			Summary:  "Some summary",
			Detail:   "Some detail",
			Extra: types.DiagnosticExtra{
				Code: "",
			},
		},
	}

	data, err := json.Marshal(diags)
	require.NoError(t, err, "marshal")

	var newDiags types.Diagnostics
	err = json.Unmarshal(data, &newDiags)
	require.NoError(t, err, "unmarshal")

	require.Equal(t, diags, newDiags)
}

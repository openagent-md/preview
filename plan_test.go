package preview_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/openagent-md/preview"
	"github.com/openagent-md/preview/types"
)

func TestPlanJSONHook(t *testing.T) {
	t.Parallel()

	t.Run("Empty plan", func(t *testing.T) {
		t.Parallel()

		dirFS := os.DirFS("testdata/static")
		_, diags := preview.Preview(t.Context(), preview.Input{
			PlanJSONPath:    "",
			PlanJSON:        []byte("{}"),
			ParameterValues: nil,
			Owner:           types.WorkspaceOwner{},
		}, dirFS)
		require.False(t, diags.HasErrors())
	})
}

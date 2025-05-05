package preview

import (
	"io/fs"

	"github.com/aquasecurity/trivy/pkg/iac/terraform"
	tfcontext "github.com/aquasecurity/trivy/pkg/iac/terraform/context"
	"github.com/zclconf/go-cty/cty"
	"golang.org/x/xerrors"
)

func workspaceOwnerHook(_ fs.FS, input Input) (func(ctx *tfcontext.Context, blocks terraform.Blocks, inputVars map[string]cty.Value), error) {
	ownerValue, err := input.Owner.ToCtyValue()
	if err != nil {
		return nil, xerrors.Errorf("failed to convert owner value", err)
	}

	return func(ctx *tfcontext.Context, blocks terraform.Blocks, inputVars map[string]cty.Value) {
		for _, block := range blocks.OfType("data") {
			// TODO: Does it have to be me?
			if block.TypeLabel() == "coder_workspace_owner" && block.NameLabel() == "me" {
				block.Context().Parent().Set(ownerValue,
					"data", block.TypeLabel(), block.NameLabel())
			}
		}
	}, nil
}

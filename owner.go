package preview

import (
	"io/fs"

	"github.com/aquasecurity/trivy/pkg/iac/terraform"
	tfcontext "github.com/aquasecurity/trivy/pkg/iac/terraform/context"
	"github.com/zclconf/go-cty/cty"
)

func workspaceOwnerHook(dfs fs.FS, input Input) (func(ctx *tfcontext.Context, blocks terraform.Blocks, inputVars map[string]cty.Value), error) {
	ownerValue, err := input.Owner.ToCtyValue()
	if err != nil {
		return nil, err
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

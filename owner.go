package preview

import (
	"io/fs"

	"github.com/aquasecurity/trivy/pkg/iac/terraform"
	tfcontext "github.com/aquasecurity/trivy/pkg/iac/terraform/context"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"golang.org/x/xerrors"
)

func workspaceOwnerHook(dfs fs.FS, input Input) (func(ctx *tfcontext.Context, blocks terraform.Blocks, inputVars map[string]cty.Value), error) {
	ownerValue, err := gocty.ToCtyValue(input.Owner, cty.Object(map[string]cty.Type{
		"id":             cty.String,
		"name":           cty.String,
		"full_name":      cty.String,
		"email":          cty.String,
		"ssh_public_key": cty.String,
		"groups":         cty.List(cty.String),
		"login_type":     cty.String,
		"rbac_roles": cty.List(cty.Object(
			map[string]cty.Type{
				"name":   cty.String,
				"org_id": cty.String,
			},
		)),
	}))
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

package types

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"golang.org/x/xerrors"
)

// Based on https://github.com/coder/terraform-provider-coder/blob/9a745586b23a9cb5de2f65a2dcac12e48b134ffa/provider/workspace_owner.go#L72
type WorkspaceOwner struct {
	ID           string `json:"id" cty:"id"`
	Name         string `json:"name" cty:"name"`
	FullName     string `json:"full_name" cty:"full_name"`
	Email        string `json:"email" cty:"email"`
	SSHPublicKey string `json:"ssh_public_key" cty:"ssh_public_key"`
	// SSHPrivateKey is intentionally omitted for now, due to the security risk
	// that exposing it poses.
	// SSHPrivateKey string `json:"ssh_private_key" cty:"ssh_private_key"`
	Groups []string `json:"groups" cty:"groups"`
	// SessionToken is intentionally omitted for now, due to the security risk
	// that exposing it poses.
	// SessionToken string `json:"session_token" cty:"session_token"`
	// OIDCAccessToken is intentionally omitted for now, due to the security risk
	// that exposing it poses.
	// OIDCAccessToken string `json:"oidc_access_token" cty:"oidc_access_token"`
	LoginType string                   `json:"login_type" cty:"login_type"`
	RBACRoles []WorkspaceOwnerRBACRole `json:"rbac_roles" cty:"rbac_roles"`
}

type WorkspaceOwnerRBACRole struct {
	Name  string `json:"name" cty:"name"`
	OrgID string `json:"org_id" cty:"org_id"`
}

func (o *WorkspaceOwner) ToCtyValue() (cty.Value, error) {
	if o.Groups == nil {
		o.Groups = make([]string, 0)
	}
	if o.RBACRoles == nil {
		o.RBACRoles = make([]WorkspaceOwnerRBACRole, 0)
	}

	ownerValue, err := gocty.ToCtyValue(o, cty.Object(map[string]cty.Type{
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
		return cty.Value{}, xerrors.Errorf("failed to convert owner value: %w", err)
	}
	return ownerValue, nil
}

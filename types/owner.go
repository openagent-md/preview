package types

import (
	"github.com/google/uuid"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// Based on https://github.com/coder/terraform-provider-coder/blob/9a745586b23a9cb5de2f65a2dcac12e48b134ffa/provider/workspace_owner.go#L72
type WorkspaceOwner struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	SSHPublicKey string    `json:"ssh_public_key"`
	// SSHPrivateKey is intentionally omitted for now, due to the security risk
	// that exposing it poses.
	// SSHPrivateKey string `json:"ssh_private_key"`
	Groups []string `json:"groups"`
	// SessionToken is intentionally omitted for now, due to the security risk
	// that exposing it poses.
	// SessionToken string `json:"session_token"`
	// OIDCAccessToken is intentionally omitted for now, due to the security risk
	// that exposing it poses.
	// OIDCAccessToken string `json:"oidc_access_token"`
	LoginType string                   `json:"login_type"`
	RBACRoles []WorkspaceOwnerRBACRole `json:"rbac_roles"`
}

func (o *WorkspaceOwner) ToCtyValue() (cty.Value, error) {
	convertedGroups, err := gocty.ToCtyValue(o.Groups, cty.List(cty.String))
	if err != nil {
		return cty.Value{}, err
	}

	roleValues := make([]cty.Value, 0, len(o.RBACRoles))
	for _, role := range o.RBACRoles {
		roleValue, err := role.ToCtyValue()
		if err != nil {
			return cty.Value{}, err
		}
		roleValues = append(roleValues, roleValue)
	}
	var convertedRoles cty.Value = cty.ListValEmpty(WorkspaceOwnerRBACRole{}.CtyType())
	if len(roleValues) > 0 {
		convertedRoles = cty.ListVal(roleValues)
	}

	return cty.ObjectVal(map[string]cty.Value{
		"id":             cty.StringVal(o.ID.String()),
		"name":           cty.StringVal(o.Name),
		"full_name":      cty.StringVal(o.FullName),
		"email":          cty.StringVal(o.Email),
		"ssh_public_key": cty.StringVal(o.SSHPublicKey),
		"groups":         convertedGroups,
		"login_type":     cty.StringVal(o.LoginType),
		"rbac_roles":     convertedRoles,
	}), nil
}

type WorkspaceOwnerRBACRole struct {
	Name  string    `json:"name"`
	OrgID uuid.UUID `json:"org_id"`
}

func (_ WorkspaceOwnerRBACRole) CtyType() cty.Type {
	return cty.Object(map[string]cty.Type{
		"name":   cty.String,
		"org_id": cty.String,
	})
}

func (r *WorkspaceOwnerRBACRole) ToCtyValue() (cty.Value, error) {
	return cty.ObjectVal(map[string]cty.Value{
		"name":   cty.StringVal(r.Name),
		"org_id": cty.StringVal(r.OrgID.String()),
	}), nil
}

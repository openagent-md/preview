package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coder/preview/types"
)

func TestToCtyValue(t *testing.T) {
	t.Parallel()

	owner := types.WorkspaceOwner{
		ID:           "f6457744-3e16-45b2-b3b0-80c2df491c99",
		Name:         "Nissa",
		FullName:     "Nissa, Worldwaker",
		Email:        "nissa@coder.com",
		SSHPublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSHXs/HCgZlpEBOXLvLw4KaOrhy1DM1Vw6M/HPVE/UA\n",
		Groups:       []string{"Everyone", "Planeswalkers", "Green"},
		LoginType:    "password",
		RBACRoles: []types.WorkspaceOwnerRBACRole{
			{Name: "User Admin"},
			{Name: "Organization User Admin", OrgID: "5af9253a-ecde-4a71-b8f5-c8d15be9e52b"},
		},
	}

	ownerValue, err := owner.ToCtyValue()
	require.NoError(t, err)

	require.Equal(t, owner.ID, ownerValue.AsValueMap()["id"].AsString())
	require.Equal(t, owner.Name, ownerValue.AsValueMap()["name"].AsString())
	require.Equal(t, owner.SSHPublicKey, ownerValue.AsValueMap()["ssh_public_key"].AsString())
	for i, it := range owner.Groups {
		require.Equal(t, it, ownerValue.AsValueMap()["groups"].AsValueSlice()[i].AsString())
	}
	for i, it := range owner.RBACRoles {
		roleValueMap := ownerValue.AsValueMap()["rbac_roles"].AsValueSlice()[i].AsValueMap()
		require.Equal(t, it.Name, roleValueMap["name"].AsString())
		require.Equal(t, it.OrgID, roleValueMap["org_id"].AsString())
	}
}

func TestToCtyValueWithNilLists(t *testing.T) {
	t.Parallel()

	owner := types.WorkspaceOwner{
		ID:           "f6457744-3e16-45b2-b3b0-80c2df491c99",
		Name:         "Nissa",
		FullName:     "Nissa, Worldwaker",
		Email:        "nissa@coder.com",
		SSHPublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSHXs/HCgZlpEBOXLvLw4KaOrhy1DM1Vw6M/HPVE/UA\n",
		Groups:       nil,
		LoginType:    "password",
		RBACRoles:    nil,
	}

	ownerValue, err := owner.ToCtyValue()
	require.NoError(t, err)
	require.Empty(t, ownerValue.AsValueMap()["groups"].AsValueSlice())
	require.Empty(t, ownerValue.AsValueMap()["rbac_roles"].AsValueSlice())
}

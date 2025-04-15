package types

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToCtyValue(t *testing.T) {
	owner := WorkspaceOwner{
		ID:           uuid.MustParse("f6457744-3e16-45b2-b3b0-80c2df491c99"),
		Name:         "Nissa",
		FullName:     "Nissa, Worldwaker",
		Email:        "nissa@coder.com",
		SSHPublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSHXs/HCgZlpEBOXLvLw4KaOrhy1DM1Vw6M/HPVE/UA\n",
		Groups:       []string{"Everyone", "Planeswalkers", "Green"},
		LoginType:    "password",
		RBACRoles: []WorkspaceOwnerRBACRole{
			{Name: "User Admin"},
			{Name: "Organization User Admin", OrgID: uuid.MustParse("5af9253a-ecde-4a71-b8f5-c8d15be9e52b")},
		},
	}

	_, err := owner.ToCtyValue()
	require.NoError(t, err)
}

func TestToCtyValueWithNilLists(t *testing.T) {
	owner := WorkspaceOwner{
		ID:           uuid.MustParse("f6457744-3e16-45b2-b3b0-80c2df491c99"),
		Name:         "Nissa",
		FullName:     "Nissa, Worldwaker",
		Email:        "nissa@coder.com",
		SSHPublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSHXs/HCgZlpEBOXLvLw4KaOrhy1DM1Vw6M/HPVE/UA\n",
		Groups:       nil,
		LoginType:    "password",
		RBACRoles:    nil,
	}

	_, err := owner.ToCtyValue()
	require.NoError(t, err)
}

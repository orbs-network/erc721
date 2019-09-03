package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestERC721_balanceOf(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newERC721Harness()
	h.deployContract(t, owner)

	balance := h.balanceOf(t, owner, owner.AddressAsBytes())
	require.EqualValues(t, 0, balance)
}

func TestERC721_mint(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newERC721Harness()
	h.deployContract(t, owner)

	paintBlackSquare(t, h, owner)
}


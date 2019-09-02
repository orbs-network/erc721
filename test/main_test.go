package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestERC721_balanceOf(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, sender)

	balance := h.balanceOf(t, sender, sender.AddressAsBytes())
	require.EqualValues(t, 0, balance)
}

func TestERC721_mint(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, sender)

	tokenId := h.mint(t, sender, `{"title":"Black Square","type":"Painting"}`)
	require.EqualValues(t, 0, tokenId)

	metadata := h.tokenMetadata(t, sender, tokenId)
	require.EqualValues(t, `{"title":"Black Square","type":"Painting"}`, metadata)

	balance := h.balanceOf(t, sender, sender.AddressAsBytes())
	require.EqualValues(t, 1, balance)

	tokenOwner := h.ownerOf(t, sender, tokenId)
	require.EqualValues(t, sender.AddressAsBytes(), tokenOwner)
}

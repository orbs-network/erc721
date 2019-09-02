package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestERC721_balanceOf(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, owner)

	balance := h.balanceOf(t, owner, owner.AddressAsBytes())
	require.EqualValues(t, 0, balance)
}

func TestERC721_mint(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, owner)

	tokenId := h.mint(t, owner, `{"title":"Black Square","type":"Painting"}`)
	require.EqualValues(t, 0, tokenId)

	metadata := h.tokenMetadata(t, owner, tokenId)
	require.EqualValues(t, `{"title":"Black Square","type":"Painting"}`, metadata)

	balance := h.balanceOf(t, owner, owner.AddressAsBytes())
	require.EqualValues(t, 1, balance)

	tokenOwner := h.ownerOf(t, owner, tokenId)
	require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)
}

func TestERC721_safeTransferFrom(t *testing.T) {
	owner, _ := orbs.CreateAccount()
	buyer, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, owner)

	tokenId := h.mint(t, owner, `{"title":"Black Square","type":"Painting"}`)
	require.EqualValues(t, 0, tokenId)

	metadata := h.tokenMetadata(t, owner, tokenId)
	require.EqualValues(t, `{"title":"Black Square","type":"Painting"}`, metadata)

	ownerBalance := h.balanceOf(t, owner, owner.AddressAsBytes())
	buyerBalance := h.balanceOf(t, owner, buyer.AddressAsBytes())
	require.EqualValues(t, 1, ownerBalance)
	require.EqualValues(t, 0, buyerBalance)

	tokenOwner := h.ownerOf(t, owner, tokenId)
	require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

	err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
	require.NoError(t, err)

	ownerBalance = h.balanceOf(t, owner, owner.AddressAsBytes())
	buyerBalance = h.balanceOf(t, owner, buyer.AddressAsBytes())
	require.EqualValues(t, 0, ownerBalance)
	require.EqualValues(t, 1, buyerBalance)

	tokenOwner = h.ownerOf(t, owner, tokenId)
	require.EqualValues(t, buyer.AddressAsBytes(), tokenOwner)
}

func TestERC721_safeTransferFromWrongAddress(t *testing.T) {
	owner, _ := orbs.CreateAccount()
	buyer, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, owner)

	tokenId := h.mint(t, owner, `{"title":"Black Square","type":"Painting"}`)
	require.EqualValues(t, 0, tokenId)

	metadata := h.tokenMetadata(t, owner, tokenId)
	require.EqualValues(t, `{"title":"Black Square","type":"Painting"}`, metadata)

	checkIfNothingChanged := func() {
		ownerBalance := h.balanceOf(t, owner, owner.AddressAsBytes())
		buyerBalance := h.balanceOf(t, owner, buyer.AddressAsBytes())
		require.EqualValues(t, 1, ownerBalance)
		require.EqualValues(t, 0, buyerBalance)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)
	}

	checkIfNothingChanged()

	t.Run("from wrong address", func(t *testing.T) {
		err := h.safeTransferFrom(t, owner, []byte{1, 2, 3}, buyer.AddressAsBytes(), tokenId)
		require.EqualError(t, err, "transfer not authorized")

		checkIfNothingChanged()
	})

	t.Run("to malformed address", func(t *testing.T) {
		err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), []byte{1, 2, 3}, tokenId)
		require.EqualError(t, err, "transfer not authorized")

		checkIfNothingChanged()
	})

	t.Run("non-existent token", func(t *testing.T) {
		err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), []byte{1, 2, 3}, 1974)
		require.EqualError(t, err, "transfer not authorized")

		checkIfNothingChanged()
	})
}
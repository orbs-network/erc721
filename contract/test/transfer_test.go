package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestERC721_transferFrom(t *testing.T) {
	owner, _ := orbs.CreateAccount()
	buyer, _ := orbs.CreateAccount()

	h := newERC721Harness()
	h.deployContract(t, owner)

	t.Run("transfer once", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)

		ownerBalance := h.balanceOf(t, owner, owner.AddressAsBytes())
		buyerBalance := h.balanceOf(t, owner, buyer.AddressAsBytes())
		require.EqualValues(t, 1, ownerBalance)
		require.EqualValues(t, 0, buyerBalance)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

		err := h.transferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		ownerBalance = h.balanceOf(t, owner, owner.AddressAsBytes())
		buyerBalance = h.balanceOf(t, owner, buyer.AddressAsBytes())
		require.EqualValues(t, 0, ownerBalance)
		require.EqualValues(t, 1, buyerBalance)

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))
	})

	t.Run("transfer twice from the same address", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

		err := h.transferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))

		err = h.transferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.EqualError(t, err, "transfer not authorized")
	})

	t.Run("from a single approved address", func(t *testing.T) {
		approvedForSingleSale, _ := orbs.CreateAccount()
		tokenId := paintBlackSquare(t, h, owner)

		err := h.approve(t, owner, approvedForSingleSale.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		err = h.transferFrom(t, approvedForSingleSale, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))
	})

	t.Run("from an universal approved address ", func(t *testing.T) {
		approvedAll, _ := orbs.CreateAccount()
		tokenId := paintBlackSquare(t, h, owner)

		err := h.setApprovalForAll(t, owner, approvedAll.AddressAsBytes(), 1)
		require.NoError(t, err)

		err = h.transferFrom(t, approvedAll, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))
	})
}

func TestERC721_transferFromWrongAddress(t *testing.T) {
	owner, _ := orbs.CreateAccount()
	buyer, _ := orbs.CreateAccount()

	h := newERC721Harness()
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
		err := h.transferFrom(t, owner, []byte{1, 2, 3}, buyer.AddressAsBytes(), tokenId)
		require.EqualError(t, err, "transfer not authorized")

		checkIfNothingChanged()
	})

	t.Run("to malformed address", func(t *testing.T) {
		err := h.transferFrom(t, owner, owner.AddressAsBytes(), []byte{1, 2, 3}, tokenId)
		require.EqualError(t, err, "transfer not authorized")

		checkIfNothingChanged()
	})

	t.Run("non-existent token", func(t *testing.T) {
		err := h.transferFrom(t, owner, owner.AddressAsBytes(), []byte{1, 2, 3}, 1974)
		require.EqualError(t, err, "transfer not authorized")

		checkIfNothingChanged()
	})
}

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

	paintBlackSquare(t, h, owner)
}

func TestERC721_safeTransferFrom(t *testing.T) {
	owner, _ := orbs.CreateAccount()
	buyer, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, owner)

	t.Run("transfer once", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)

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

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))
	})

	t.Run("transfer twice from the same address", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

		err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))

		err = h.safeTransferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.EqualError(t, err, "transfer not authorized")
	})

	t.Run("from a single approved address", func(t *testing.T) {
		approvedForSingleSale, _ := orbs.CreateAccount()
		tokenId := paintBlackSquare(t, h, owner)

		err := h.approve(t, owner, approvedForSingleSale.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		err = h.safeTransferFrom(t, approvedForSingleSale, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.EqualValues(t, buyer.AddressAsBytes(), h.ownerOf(t, owner, tokenId))
	})

	t.Run("from an universal approved address ", func(t *testing.T) {
		// FIXME implement
	})
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

func TestERC721_approve(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, owner)

	t.Run("approve once", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		approvedAddress, _ := orbs.CreateAccount()
		require.Empty(t, h.getApproved(t, owner, tokenId))

		h.approve(t, owner, approvedAddress.AddressAsBytes(), tokenId)
		require.EqualValues(t, approvedAddress.AddressAsBytes(), h.getApproved(t, owner, tokenId))
	})

	t.Run("remove approval after transfer", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		approvedAddress, _ := orbs.CreateAccount()
		buyer, _ := orbs.CreateAccount()

		err := h.approve(t, owner, approvedAddress.AddressAsBytes(), tokenId)
		require.NoError(t, err)
		require.EqualValues(t, approvedAddress.AddressAsBytes(), h.getApproved(t, owner, tokenId))

		err = h.safeTransferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.Empty(t, h.getApproved(t, owner, tokenId))
	})

	t.Run("approve bad address", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		require.Empty(t, h.getApproved(t, owner, tokenId))

		err := h.approve(t, owner, []byte{1, 2, 3}, tokenId)
		require.EqualError(t, err, "approval not authorized")

		require.Empty(t, h.getApproved(t, owner, tokenId))
	})
}
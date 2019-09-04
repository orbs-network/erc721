package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestERC721_approve(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newERC721Harness()
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

		err = h.transferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
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

func TestERC721_setApprovalForAll(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newERC721Harness()
	h.deployContract(t, owner)

	t.Run("approve once", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		approvedAddress, _ := orbs.CreateAccount()
		require.Empty(t, h.getApproved(t, owner, tokenId))

		h.approve(t, owner, approvedAddress.AddressAsBytes(), tokenId)
		require.EqualValues(t, approvedAddress.AddressAsBytes(), h.getApproved(t, owner, tokenId))
	})

	t.Run("remove approval", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		approvedAddress, _ := orbs.CreateAccount()
		anotherApprovedAddress, _ := orbs.CreateAccount()
		buyer, _ := orbs.CreateAccount()

		require.EqualValues(t, 0, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), approvedAddress.AddressAsBytes()))

		err := h.setApprovalForAll(t, owner, approvedAddress.AddressAsBytes(), 1)
		require.NoError(t, err)
		err = h.setApprovalForAll(t, owner, anotherApprovedAddress.AddressAsBytes(), 1)
		require.NoError(t, err)
		require.EqualValues(t, 1, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), approvedAddress.AddressAsBytes()))
		require.EqualValues(t, 1, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), anotherApprovedAddress.AddressAsBytes()))

		err = h.setApprovalForAll(t, owner, approvedAddress.AddressAsBytes(), 0)
		require.NoError(t, err)

		require.EqualValues(t, 0, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), approvedAddress.AddressAsBytes()))
		require.EqualValues(t, 1, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), anotherApprovedAddress.AddressAsBytes()))

		err = h.transferFrom(t, owner, approvedAddress.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.EqualError(t, err, "transfer not authorized")
	})

	t.Run("transfer does not affect approval", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		approvedAddress, _ := orbs.CreateAccount()
		buyer, _ := orbs.CreateAccount()

		require.EqualValues(t, 0, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), approvedAddress.AddressAsBytes()))

		err := h.setApprovalForAll(t, owner, approvedAddress.AddressAsBytes(), 1)
		require.NoError(t, err)
		require.EqualValues(t, 1, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), approvedAddress.AddressAsBytes()))

		err = h.transferFrom(t, owner, owner.AddressAsBytes(), buyer.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		require.EqualValues(t, 1, h.isApprovedForAll(t, owner, owner.AddressAsBytes(), approvedAddress.AddressAsBytes()))
	})

	t.Run("approve bad address", func(t *testing.T) {
		tokenId := paintBlackSquare(t, h, owner)
		require.Empty(t, h.getApproved(t, owner, tokenId))

		err := h.approve(t, owner, []byte{1, 2, 3}, tokenId)
		require.EqualError(t, err, "approval not authorized")

		require.Empty(t, h.getApproved(t, owner, tokenId))
	})
}
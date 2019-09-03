package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestERC721_safeTransferFrom(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newERC721Harness()
	h.deployContract(t, owner)

	r := newErc721ReceiverHarness()
	r.deployContract(t, owner)

	t.Run("transfer once", func(t *testing.T) {
		r.acceptTokens(t, owner)
		receiverContractNameAsBytes := orbs.ContractNameToAddressAsBytes(r.contractName)
		receivedTokens := r.receivedTokens(t, owner)

		tokenId := paintBlackSquare(t, h, owner)

		ownerBalance := h.balanceOf(t, owner, owner.AddressAsBytes())
		buyerBalance := h.balanceOf(t, owner, receiverContractNameAsBytes)
		require.EqualValues(t, 1, ownerBalance)
		require.EqualValues(t, 0, buyerBalance)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

		err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), receiverContractNameAsBytes, tokenId, []byte(r.contractName))
		require.NoError(t, err)

		ownerBalance = h.balanceOf(t, owner, owner.AddressAsBytes())
		buyerBalance = h.balanceOf(t, owner, receiverContractNameAsBytes)
		require.EqualValues(t, 0, ownerBalance)
		require.EqualValues(t, 1, buyerBalance)

		require.EqualValues(t, receiverContractNameAsBytes, h.ownerOf(t, owner, tokenId))
		require.EqualValues(t, receivedTokens + 1, r.receivedTokens(t, owner))
	})

	t.Run("transfer twice from the same address", func(t *testing.T) {
		r.acceptTokens(t, owner)
		receiverContractNameAsBytes := orbs.ContractNameToAddressAsBytes(r.contractName)
		receivedTokens := r.receivedTokens(t, owner)

		tokenId := paintBlackSquare(t, h, owner)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

		err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), receiverContractNameAsBytes, tokenId, []byte(r.contractName))
		require.NoError(t, err)

		require.EqualValues(t, receiverContractNameAsBytes, h.ownerOf(t, owner, tokenId))
		require.EqualValues(t, receivedTokens + 1, r.receivedTokens(t, owner))

		err = h.safeTransferFrom(t, owner, owner.AddressAsBytes(), receiverContractNameAsBytes, tokenId, []byte(r.contractName))
		require.EqualError(t, err, "transfer not authorized")
	})


	t.Run("from a single approved address", func(t *testing.T) {
		r.acceptTokens(t, owner)
		receiverContractNameAsBytes := orbs.ContractNameToAddressAsBytes(r.contractName)
		receivedTokens := r.receivedTokens(t, owner)

		approvedForSingleSale, _ := orbs.CreateAccount()
		tokenId := paintBlackSquare(t, h, owner)

		err := h.approve(t, owner, approvedForSingleSale.AddressAsBytes(), tokenId)
		require.NoError(t, err)

		err = h.safeTransferFrom(t, approvedForSingleSale, owner.AddressAsBytes(), receiverContractNameAsBytes, tokenId, []byte(r.contractName))
		require.NoError(t, err)

		require.EqualValues(t, receiverContractNameAsBytes, h.ownerOf(t, owner, tokenId))
		require.EqualValues(t, receivedTokens + 1, r.receivedTokens(t, owner))
	})

	t.Run("from an universal approved address ", func(t *testing.T) {
		r.acceptTokens(t, owner)
		receiverContractNameAsBytes := orbs.ContractNameToAddressAsBytes(r.contractName)
		receivedTokens := r.receivedTokens(t, owner)

		approvedAll, _ := orbs.CreateAccount()
		tokenId := paintBlackSquare(t, h, owner)

		err := h.setApprovalForAll(t, owner, approvedAll.AddressAsBytes(), 1)
		require.NoError(t, err)

		err = h.safeTransferFrom(t, approvedAll, owner.AddressAsBytes(), receiverContractNameAsBytes, tokenId, []byte(r.contractName))
		require.NoError(t, err)

		require.EqualValues(t, receiverContractNameAsBytes, h.ownerOf(t, owner, tokenId))
		require.EqualValues(t, receivedTokens + 1, r.receivedTokens(t, owner))
	})

	t.Run("rolls back if the receiver does not accept tokens", func(t *testing.T) {
		r.rejectTokens(t, owner)
		receiverContractNameAsBytes := orbs.ContractNameToAddressAsBytes(r.contractName)
		receivedTokens := r.receivedTokens(t, owner)

		tokenId := paintBlackSquare(t, h, owner)

		tokenOwner := h.ownerOf(t, owner, tokenId)
		require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

		err := h.safeTransferFrom(t, owner, owner.AddressAsBytes(), receiverContractNameAsBytes, tokenId, []byte(r.contractName))
		require.EqualError(t, err, "invalid callback return value")

		require.EqualValues(t, owner.AddressAsBytes(), h.ownerOf(t, owner, tokenId))
		require.EqualValues(t, receivedTokens, r.receivedTokens(t, owner))
	})
}
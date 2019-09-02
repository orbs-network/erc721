package main

import (
	"bytes"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
)

func approve(approvedAddress []byte, tokenId uint64) {
	owner := ownerOf(tokenId)

	if !bytes.Equal(owner, address.GetSignerAddress()) {
		panic("approval not authorized")
	}

	if len(approvedAddress) != 20 {
		panic("approval not authorized")
	}

	_approve(approvedAddress, tokenId)
}

func getApproved(tokenId uint64) []byte {
	return state.ReadBytes(_approvalKey(tokenId))
}

func _approve(approvedAddress []byte, tokenId uint64) {
	state.WriteBytes(_approvalKey(tokenId), approvedAddress)
}

func _revokeApproval(tokenId uint64) {
	state.Clear(_approvalKey(tokenId))
}

func _checkApproval(from []byte, tokenId uint64) {
	approvedAddress := getApproved(tokenId)

	if approved := bytes.Equal(from, address.GetSignerAddress()) || bytes.Equal(approvedAddress, address.GetSignerAddress()); !approved {
		panic("transfer not authorized")
	}
}

func _approvalKey(tokenId uint64) []byte {
	return []byte("approved_single." + strconv.FormatUint(tokenId, 10))
}
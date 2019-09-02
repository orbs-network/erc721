package main

import (
	"bytes"
	"encoding/hex"
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


func setApprovalForAll(operator []byte, approved uint32) {
	if len(operator) != 20 {
		panic("approval not authorized")
	}

	if approved > 1 { // enforce boolean
		panic("approval not authorized")
	}

	_setOperatorPrivileges(address.GetSignerAddress(), operator, approved)
}

func isApprovedForAll(owner []byte, operator []byte) uint32 {
	return _getOperatorPrivileges(owner, operator)
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
	return []byte("approved." + strconv.FormatUint(tokenId, 10))
}

func _setOperatorPrivileges(owner []byte, operator []byte, value uint32)  {
	state.WriteUint32(_operatorKey(owner, operator), value)
}

func _getOperatorPrivileges(owner []byte, operator []byte) uint32 {
	return state.ReadUint32(_operatorKey(owner, operator))
}

func _operatorKey(owner []byte, operator []byte) []byte {
	return []byte("operator."+hex.EncodeToString(owner)+"."+hex.EncodeToString(operator))
}
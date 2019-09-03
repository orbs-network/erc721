package main

import (
	"bytes"
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/events"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
)

func approve(approvedAddress []byte, tokenId uint64) {
	owner := ownerOf(tokenId)

	if !bytes.Equal(owner, address.GetCallerAddress()) {
		panic("approval not authorized")
	}

	if len(approvedAddress) != 20 {
		panic("approval not authorized")
	}

	_approvedBy(owner, approvedAddress, tokenId)
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

	_setOperatorPrivileges(address.GetCallerAddress(), operator, approved)
}

func isApprovedForAll(owner []byte, operator []byte) uint32 {
	return _getOperatorPrivileges(owner, operator)
}

func _approvedBy(owner []byte, approvedAddress []byte, tokenId uint64) {
	state.WriteBytes(_approvalKey(tokenId), approvedAddress)
	events.EmitEvent(Approval, owner, approvedAddress, tokenId)
}

func _revokeApprovalBy(owner []byte, tokenId uint64) {
	state.Clear(_approvalKey(tokenId))
	events.EmitEvent(Approval, owner, []byte{}, tokenId)
}

func _checkApproval(from []byte, tokenId uint64) {
	approvedAddress := getApproved(tokenId)
	approvedByOperator := _getOperatorPrivileges(from, address.GetCallerAddress()) == 1

	approved := bytes.Equal(from, address.GetCallerAddress()) || bytes.Equal(approvedAddress, address.GetCallerAddress()) || approvedByOperator;
	if  !approved {
		panic("transfer not authorized")
	}
}

func _approvalKey(tokenId uint64) []byte {
	return []byte("approved." + strconv.FormatUint(tokenId, 10))
}

func _setOperatorPrivileges(owner []byte, operator []byte, value uint32)  {
	state.WriteUint32(_operatorKey(owner, operator), value)
	events.EmitEvent(ApprovalForAll, owner, operator, value)
}

func _getOperatorPrivileges(owner []byte, operator []byte) uint32 {
	return state.ReadUint32(_operatorKey(owner, operator))
}

func _operatorKey(owner []byte, operator []byte) []byte {
	return []byte("operator."+hex.EncodeToString(owner)+"."+hex.EncodeToString(operator))
}
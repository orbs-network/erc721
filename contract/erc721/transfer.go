package main

import (
	"bytes"
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/events"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
)

func transferFrom(from []byte, to []byte, tokenId uint64) {
	_checkTransferRights(from, to, tokenId)
	_transfer(from, to, tokenId)
}

func safeTransferFrom(from []byte, to []byte, tokenId uint64, toContractName string, data []byte) {
	_checkTransferRights(from, to, tokenId)
	_transfer(from, to, tokenId)
	_callOnERC721Received(address.GetCallerAddress(), from, to, tokenId, toContractName, data)
}

func _transfer(from []byte, to []byte, tokenId uint64) {
	_decBalance(from)
	_setOwner(tokenId, to)
	_revokeApprovalBy(from, tokenId)
	_incBalance(to)

	events.EmitEvent(Transfer, from, to, tokenId)
}

func _checkTransferRights(from []byte, to []byte, tokenId uint64) {
	owner := ownerOf(tokenId)
	_checkApproval(from, tokenId)

	if !bytes.Equal(from, owner) {
		panic("transfer not authorized")
	}

	if len(to) != 20 {
		panic("transfer not authorized")
	}
}

var MAGIC_ON_ERC721_RECEIVED, _ = hex.DecodeString("150b7a02")

func _callOnERC721Received(operator []byte, from []byte, to []byte, tokenId uint64, toContractName string, data []byte) {
	if !bytes.Equal(to, address.GetContractAddress(toContractName)) {
		panic("unknown contract name: " + toContractName)
	}

	value := service.CallMethod(toContractName, "onERC721Received", operator, from, tokenId, data)[0].([]byte)
	if !bytes.Equal(value, MAGIC_ON_ERC721_RECEIVED) {
		panic("invalid callback return value")
	}
}

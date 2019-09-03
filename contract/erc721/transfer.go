package main

import (
	"bytes"
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
)

func transferFrom(from []byte, to []byte, tokenId uint64) {
	_checkTransferRights(from, to, tokenId)
	_transfer(from, to, tokenId)
}

func safeTransferFrom(from []byte, to []byte, tokenId uint64, data []byte) {
	println("from", hex.EncodeToString(from))
	println("to", hex.EncodeToString(to))

	_checkTransferRights(from, to, tokenId)
	_transfer(from, to, tokenId)
	_callOnERC721Received(address.GetCallerAddress(), from, to, tokenId, data)
}

func _transfer(from []byte, to []byte, tokenId uint64) {
	_decBalance(from)
	_setOwner(tokenId, to)
	_revokeApproval(tokenId)
	_incBalance(to)
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

func _callOnERC721Received(operator []byte, from []byte, to []byte, tokenId uint64, data []byte) {
	contractName := string(data)
	if !bytes.Equal(to, address.GetContractAddress(contractName)) {
		panic("unknown contract name:" + contractName)
	}

	value := service.CallMethod(contractName, "onERC721Received", operator, from, tokenId, data)[0].([]byte)
	// FIXME temporary solution
	if !bytes.Equal(value, []byte{1, 2, 3, 4}) {
		panic("invalid callback return value")
	}
}
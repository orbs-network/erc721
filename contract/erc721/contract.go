package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
)

var PUBLIC = sdk.Export(
	name, symbol,
	mint, tokenMetadata,
	balanceOf, ownerOf,
	transferFrom, safeTransferFrom,
	approve, getApproved,
	setApprovalForAll, isApprovedForAll)

var SYSTEM = sdk.Export(_init)

var EVENTS = sdk.Export(Transfer, Approval, ApprovalForAll)

func _init() {
	_setNextTokenID(0)
}

func Transfer(from []byte, to []byte, tokenId uint64) {}
func Approval(owner []byte, address []byte, tokenId uint64) {}
func ApprovalForAll(owner []byte, operator []byte, permissions uint32) {}

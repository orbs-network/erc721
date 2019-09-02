package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
)

var PUBLIC = sdk.Export(
	mint, tokenMetadata,
	balanceOf, ownerOf,
	safeTransferFrom,
	approve, getApproved,
	setApprovalForAll, isApprovedForAll)

var SYSTEM = sdk.Export(_init)

func _init() {
	_setNextTokenID(0)
}

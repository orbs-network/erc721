package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
)

var PUBLIC = sdk.Export(safeTransferFrom, balanceOf, ownerOf, mint, tokenMetadata, approve, getApproved)
var SYSTEM = sdk.Export(_init)

func _init() {
	_setNextTokenID(0)
}

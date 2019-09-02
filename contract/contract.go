package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
)

var PUBLIC = sdk.Export(balanceOf, ownerOf, mint, tokenMetadata)
var SYSTEM = sdk.Export(_init)

func _init() {
	_setNextTokenID(0)
}

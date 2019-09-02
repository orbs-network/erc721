package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
)

func ownerOf(tokenId uint64) []byte {
	return state.ReadBytes(_ownerKey(tokenId))
}

func _ownerKey(tokenId uint64) []byte {
	return []byte("owner."+strconv.FormatUint(tokenId, 10))
}

func _setOwner(tokenId uint64, address []byte) {
	state.WriteBytes(_ownerKey(tokenId), address)
}
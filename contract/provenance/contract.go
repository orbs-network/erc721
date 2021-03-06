package main

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(
	onERC721Received,
	acceptTokens, rejectTokens,
	provenance,
)
var SYSTEM = sdk.Export(_init)

func _init() {

}

var CALLBACK_COUNTER = []byte("callback_counter")
var ACCEPT = []byte("accept_tokens")

var MAGIC_ON_ERC721_RECEIVED, _ = hex.DecodeString("150b7a02")

func onERC721Received(from []byte, to []byte, tokenId uint64, data []byte) []byte {
	if _shouldAcceptTokens() {
		_appendTransferEvent(from, to, tokenId, data)
		return MAGIC_ON_ERC721_RECEIVED
	}

	return nil
}

// FIXME: contract owner only
func acceptTokens() {
	state.WriteUint32(ACCEPT, 1)
}

// FIXME: contract owner only
func rejectTokens() {
	state.Clear(ACCEPT)
}

func _shouldAcceptTokens() bool {
	return state.ReadUint32(ACCEPT) == 1
}

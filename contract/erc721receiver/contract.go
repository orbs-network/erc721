package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(onERC721Received, acceptTokens, rejectTokens, receivedTokens)
var SYSTEM = sdk.Export(_init)

func _init() {

}

var CALLBACK_COUNTER = []byte("callback_counter")
var ACCEPT = []byte("accept_tokens")

func onERC721Received(operator []byte, from []byte, tokenId uint64, data []byte) []byte {
	if _shouldAcceptTokens() {
		_incCallbackCounter()
		return []byte{1, 2, 3, 4}
	}

	return nil
}

func acceptTokens() {
	state.WriteUint32(ACCEPT, 1)
}

func rejectTokens() {
	state.Clear(ACCEPT)
}

func _shouldAcceptTokens() bool {
	return state.ReadUint32(ACCEPT) == 1
}

func receivedTokens() uint64 {
	return state.ReadUint64(CALLBACK_COUNTER)
}

func _incCallbackCounter() {
	state.WriteUint64(CALLBACK_COUNTER, receivedTokens() + 1)
}
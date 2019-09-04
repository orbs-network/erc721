package main

import (
	"bytes"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var CALLBACK_CONTRACT = []byte("callback")

// FIXME: contract owner only
func setCallbackContract(name string) {
	state.WriteString(CALLBACK_CONTRACT, name)
}

func _getCallbackContract() string {
	return state.ReadString(CALLBACK_CONTRACT)
}

func _executeCallback(from []byte, to []byte, tokenId uint64, data []byte) {
	if callbackName := _getCallbackContract(); callbackName != "" {
		value := service.CallMethod(callbackName, "onERC721Received", from, to, tokenId, data)[0].([]byte)
		if !bytes.Equal(value, MAGIC_ON_ERC721_RECEIVED) {
			panic("invalid callback return value")
		}
	}
}

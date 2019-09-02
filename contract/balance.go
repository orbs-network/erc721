package main

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

func balanceOf(address []byte) uint64 {
	return state.ReadUint64(_balanceKey(address))
}

func _incBalance(address []byte) {
	if address != nil {
		balance := balanceOf(address) + 1
		state.WriteUint64(_balanceKey(address), balance)
	}
}

func _decBalance(address []byte) {
	if address != nil {
		balance := balanceOf(address) - 1
		state.WriteUint64(_balanceKey(address), balance)
	}
}

func _balanceKey(address []byte) []byte {
	return []byte("balance." + hex.EncodeToString(address))
}
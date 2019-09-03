package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var NEXT_TOKEN_ID = []byte("next_token_id")

// Returns new token id
func mint(jsonMetadata string) uint64 {
	tokenId := _getNextTokenID()

	_setMetadata(tokenId, jsonMetadata)
	_mint(address.GetCallerAddress(), tokenId)

	_setNextTokenID(tokenId + 1)
	return tokenId
}


func _mint(address []byte, tokenId uint64) {
	_transfer(nil, address, tokenId)
}

func _setNextTokenID(nextTokenId uint64) {
	state.WriteUint64(NEXT_TOKEN_ID, nextTokenId)
}

func _getNextTokenID() uint64 {
	return state.ReadUint64(NEXT_TOKEN_ID)
}


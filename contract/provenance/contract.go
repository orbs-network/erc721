package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/contract-external-libraries-go/v1/list"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/env"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
	"time"
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

type Transfer struct {
	From []byte
	To []byte
	TokenId uint64
	Timestamp uint64

	// FIXME add description
	// Description string
}

type TransferExport struct {
	From string
	To string
	TokenId uint64
	Timestamp time.Time
}

func onERC721Received(from []byte, to []byte, tokenId uint64, data []byte) []byte {
	if _shouldAcceptTokens() {
		tokenList := _tokenList(tokenId)
		tokenList.Append(Transfer{
			From: from,
			To: to,
			TokenId: tokenId,
			Timestamp: env.GetBlockTimestamp(),
		})

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

func provenance(tokenId uint64) string {
	var transfers []TransferExport

	for i := _tokenList(tokenId).Iterator(); i.Next(); {
		t := i.Value().(*Transfer)
		transfers = append(transfers, TransferExport{
			From: _encodeAddress(t.From),
			To: _encodeAddress(t.To),
			TokenId: t.TokenId,
			Timestamp: time.Unix(0, int64(t.Timestamp)),
		})
	}

	rawJSON, _ := json.Marshal(transfers)
	return string(rawJSON)
}

func _tokenList(tokenId uint64) list.List {
	return list.NewAppendOnlyList("provenance."+strconv.FormatUint(tokenId, 10), list.StructSerializer, list.StructDeserializer(Transfer{}))
}

func _encodeAddress(address []byte) string {
	return "0x" + hex.EncodeToString(address)
}
package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/contract-external-libraries-go/v1/list"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/env"
	"strconv"
	"time"
)

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

func _appendTransferEvent(from []byte, to []byte, tokenId uint64, data []byte) {
	_tokenList(tokenId).Append(Transfer{
		From: from,
		To: to,
		TokenId: tokenId,
		Timestamp: env.GetBlockTimestamp(),
	})
}

func _tokenList(tokenId uint64) list.List {
	return list.NewAppendOnlyList("provenance."+strconv.FormatUint(tokenId, 10), list.StructSerializer, list.StructDeserializer(Transfer{}))
}

func _encodeAddress(address []byte) string {
	return "0x" + hex.EncodeToString(address)
}

package main

import "bytes"

func safeTransferFrom(from []byte, to []byte, tokenId uint64) {
	_checkTransferRights(from, to, tokenId)
	_transfer(from, to, tokenId)
}

func _transfer(from []byte, to []byte, tokenId uint64) {
	_decBalance(from)
	_setOwner(tokenId, to)
	_incBalance(to)
}

func _checkTransferRights(from []byte, to []byte, tokenId uint64) {
	owner := ownerOf(tokenId)
	if !bytes.Equal(from, owner) {
		panic("transfer not authorized")
	}

	if len(to) != 40 {
		panic("transfer not authorized")
	}
}
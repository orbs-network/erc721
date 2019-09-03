package main

import "bytes"

func transferFrom(from []byte, to []byte, tokenId uint64) {
	_checkTransferRights(from, to, tokenId)
	_transfer(from, to, tokenId)
}

func safeTransferFrom(from []byte, to []byte, tokenId uint64) {
	// FIXME implement
	panic("not implemented")
}

func _transfer(from []byte, to []byte, tokenId uint64) {
	_decBalance(from)
	_setOwner(tokenId, to)
	_revokeApproval(tokenId)
	_incBalance(to)
}

func _checkTransferRights(from []byte, to []byte, tokenId uint64) {
	owner := ownerOf(tokenId)
	_checkApproval(from, tokenId)

	if !bytes.Equal(from, owner) {
		panic("transfer not authorized")
	}

	if len(to) != 20 {
		panic("transfer not authorized")
	}
}
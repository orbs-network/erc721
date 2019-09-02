package main

func _transfer(from []byte, to []byte, tokenId uint64) {
	_decBalance(from)
	_setOwner(tokenId, to)
	_incBalance(to)
}

package main

import (
"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
"strconv"
)

type List interface {
	Append(item interface{}) (length uint64)
	Get(index uint64) interface{}
	Length() uint64
	Iterator() Iterator
}

func NewAppendOnlyList(name string, serializer Serializer, deserializer Deserializer) List {
	return &list{
		name,
		serializer,
		deserializer,
	}
}

type list struct {
	name string
	serializer Serializer
	deserializer Deserializer
}

func (l *list) Append(item interface{}) (length uint64) {
	index := l.Length()
	l.serializer(l.itemKeyName(index), item)

	length = index + 1
	state.WriteUint64(l.lengthKeyName(), length)

	return
}

func (l *list) Get(index uint64) interface{} {
	key := l.itemKeyName(index)
	return l.deserializer(key)
}

func (l *list) Length() uint64 {
	return state.ReadUint64(l.lengthKeyName())
}

func (l *list) Iterator() Iterator {
	return NewListIterator(l)
}

func (l *list) itemKeyName(index uint64) []byte {
	return []byte(l.name+"."+strconv.FormatUint(index, 10))
}

func (l *list) lengthKeyName() []byte {
	return []byte(l.name+".length")
}
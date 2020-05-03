package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

type Serializer func(key []byte, item interface{})
type Deserializer func(key []byte) interface{}

func StringSerializer(key []byte, item interface{}) {
	state.WriteString(key, item.(string))
}

func StringDeserializer(key []byte) interface{} {
	return state.ReadString(key)
}

func BytesSerializer(key []byte, item interface{}) {
	state.WriteBytes(key, item.([]byte))
}

func BytesDeserializer(key []byte) interface{} {
	return state.ReadBytes(key)
}

func Uint64Serializer(key []byte, item interface{}) {
	state.WriteUint64(key, item.(uint64))
}

func Uint64Deserializer(key []byte) interface{} {
	return state.ReadUint64(key)
}

func Uint32Serializer(key []byte, item interface{}) {
	state.WriteUint32(key, item.(uint32))
}

func Uint32Deserializer(key []byte,) interface{} {
	return state.ReadUint32(key)
}


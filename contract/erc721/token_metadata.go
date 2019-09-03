package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
)

/**

ERC721 Metadata JSON Schema

{
    "title": "Asset Metadata",
    "type": "object",
    "properties": {
        "name": {
            "type": "string",
            "description": "Identifies the asset to which this NFT represents"
        },
        "description": {
            "type": "string",
            "description": "Describes the asset to which this NFT represents"
        },
        "image": {
            "type": "string",
            "description": "A URI pointing to a resource with mime type image/* representing the asset to which this NFT represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive."
        }
    }
}

 **/

func tokenMetadata(tokenId uint64) string {
	return _getMetadata(tokenId)
}

func _setMetadata(tokenId uint64, metadata string) {
	state.WriteString(_metadataKey(tokenId), metadata)
}

func _getMetadata(tokenId uint64) string {
	return state.ReadString(_metadataKey(tokenId))
}

func _metadataKey(tokenId uint64) []byte {
	return []byte("metadata." + strconv.FormatUint(tokenId, 10))
}
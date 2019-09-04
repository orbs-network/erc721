# Orbs implementation of ERC721

[erc721.org](https://erc721.org)

## Events

`Transfer(from []byte, to []byte, tokenId uint64)`

`Approval(owner []byte, address []byte, tokenId uint64)`

`ApprovalForAll(owner []byte, operator []byte, permissions uint32)`

## Methods

### Token minting and metadata

`mint(jsonMetadata string) uint64`

`tokenMetadata(tokenId uint64) string`

### ERC721

`name() string`

`symbol() string`

`balanceOf(address []byte) uint64`

`ownerOf(tokenId uint64) []byte`

`transferFrom(from []byte, to []byte, tokenId uint64)`

`safeTransferFrom(from []byte, to []byte, tokenId uint64, toContractName string, data []byte)`

`approve(approvedAddress []byte, tokenId uint64)`

`getApproved(tokenId uint64) []byte`

`setApprovalForAll(operator []byte, approved uint32)`

`isApprovedForAll(owner []byte, operator []byte) uint32`

### ERC721 Receiver

`onERC721Received(operator []byte, from []byte, tokenId uint64, data []byte) []byte`

## Testing

Some tests are a bit flaky due to a delay between transaction commit and state storage update, so you may need to run them multiple times.

```
gamma-cli start -env experimental
```

```
go test ./test/... -v
```
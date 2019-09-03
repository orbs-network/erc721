package test

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

type erc721Harness struct {
	client       *orbs.OrbsClient
	contractName string
}

func newERC721Harness() *erc721Harness {
	return &erc721Harness{
		client:       orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET),
		contractName: fmt.Sprintf("ERC721X%d", time.Now().UnixNano()),
	}
}

func (h *erc721Harness) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	fileNames, _ := ioutil.ReadDir("../contract/erc721")
	var contractSources [][]byte
	for _, fileName := range fileNames {
		source, _ := ioutil.ReadFile("../contract/erc721/" + fileName.Name())
		contractSources = append(contractSources, source)
	}

	deployTx, _, err := h.client.CreateDeployTransaction(sender.PublicKey, sender.PrivateKey,
		h.contractName, orbs.PROCESSOR_TYPE_NATIVE, contractSources...)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *erc721Harness) balanceOf(t *testing.T, sender *orbs.OrbsAccount, address []byte) uint64 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "balanceOf", address)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(uint64)
}

func (h *erc721Harness) ownerOf(t *testing.T, sender *orbs.OrbsAccount, tokenId uint64) []byte {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "ownerOf", tokenId)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].([]byte)
}

func (h *erc721Harness) tokenMetadata(t *testing.T, sender *orbs.OrbsAccount, tokenId uint64) string {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "tokenMetadata", tokenId)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(string)
}

func (h *erc721Harness) getApproved(t *testing.T, sender *orbs.OrbsAccount, tokenId uint64) []byte {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "getApproved", tokenId)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].([]byte)
}

func (h *erc721Harness) isApprovedForAll(t *testing.T, sender *orbs.OrbsAccount, address []byte, operator []byte) uint32 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "isApprovedForAll", address, operator)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(uint32)
}

func (h *erc721Harness) mint(t *testing.T, sender *orbs.OrbsAccount, jsonMetadata string) uint64 {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "mint", jsonMetadata)
	require.NoError(t, err)

	response, err := h.client.SendTransaction(tx)
	require.NoError(t, err)

	return response.OutputArguments[0].(uint64)
}

func (h *erc721Harness) transferFrom(t *testing.T, sender *orbs.OrbsAccount, from []byte, to []byte, tokenId uint64) error {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "transferFrom", from, to, tokenId)
	require.NoError(t, err)

	response, err := h.client.SendTransaction(tx)
	require.NoError(t, err)

	if response.ExecutionResult != codec.EXECUTION_RESULT_SUCCESS {
		return fmt.Errorf(response.OutputArguments[0].(string))
	}

	return nil
}

func (h *erc721Harness) safeTransferFrom(t *testing.T, sender *orbs.OrbsAccount, from []byte, to []byte, tokenId uint64) error {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "safeTransferFrom", from, to, tokenId)
	require.NoError(t, err)

	response, err := h.client.SendTransaction(tx)
	require.NoError(t, err)

	if response.ExecutionResult != codec.EXECUTION_RESULT_SUCCESS {
		return fmt.Errorf(response.OutputArguments[0].(string))
	}

	return nil
}

func (h *erc721Harness) approve(t *testing.T, sender *orbs.OrbsAccount, address []byte, tokenId uint64) error {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "approve", address, tokenId)
	require.NoError(t, err)

	response, err := h.client.SendTransaction(tx)
	require.NoError(t, err)

	if response.ExecutionResult != codec.EXECUTION_RESULT_SUCCESS {
		return fmt.Errorf(response.OutputArguments[0].(string))
	}

	return nil
}

func (h *erc721Harness) setApprovalForAll(t *testing.T, sender *orbs.OrbsAccount, operator []byte, approved uint32) error {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "setApprovalForAll", operator, approved)
	require.NoError(t, err)

	response, err := h.client.SendTransaction(tx)
	require.NoError(t, err)

	if response.ExecutionResult != codec.EXECUTION_RESULT_SUCCESS {
		return fmt.Errorf(response.OutputArguments[0].(string))
	}

	return nil
}

func paintBlackSquare(t *testing.T, h *erc721Harness, owner *orbs.OrbsAccount) uint64 {
	balance := h.balanceOf(t, owner, owner.AddressAsBytes())
	tokenId := h.mint(t, owner, `{"title":"Black Square","type":"Painting"}`)

	metadata := h.tokenMetadata(t, owner, tokenId)
	require.EqualValues(t, `{"title":"Black Square","type":"Painting"}`, metadata)
	require.EqualValues(t, balance+1, h.balanceOf(t, owner, owner.AddressAsBytes()))

	tokenOwner := h.ownerOf(t, owner, tokenId)
	require.EqualValues(t, owner.AddressAsBytes(), tokenOwner)

	return tokenId
}
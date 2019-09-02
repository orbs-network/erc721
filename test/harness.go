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

type harness struct {
	client       *orbs.OrbsClient
	contractName string
}

func newHarness() *harness {
	return &harness{
		client:       orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET),
		contractName: fmt.Sprintf("ERC721X%d", time.Now().UnixNano()),
	}
}

func (h *harness) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	fileNames, _ := ioutil.ReadDir("../contract")
	var contractSources [][]byte
	for _, fileName := range fileNames {
		source, _ := ioutil.ReadFile("../contract/" + fileName.Name())
		contractSources = append(contractSources, source)
	}

	deployTx, _, err := h.client.CreateDeployTransaction(sender.PublicKey, sender.PrivateKey,
		h.contractName, orbs.PROCESSOR_TYPE_NATIVE, contractSources...)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *harness) balanceOf(t *testing.T, sender *orbs.OrbsAccount, address []byte) uint64 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "balanceOf", address)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err, queryResponse.OutputArguments[0])

	return queryResponse.OutputArguments[0].(uint64)
}

func (h *harness) ownerOf(t *testing.T, sender *orbs.OrbsAccount, tokenId uint64) []byte {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "ownerOf", tokenId)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err, queryResponse.OutputArguments[0])

	return queryResponse.OutputArguments[0].([]byte)
}

func (h *harness) tokenMetadata(t *testing.T, sender *orbs.OrbsAccount, tokenId uint64) string {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "tokenMetadata", tokenId)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err, queryResponse.OutputArguments[0])

	return queryResponse.OutputArguments[0].(string)
}


func (h *harness) mint(t *testing.T, sender *orbs.OrbsAccount, jsonMetadata string) uint64 {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "mint", jsonMetadata)
	require.NoError(t, err)

	response, err := h.client.SendTransaction(tx)
	require.NoError(t, err, response.OutputArguments[0])

	return response.OutputArguments[0].(uint64)
}


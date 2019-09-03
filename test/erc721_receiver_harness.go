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

type erc721ReceiverHarness struct {
	client       *orbs.OrbsClient
	contractName string
}

func newErc721ReceiverHarness() *erc721ReceiverHarness {
	return &erc721ReceiverHarness{
		client:       orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET),
		contractName: fmt.Sprintf("ERC721Receiver%d", time.Now().UnixNano()),
	}
}

func (h *erc721ReceiverHarness) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	source, _ := ioutil.ReadFile("../contract/erc721receiver/contract.go")

	deployTx, _, err := h.client.CreateDeployTransaction(sender.PublicKey, sender.PrivateKey,
		h.contractName, orbs.PROCESSOR_TYPE_NATIVE, source)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *erc721ReceiverHarness) receivedTokens(t *testing.T, sender *orbs.OrbsAccount, address []byte) uint64 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "receivedTokens", address)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(uint64)
}

func (h *erc721ReceiverHarness) acceptTokens(t *testing.T, sender *orbs.OrbsAccount) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "acceptTokens")
	require.NoError(t, err)

	_, err = h.client.SendTransaction(tx)
	require.NoError(t, err)
}

func (h *erc721ReceiverHarness) rejectTokens(t *testing.T, sender *orbs.OrbsAccount) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "rejectTokens")
	require.NoError(t, err)

	_, err = h.client.SendTransaction(tx)
	require.NoError(t, err)
}
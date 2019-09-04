const { readFileSync, readdirSync } = require("fs");
const { join } = require("path");

const {
	Client, createAccount,
	PROCESSOR_TYPE_NATIVE, NetworkType
} = require("orbs-client-sdk");

function getClient() {
    const endpoint = process.env.ORBS_NODE_ADDRESS || "http://localhost:8080";
    const chain = Number(process.env.ORBS_VCHAIN) || 42;
    return new Client(endpoint, chain, NetworkType.NETWORK_TYPE_TEST_NET);
}

function getContractCode() {
	const dir = __dirname + "/../contract/erc721";
	return readdirSync(dir).map(f => readFileSync(join(dir, f)));
}

async function deploy(client, contractOwner, contractName) {
    const [tx, txid] = client.createDeployTransaction(contractOwner.publicKey, contractOwner.privateKey,
		contractName, PROCESSOR_TYPE_NATIVE, ...getContractCode());
    const receipt = await client.sendTransaction(tx);
	if (receipt.executionResult !== 'SUCCESS') {
		throw new Error(receipt.outputArguments[0].value);
	}
}

module.exports = {
	getContractCode,
	getClient,
	deploy
}

if (!module.parent) {
	(async () => {
		try {
			await deploy(getClient(), createAccount(), "ERC721")
			console.log("Deployed ERC721 smart contract successfully");
		} catch (e) {
			console.error(e);
		}
	})();
}

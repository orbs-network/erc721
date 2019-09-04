const {
	argString, argBytes, argUint64, argUint32, argAddress,
	bytesToAddress,
} = require("orbs-client-sdk");

function getErrorFromReceipt(receipt) {
    const value = receipt.outputArguments.length == 0 ? receipt.executionResult : receipt.outputArguments[0].value;
    return new Error(value);
}

class Provenance {
	constructor(orbsClient, contractName, publicKey, privateKey) {
		this.client = orbsClient;
		this.contractName = contractName;
		this.publicKey = publicKey;
		this.privateKey = privateKey;
	}

	async provenance(tokenId) {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"provenance",
			[
				argUint64(tokenId)
			]
		);

        const receipt = await this.client.sendQuery(query);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}

		return JSON.parse(receipt.outputArguments[0].value);
	}

	async acceptTokens() {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"acceptTokens",
			[]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}

    async rejectTokens() {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"rejectTokens",
			[]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}
}

module.exports = {
	Provenance
};

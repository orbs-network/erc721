const {
	argString, argBytes, argUint64, argUint32, argAddress,
	bytesToAddress,
} = require("orbs-client-sdk");

function getErrorFromReceipt(receipt) {
    const value = receipt.outputArguments.length == 0 ? receipt.executionResult : receipt.outputArguments[0].value;
    return new Error(value);
}

class ERC721 {
	constructor(orbsClient, contractName, publicKey, privateKey) {
		this.client = orbsClient;
		this.contractName = contractName;
		this.publicKey = publicKey;
		this.privateKey = privateKey;
	}

	async mint(metadata) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"mint",
			[
				argString(JSON.stringify(metadata))
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}

		return receipt.outputArguments[0].value;
	}

	async tokenMetadata(tokenId) {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"tokenMetadata",
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

	async transfer(from, to, tokenId) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"transferFrom",
			[
				argAddress(from),
				argAddress(to),
				argUint64(tokenId)
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}

	async ownerOf(tokenId) {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"ownerOf",
			[
				argUint64(tokenId)
			]
		);

		const receipt = await this.client.sendQuery(query);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}

		return bytesToAddress(receipt.outputArguments[0].value);
	}
}

module.exports = {
	ERC721
};

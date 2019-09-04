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

	async name() {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"name",
			[]
		);

		const receipt = await this.client.sendQuery(query);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}

		return receipt.outputArguments[0].value;
	}

	async symbol() {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"symbol",
			[]
		);

		const receipt = await this.client.sendQuery(query);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}

		return receipt.outputArguments[0].value;
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

	async transferFrom(fromAddress, toAddress, tokenId) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"transferFrom",
			[
				argAddress(fromAddress),
				argAddress(toAddress),
				argUint64(tokenId)
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}

	async safeTransferFrom(fromAddress, toAddress, tokenId, toContractName, bytes) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"transferFrom",
			[
				argAddress(fromAddress),
				argAddress(toAddress),
				argUint64(tokenId),
				argString(toContractName),
				argBytes(bytes)
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

	async approve(approvedAddress, tokenId) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"approve",
			[
				argAddress(approvedAddress),
				argUint64(tokenId)
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}

	async getApproved(tokenId) {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"getApproved",
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

	async setApprovalForAll(operatorAddress, approved) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"setApprovalForAll",
			[
				argAddress(operatorAddress),
				argUint32(approved ? 1 : 0)
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}

	async isApprovedForAll(ownerAddress, operatorAddress) {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"isApprovedForAll",
			[
				argAddress(ownerAddress),
				argAddress(operatorAddress)
			]
		);

		const receipt = await this.client.sendQuery(query);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}

		return receipt.outputArguments[0].value > 0;
	}

	async setCallbackContract(contractName) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"setCallbackContract",
			[
				argString(contractName),
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw getErrorFromReceipt(receipt);
		}
	}
}

module.exports = {
	ERC721
};

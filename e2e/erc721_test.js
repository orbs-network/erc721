const expect = require("expect.js");
const { createAccount } = require("orbs-client-sdk");
const { ERC721 } = require("../src/erc721");
const { deploy, getClient } = require("../src/deploy");

const blackSquare = {
	title: "Black Square",
	type: "Painting"
};

describe("ERC721", () => {
    it("updates contract state", async () => {
		const contractOwner = createAccount();
		const contractName = "ERC721" + new Date().getTime();

		await deploy(getClient(), contractOwner, contractName);

		const seller = createAccount();
		const sellerERC721 = new ERC721(getClient(), contractName, seller.publicKey, seller.privateKey);

		const tokenId = await sellerERC721.mint(blackSquare);
		expect(tokenId).to.eql(0n);

		// const buyer = createAccount();
		// const buyerERC721 = new ERC721(getClient(), contractName, seller.buyerKey, seller.buyerKey);
	});
});

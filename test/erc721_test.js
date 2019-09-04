const expect = require("expect.js");
const { createAccount } = require("orbs-client-sdk");
const { ERC721 } = require("../src/erc721");
const { deploy, getClient } = require("../src/deploy");

const blackSquare = {
	title: "Black Square",
	type: "Painting"
};

describe("ERC721", () => {
    it("transfer flow", async () => {
		const contractOwner = createAccount();
		const contractName = "ERC721" + new Date().getTime();

		await deploy(getClient(), contractOwner, contractName);

		const seller = createAccount();
		const sellerERC721 = new ERC721(getClient(), contractName, seller.publicKey, seller.privateKey);

		const tokenId = await sellerERC721.mint(blackSquare);
		expect(tokenId).to.eql(0n);

		expect(await sellerERC721.tokenMetadata(tokenId)).to.eql(blackSquare);

		const buyer = createAccount();
		await sellerERC721.transfer(seller.address, buyer.address, tokenId);
		expect(await sellerERC721.ownerOf(tokenId)).to.eql(buyer.address);
	});
});

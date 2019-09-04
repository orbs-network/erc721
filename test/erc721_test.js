const expect = require("expect.js");
const { createAccount } = require("orbs-client-sdk");
const { ERC721 } = require("../src/erc721");
const { deployERC721, getClient } = require("../src/deploy");

const blackSquare = {
	title: "Black Square",
	type: "Painting"
};

describe("ERC721", () => {
    it("transferFrom", async () => {
		const contractOwner = createAccount();
		const contractName = "ERC721" + new Date().getTime();

		await deployERC721(getClient(), contractOwner, contractName);

		const seller = createAccount();
		const sellerERC721 = new ERC721(getClient(), contractName, seller.publicKey, seller.privateKey);

		expect(await sellerERC721.name()).to.eql("ORBS721");
		expect(await sellerERC721.symbol()).to.eql("0721");

		const tokenId = await sellerERC721.mint(blackSquare);
		expect(tokenId).to.eql(0n);

		expect(await sellerERC721.tokenMetadata(tokenId)).to.eql(blackSquare);

		const buyer = createAccount();
		await sellerERC721.transferFrom(seller.address, buyer.address, tokenId);
		expect(await sellerERC721.ownerOf(tokenId)).to.eql(buyer.address);
	});

	it("transferFrom with approval per token", async () => {
		const contractOwner = createAccount();
		const contractName = "ERC721" + new Date().getTime();

		await deployERC721(getClient(), contractOwner, contractName);

		const seller = createAccount();
		const sellerERC721 = new ERC721(getClient(), contractName, seller.publicKey, seller.privateKey);

		const tokenId = await sellerERC721.mint(blackSquare);

		const approvedSeller = createAccount();
		await sellerERC721.approve(approvedSeller.address, tokenId);
		expect(await sellerERC721.getApproved(tokenId)).to.eql(approvedSeller.address);

		const buyer = createAccount();
		const approvedSellerERC721 = new ERC721(getClient(), contractName, approvedSeller.publicKey, approvedSeller.privateKey);
		await approvedSellerERC721.transferFrom(seller.address, buyer.address, tokenId);
		expect(await sellerERC721.ownerOf(tokenId)).to.eql(buyer.address);
	});

	it("transferFrom with approval from operator", async () => {
		const contractOwner = createAccount();
		const contractName = "ERC721" + new Date().getTime();

		await deployERC721(getClient(), contractOwner, contractName);

		const seller = createAccount();
		const sellerERC721 = new ERC721(getClient(), contractName, seller.publicKey, seller.privateKey);

		const tokenId = await sellerERC721.mint(blackSquare);

		const operator = createAccount();
		await sellerERC721.setApprovalForAll(operator.address, true);
		expect(await sellerERC721.isApprovedForAll(seller.address, operator.address)).to.eql(true);

		const buyer = createAccount();
		const operatorERC721 = new ERC721(getClient(), contractName, operator.publicKey, operator.privateKey);
		await operatorERC721.transferFrom(seller.address, buyer.address, tokenId);
		expect(await sellerERC721.ownerOf(tokenId)).to.eql(buyer.address);
	});
});

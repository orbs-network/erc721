const expect = require("expect.js");
const { createAccount } = require("orbs-client-sdk");
const { ERC721 } = require("../src/erc721");
const { Provenance } = require("../src/provenance");
const { deployERC721, deployProvenance, getClient } = require("../src/deploy");

const blackSquare = {
	title: "Black Square",
	type: "Painting"
};

describe("ERC721 with provenance", () => {
    it("transferFrom", async () => {
		const contractOwner = createAccount();
        const erc721ContractName = "ERC721" + new Date().getTime();
        const provenanceContractName = "Provenance" + new Date().getTime();

        await deployERC721(getClient(), contractOwner, erc721ContractName);
        await deployProvenance(getClient(), contractOwner, provenanceContractName);

        const ownerERC721 = new ERC721(getClient(), erc721ContractName, contractOwner.publicKey, contractOwner.privateKey);
        await ownerERC721.setCallbackContract(provenanceContractName);

        const ownerProvenance = new Provenance(getClient(), provenanceContractName, contractOwner.publicKey, contractOwner.privateKey);
        await ownerProvenance.acceptTokens();

		const seller = createAccount();
		const sellerERC721 = new ERC721(getClient(), erc721ContractName, seller.publicKey, seller.privateKey);
		
		const tokenId = await sellerERC721.mint(blackSquare);
		expect(tokenId).to.eql(0n);

		expect(await sellerERC721.tokenMetadata(tokenId)).to.eql(blackSquare);

		const buyer = createAccount();
		await sellerERC721.transferFrom(seller.address, buyer.address, tokenId);
        expect(await sellerERC721.ownerOf(tokenId)).to.eql(buyer.address);

        const provenance = await ownerProvenance.provenance(tokenId);
        expect(provenance[0].From).to.eql(seller.address.toLowerCase());
        expect(provenance[0].To).to.eql(buyer.address.toLowerCase());
        expect(provenance[0].TokenId).to.eql(0n);
        expect(provenance[0].Timestamp).to.be.a("string");
	});

});

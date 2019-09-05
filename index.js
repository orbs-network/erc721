const { ERC721 } = require("./src/erc721");
const { Provenance } = require("./src/provenance");
const { deployERC721, deployProvenance } = require("./src/deploy");

module.exports = {
    ERC721,
    Provenance,
    deployERC721,
    deployProvenance
}
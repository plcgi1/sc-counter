// Help Truffle find `TruffleTutorial.sol` in the `/contracts` directory
const AddressCallerCounter = artifacts.require("AddressCallerCounter");

const creator = '0x47FD271c5C2b8431061A789DeEC503A9c2c8F2b5'

module.exports = function(deployer) {
  // Command Truffle to deploy the Smart Contract
  deployer.deploy(AddressCallerCounter);
};
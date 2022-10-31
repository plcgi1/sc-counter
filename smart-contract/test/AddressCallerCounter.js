const { assert } = require("chai")
const AddressCallerCounter = artifacts.require("./AddressCallerCounter.sol")

require("chai")
  .use(require("chai-as-promised"))
  .should()

contract('AddressCallerCounter', ([contractOwner, secondAddress, thirdAddress]) => {
  let addressCallerCounter

  // this would attach the deployed smart contract and its methods
  // to the `truffleTutorial` variable before all other tests are run
  before(async () => {
    addressCallerCounter = await AddressCallerCounter.deployed()
  })

  // check if deployment goes smooth
  describe('deployment', () => {
    // check if the smart contract is deployed
    // by checking the address of the smart contract
    it('deploys successfully', async () => {
      const address = await addressCallerCounter.address

      // assert.notEqual(address, '')
      assert.notEqual(address, undefined)
      assert.notEqual(address, null)
      assert.notEqual(address, 0x0)
    })
  })

  describe('incAddressCounter', () => {
    it('caller registered', async () => {
      // set new message
      await addressCallerCounter.incAddressCounter({ from: contractOwner })

      const owner = await addressCallerCounter.owner()

      // `from` helps us identify by any address in the test
      let res = await addressCallerCounter.getCountByAddress(contractOwner)
      //
      assert.equal(res, 1)

      await addressCallerCounter.incAddressCounter({ from: contractOwner })
      res = await addressCallerCounter.getCountByAddress(contractOwner)
      assert.equal(res, 2)
      //
      await addressCallerCounter.incAddressCounter({ from: secondAddress })
      res = await addressCallerCounter.getCountByAddress(secondAddress)
      assert.equal(res, 1)
    })

    it('address that is not the owner fails to set a message', async () => {
      await addressCallerCounter.getCountByAddress(contractOwner, { from: secondAddress })
        .should.be.rejected
      await addressCallerCounter.getCountByAddress(contractOwner, { from: thirdAddress })
        .should.be.rejected
    })

    it('address owner', async () => {
      const owner = await addressCallerCounter.owner()
      assert.equal(owner, contractOwner)
    })
  })
})
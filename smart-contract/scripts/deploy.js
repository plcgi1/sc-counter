async function main() {
  const AddressCallerCounter = await ethers.getContractFactory("AddressCallerCounter");

  // Start deployment, returning a promise that resolves to a contract object
  const addressCallerCounter = await AddressCallerCounter.deploy();
  console.log("Contract deployed to address:", addressCallerCounter.address);
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
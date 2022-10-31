network=goerli
address=0xFF40bC796f41a5e521bECA3DCe672d16eB905520

[install](https://www.web3.university/tracks/create-a-smart-contract/deploy-your-first-smart-contract)

## Truffle usage
 
1. Install truffle
2. Run truffle
3. Deploy contract
```
cd smart-contract
truffle migrate --reset
```
4. Copy file smart contract smart-contract/build/contracts/AddressCallerCounter.json to 
    frontend/src/contracts/AddressCallerCounter.json
5. ```cd go-listener && make start```
6. ```cd frontend && yarn start```
7. Add truffle users to metamask    
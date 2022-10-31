const { usePlugin } = require('@nomiclabs/buidler/config');

usePlugin('@nomiclabs/buidler-waffle');
usePlugin('solidity-coverage');
usePlugin('buidler-deploy');

const ethers = require('ethers');
const fs = require('fs');

const buidlerConfig = {
  solc: {
    version: '0.6.10',
  },
  paths: {
    sources: './contracts',
    deploy: 'deploy',
    deployments: 'deployments',
  },
};

let deplymentConfigs = {};
if (process.env.NODE_ENV !== 'test') {
  require('./scripts/deployConfigs')(); // eslint-disable-line global-require
  const configs = require('./configs')(); // eslint-disable-line global-require

  if (!configs.DEPLOYMENT_ACCOUNT_PK) {
    if (!configs.KEYSTORE || !configs.KEYSTORE_PASSWORD) {
      throw new Error(
        'Keystore file needed. Add KEYSTORE and KEYSTORE_PASSWORD to your configs.json file.'
      );
    }
    const keystore = fs.readFileSync(configs.KEYSTORE);
    const address = ethers.utils.getJsonWalletAddress(keystore.toString());
    console.log(`Found keystore at ${configs.KEYSTORE}! Address: ${address}`);
    const wallet = ethers.Wallet.fromEncryptedJsonSync(
      keystore.toString(),
      configs.KEYSTORE_PASSWORD
    );
    console.log('Got private key successfully!');
    configs.DEPLOYMENT_ACCOUNT_PK = wallet.privateKey;
  } else {
    console.log('Using private key specified in configs.json');
  }

  deplymentConfigs = {
    networks: {
      rinkeby: {
        url: `https://rinkeby.infura.io/v3/${configs.INFURA_PROJECT_ID}`,
        accounts: [configs.DEPLOYMENT_ACCOUNT_PK],
        gas: 4000000,
        gasPrice: configs.GAS_PRICE,
      },
      ropsten: {
        url: `https://ropsten.infura.io/v3/${configs.INFURA_PROJECT_ID}`,
        accounts: [configs.DEPLOYMENT_ACCOUNT_PK],
        gas: 4000000,
        gasPrice: configs.GAS_PRICE,
      },
      mainnet: {
        url: `https://mainnet.infura.io/v3/${configs.INFURA_PROJECT_ID}`,
        accounts: [configs.DEPLOYMENT_ACCOUNT_PK],
        gas: 4000000,
        gasPrice: configs.GAS_PRICE,
      },
    },
    namedAccounts: {
      deployer: {
        1: 0,
        3: 0, // ropsten
        4: 0, // rinkeby
      },
      minter: {
        1: configs.MINTER,
        3: configs.MINTER,
        4: configs.MINTER,
      },
      wiper: {
        1: configs.WIPER,
        3: configs.WIPER,
        4: configs.WIPER,
      },
      registryManager: {
        1: configs.REGISTRY_MANAGER,
        3: configs.REGISTRY_MANAGER,
        4: configs.REGISTRY_MANAGER,
      },
    },
  };
}

module.exports = {
  ...buidlerConfig,
  ...deplymentConfigs,
};

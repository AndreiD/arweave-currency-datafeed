# :moneybag: Arweave Currency Datafeed 

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/AndreiD/arweave-ipfs-bridge/blob/master/LICENSE)

currency rates, on Arweave

### Features

- hourly updates
- uses https://openexchangerates.org

### How to use it

Since it updates hourly, until the balance of the wallet is 0,
I've added a tag called "TIME-HOUR"


People should query for the latest file like this

~~~~
{
  op: 'and',
  expr1: {
    op: 'equals',
    expr1: 'from',
    expr2: 'hjCuCtTDySi9_KNtRxmPIsAyYVrTXeJZNP3R1aHfNPg'
  },
  expr2: {
    op: 'equals',
    expr1: 'TIME-HOUR',
    expr2: '2019-11-17T15'
  }
}
~~~~

TIME-HOUR represents the first part of the ISO8601 time formatting (GMT)
so get the unix time, transform it to ISO8601, and leave just the hours

**If you think of a better way to do this, let me know**

#### Configuration file

~~~~
{
  "debug": true,
  "nodeURL": "https://arweave.net",
  "appID": "YOUR_OPENEXCHANGERATES_ORG_APPID_HERE",
  "walletFile": "/PATH/TO/YOUR/arweave-wallet.json",
  "server": {
    "host": "localhost",
    "port": 5555
  }
}
~~~~

Where:

- **debug** run with move verbosity
- **nodeURL** the arweave node url
- **appID** the https://openexchangerates.org APP ID
- **walletFile** the location of the wallet.json file
- **server** ip & port where this service should run

### Special thanks to:

https://github.com/Dev43/arweave-go -> for the transaction signing & transmitting code

### Bugs / Features / Questions

fell free to create an issue

### TODO://

- [x] tags
- [ ] awaiting your idea


## License

:moneybag: Arweave Currency Datafeed  is released under the MIT license.
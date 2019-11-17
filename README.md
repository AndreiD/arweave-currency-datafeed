
<p align="center">
  <h3 align="center">Arweave Currency Datafeed</h3>
  <p align="center">currency rates, on Arweave</p>
</p>

---
# :moneybag:

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/AndreiD/arweave-ipfs-bridge/blob/master/LICENSE)

currency rates, on Arweave

### Features

- hourly updates
- uses https://openexchangerates.org
- 193 currencies (it includes crypto) check them here <a href="./currencies.json">currencies.json</a>

### Get the currency rates API

People should query using ARQL for the latest file like this

Tip: you can use https://ibwslcweo3rb.arweave.net/1w40L6Z8WLrRTZiBRcp4qMPaMuJfblUlB9pqovkh5PQ 
if you want to quickly test it

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

Since it updates hourly, I've added a tag called "TIME-HOUR"

TIME-HOUR represents the first part of the ISO8601 time formatting (UTC)
so get the unix time, transform it to ISO8601, and leave just the date & hours

**If you think of a better way to do this, let me know**

#### Use it as a template

If you want to use this as a template for your project, for whatever reason, it's simple

- edit the config file
- check the logic under sync.go

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

## Sponsors

Does your company use arweave-datafeed ? Help keep the project bug-free and feature rich by sending a small donation


## License

:moneybag: Arweave Currency Datafeed  is released under the MIT license.
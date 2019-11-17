# :moneybag: Arweave Currency Datafeed 

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/AndreiD/arweave-ipfs-bridge/blob/master/LICENSE)

currency rates, on Arweave

### Features

- hourly updates

### How to use it

xxxxx

#### Configuration file

~~~~
{
  "nodeURL": "https://arweave.net",
  "fixerAPIKey": "a8b0135ed1610852099fb49273e08946",
  "intervalSeconds": "3600"
}
~~~~

Where:

- **nodeURL** the arweave node url
- **fixerAPIKey** the Fixer API Key
- **intervalSeconds** how fast the boot should run (default: 1hour)


### Special thanks to:

https://github.com/Dev43/arweave-go -> for the transaction signing & transmitting code

### Bugs / Features / Questions

fell free to create an issue

### TODO://

- [x] compression
- [x] docker
- [x] golang-ci yml
- [ ] awaiting your idea


## License

AIB is released under the MIT license.
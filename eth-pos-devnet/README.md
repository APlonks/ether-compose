Project based on https://github.com/OffchainLabs/eth-pos-devnet/tree/master

# Ethereum Proof-of-Stake Devnet

This repository provides a docker-compose file to run a fully-functional, local development network for Ethereum with proof-of-stake enabled. This configuration uses [Prysm](https://github.com/prysmaticlabs/prysm) as a consensus client and [go-ethereum](https://github.com/ethereum/go-ethereum) for execution. **It starts from proof-of-stake** and does not go through the Ethereum merge.

This sets up a single node development network with 64 deterministically-generated validator keys to drive the creation of blocks in an Ethereum proof-of-stake chain. Here's how it works:

1. We initialize a go-ethereum, proof-of-work development node from a genesis config
2. We initialize a Prysm beacon chain, proof-of-stake development node from a genesis config

The development net is fully functional and allows for the deployment of smart contracts and all the features that also come with the Prysm consensus client such as its rich set of APIs for retrieving data from the blockchain. This development net is a great way to understand the internals of Ethereum proof-of-stake and to mess around with the different settings that make the system possible.

## Installation Blockchain

### Minimum prerequisites : 

Materials:
- 2 CPU
- 4G MEMORY

Software:
- Install Docker

### Installation

```
git clone https://github.com/aplonks/Ethercompose.git && cd Ethercompose/eth-pos-devnet
# Configure the .env file by adding the IP of your device/virtual machine
sudo ./clean.sh
```
Then run Private Ethereum Blockchain using POS (Proof of Stake) as consensus algorithm
```
docker compose up --profile pos -d # for pos
```
or run Private Ethereum Blockchain using POA (Proof of Authority) as consensus algorithm
```
docker compose up --profile pos -d # for pos
```

You will see the following:

```
$ docker compose up -d
[+] Running 11/15
 ⠸ Network bcnetwork                                                     Created          2.3s
 ⠸ Volume "eth-pos-devnet_portainer_data"                                Created          2.3s
 ⠸ Volume "eth-pos-devnet_prom_data"                                     Created          2.3s
 ⠸ Volume "eth-pos-devnet_grafana_storage"                               Created          2.3s
 ✔ Container eth-pos-devnet-lite-explorer-1                              Started          0.6s
 ✔ Container eth-pos-devnet-weavescope-1                                 Started          0.3s
 ✔ Container eth-pos-devnet-grafana-1                                    Started          0.7s
 ✔ Container eth-pos-devnet-portainer-1                                  Started          0.5s
 ✔ Container eth-pos-devnet-prometheus-1                                 Started          0.7s
 ✔ Container eth-pos-devnet-create-beacon-chain-genesis-1                Exited           0.5s
 ! weavescope Published ports are discarded when using host network mode                  0.0s
 ✔ Container eth-pos-devnet-geth-genesis-1                               Exited           1.2s
 ✔ Container eth-pos-devnet-beacon-chain-1                               Started          1.4s
 ✔ Container eth-pos-devnet-validator-1                                  Started          1.6s
 ✔ Container eth-pos-devnet-geth-1                                       Started          2.1s
```

Each time you restart, you can wipe the old data using `./clean.sh`.

Next, you can inspect the logs of the different services launched. 

```
docker logs eth-pos-devnet-geth-1 -f
```

and see:

```
INFO [08-19|00:44:30.956] Imported new potential chain segment     blocks=1 txs=0 mgas=0.000 elapsed=1.356ms     mgasps=0.000 number=50 hash=e0bd7f..497d27 dirty=0.00B
INFO [08-19|00:44:31.030] Chain head was updated                   number=50 hash=e0bd7f..497d27 root=815538..801014 elapsed=1.49025ms
INFO [08-19|00:44:35.215] Imported new potential chain segment     blocks=1 txs=0 mgas=0.000 elapsed=3.243ms     mgasps=0.000 number=51 hash=a5fb7c..5e844b dirty=0.00B
INFO [08-19|00:44:35.311] Chain head was updated                   number=51 hash=a5fb7c..5e844b root=815538..801014 elapsed=1.73475ms
INFO [08-19|00:44:39.435] Imported new potential chain segment     blocks=1 txs=0 mgas=0.000 elapsed=1.355ms     mgasps=0.000 number=52 hash=b2fd97..22e230 dirty=0.00B
INFO [08-19|00:44:39.544] Chain head was updated                   number=52 hash=b2fd97..22e230 root=815538..801014 elapsed=1.167959ms
INFO [08-19|00:44:42.733] Imported new potential chain segment     blocks=1 txs=0 mgas=0.000 elapsed=2.453ms     mgasps=0.000 number=53 hash=ee046e..e56b0c dirty=0.00B
INFO [08-19|00:44:42.747] Chain head was updated                   number=53 hash=ee046e..e56b0c root=815538..801014 elapsed="821.084µs"
```

### Configure Prometheus & Grafana
- Prometheus Grafana configuration : 
https://grafana.com/grafana/dashboards/18463-go-ethereum-by-instance/
- Import data dashboard template in grafana : https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/import-dashboards/

### Blockchain removal
Delete the blockchain 
```
sudo ./clean.sh
```

Delete the blockchain and all the volumes
```
sudo ./clean.sh -f
```

# Available Features

- Starts from the Capella Ethereum hard fork
- The network launches with a [Validator Deposit Contract](https://github.com/ethereum/consensus-specs/blob/dev/solidity_deposit_contract/deposit_contract.sol) deployed at address `0x4242424242424242424242424242424242424242`. This can be used to onboard new validators into the network by depositing 32 ETH into the contract
- The default account used in the go-ethereum node is address `0x123463a4b065722e99115d6c222f267d9cabb524` which comes seeded with ETH for use in the network. This can be used to send transactions, deploy contracts, and more
- The default account, `0x123463a4b065722e99115d6c222f267d9cabb524` is also set as the fee recipient for transaction fees proposed validators in Prysm. This address will be receiving the fees of all proposer activity
- The go-ethereum JSON-RPC API is available at http://geth:8545
- The Prysm client's REST APIs are available at http://beacon-chain:3500. For more info on what these APIs are, see [here](https://ethereum.github.io/beacon-APIs/)
- The Prysm client also exposes a gRPC API at http://beacon-chain:4000

# Listening Ports

### Ethereum execution client go : **GETH**
|Ports|Process|
|:--|:--|
|8545|HTTP endpoint|
|8546|Websocket endpoint|
|8551|Authentication port for the consensus client|
|6060|Metrics Port|

### Ethereum consensus client : **Prysm** (exists only in POS)
ref : https://docs.prylabs.network/docs/prysm-usage/p2p-host-ip
|Ports|Process|
|:--|:--|
|4000|Port to let the validator access to the consensus using gRPC|
|3500|JSON-RPC for API|


### Ethereum consensus client : **Validator** (exists only in POS)
|Ports|Process|
|:--|:--|


### Block explorer : **Ethereum lite-explorer**
|Ports|Process|
|:--|:--|
|8081|Web Server|

### Metrcis Collector : **Prometheus**
|Ports|Process|
|:--|:--|
|9090|WebServer & endpoint for grafana|

### Graph & Stats on metrics : **Grafana**
|Ports|Process|
|:--|:--|
|3000|WebServer|

### Containers Manager : **Portainer**
|Ports|Process|
|:--|:--|
|8000|WebServer|

### Containers/process visualisation : **weaveworks**
|Ports|Process|
|:--|:--|
|4040|WebServer|

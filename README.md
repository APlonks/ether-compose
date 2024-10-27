![](/images/Image_README_ether-compose.png)
Project based on https://github.com/OffchainLabs/eth-pos-devnet/tree/master

# 1. Ethereum Proof-of-Stake Devnet

- [1. Ethereum Proof-of-Stake Devnet](#1-ethereum-proof-of-stake-devnet)
  - [1.1. Minimum prerequisites:](#11-minimum-prerequisites)
  - [1.2. Installation](#12-installation)
  - [1.3. IMPORTANT NOTES](#13-important-notes)
  - [1.4. Blockchain removal](#14-blockchain-removal)
- [2. Available Features](#2-available-features)
- [3. Listening Ports](#3-listening-ports)
  - [3.1. Ethereum execution client go: **GETH**](#31-ethereum-execution-client-go-geth)
  - [3.2. Ethereum consensus client: **Prysm**](#32-ethereum-consensus-client-prysm)
  - [3.3. Ethereum consensus client: **Validator**](#33-ethereum-consensus-client-validator)
  - [3.4. Block explorer: **Ethereum lite-explorer**](#34-block-explorer-ethereum-lite-explorer)
  - [3.5. Faucet: **ether-faucet**](#35-faucet-ether-faucet)


This repository provides a docker-compose file to run a fully-functional, local development network for Ethereum with proof-of-stake enabled. This configuration uses [Prysm](https://github.com/prysmaticlabs/prysm) as a consensus client and [go-ethereum](https://github.com/ethereum/go-ethereum) for execution. **It starts from proof-of-stake** and does not go through the Ethereum merge.

This sets up a single node development network with 64 deterministically-generated validator keys to drive the creation of blocks in an Ethereum proof-of-stake chain. Here's how it works:

1. We initialize a go-ethereum, proof-of-work development node from a genesis config
2. We initialize a Prysm beacon chain, proof-of-stake development node from a genesis config

The development net is fully functional and allows for the deployment of smart contracts and all the features that also come with the Prysm consensus client such as its rich set of APIs for retrieving data from the blockchain. This development net is a great way to understand the internals of Ethereum proof-of-stake and to mess around with the different settings that make the system possible.

## 1.1. Minimum prerequisites:

Materials:
- 2 CPU cores
- 4 Go RAM

Software:
- Install Docker (recommended version latest) : [Install Docker Engine](https://docs.docker.com/engine/install/)

## 1.2. Installation

Clone the project
```
git clone https://github.com/APlonks/ether-compose.git && cd ether-compose
```

Configure the .env file by adding the IP of your device/virtual machine.
This adress will be used by the explorer to access 
```
IP_SERVER=<server_ip>
# Example :
# IP_SERVER=192.168.1.45
```

Clean your environnement
```
sudo ./clean.sh
```

Create blockchain network called bcnetwork if it doesn't exist.
```
docker network create --driver bridge bcnetwork
```

Then run Private Ethereum Blockchain using POS (Proof of Stake) as consensus algorithm.
```
docker compose up -d
```

You can add profile : 
- explorer : to deploy [Ethereum Lite Explorer by Alethio](https://github.com/Alethio/ethereum-lite-explorer)
- ether-faucet : to deploy [ether-faucet by APlonks](https://github.com/APlonks/ether-faucet)
```
docker compose --profile explorer --profile ether-faucet up -d
```


The configuration can be found further in this documentation.

You will see the following:

```bash
$ docker compose --profile explorer --profile ether-faucet up -d
[+] Running 9/9
 ✔ Container ether-compose-lite-explorer-1                Started                                0.6s 
 ✔ Container ether-faucet-backend                         St...                                  0.5s 
 ✔ Container ether-faucet-frontend                        S...                                   0.5s 
 ✔ Container redis-stack                                  Started                                0.6s 
 ✔ Container ether-compose-create-beacon-chain-genesis-1  Exited                                 0.9s 
 ✔ Container geth-genesis-pos                             Exited                                 1.7s 
 ✔ Container ether-compose-beacon-chain-1                 Started                                1.2s 
 ✔ Container geth                                         Started                                1.9s 
 ✔ Container ether-compose-validator-1                    Started                                1.4s
```

Each time you restart, you can wipe the old data using `./clean.sh`.

Next, you can inspect the logs of the different services launched. 

```
docker logs geth -f
```

and see:

```
INFO [10-27|08:00:31.469] Imported new potential chain segment     number=1 hash=8654d2..2105c7 blocks=1 txs=0 mgas=0.000 elapsed=2.304ms     mgasps=0.000 triedirty=0.00B
INFO [10-27|08:00:31.552] Chain head was updated                   number=1 hash=8654d2..2105c7 root=cf6120..825bd5 elapsed=2.40845ms
ERROR[10-27|08:00:31.552] Nil finalized block cannot evict old blobs
INFO [10-27|08:00:31.555] Indexed transactions                     blocks=2 txs=0 tail=0 elapsed=2.319ms
INFO [10-27|08:00:31.559] Starting work on payload                 id=0x03f858e7eb51aa3c
INFO [10-27|08:00:31.559] Updated payload                          id=0x03f858e7eb51aa3c number=2 hash=0df6ea..6e4149 txs=0 withdrawals=0 gas=0 fees=0 root=cf6120..825bd5 elapsed="62.378µs"
INFO [10-27|08:00:43.269] Stopping work on payload                 id=0x03f858e7eb51aa3c reason=delivery
INFO [10-27|08:00:43.283] Imported new potential chain segment     number=2 hash=0df6ea..6e4149 blocks=1 txs=0 mgas=0.000 elapsed=2.134ms     mgasps=0.000 triedirty=0.00B
INFO [10-27|08:00:43.324] Chain head was updated                   number=2 hash=0df6ea..6e4149 root=cf6120..825bd5 elapsed=2.033147ms
ERROR[10-27|08:00:43.324] Nil finalized block cannot evict old blobs
INFO [10-27|08:00:43.327] Starting work on payload                 id=0x030e91dc81563803
INFO [10-27|08:00:43.328] Updated payload                          id=0x030e91dc81563803 number=3 hash=91760b..310628 txs=0 withdrawals=0 gas=0 fees=0 root=cf6120..825bd5 elapsed="74.611µs"
INFO [10-27|08:00:55.284] Stopping work on payloasd                 id=0x030e91dc81563803 reason=delivery
INFO [10-27|08:00:55.297] Imported new potential chain segment     number=3 hash=91760b..310628 blocks=1 txs=0 mgas=0.000 elapsed=2.026ms     mgasps=0.000 triedirty=0.00B
```

## 1.3. IMPORTANT NOTES
Up to the 24th block, no block is finalized, so the following error will be visible in Geth's logs:
```log
ERROR[10-27|00:20:21.329] Nil finalized block cannot evict old blobs
```
This is an expected error, so don't worry, it will be removed after the 24th block, which will be the first finalized block.


## 1.4. Blockchain removal
Delete the blockchain 
```
sudo ./clean.sh
```

# 2. Available Features

- Starts from the Deneb Ethereum hard fork
- The network launches with a [Validator Deposit Contract](https://github.com/ethereum/consensus-specs/blob/dev/solidity_deposit_contract/deposit_contract.sol) deployed at address `0x4242424242424242424242424242424242424242`. This can be used to onboard new validators into the network by depositing 32 ETH into the contract
- The default account used in the go-ethereum node is address `0x123463a4b065722e99115d6c222f267d9cabb524` which comes seeded with ETH for use in the network. This can be used to send transactions, deploy contracts, and more
- The default account, `0x123463a4b065722e99115d6c222f267d9cabb524` is also set as the fee recipient for transaction fees proposed validators in Prysm. This address will be receiving the fees of all proposer activity
- The go-ethereum JSON-RPC API is available at http://geth:8545
- The Prysm client's REST APIs are available at http://beacon-chain:3500. For more info on what these APIs are, see [here](https://ethereum.github.io/beacon-APIs/)
- The Prysm client also exposes a gRPC API at http://beacon-chain:4000

# 3. Listening Ports

## 3.1. Ethereum execution client go: **GETH**
| Ports | Process                                      |
| :---- | :------------------------------------------- |
| 8545  | HTTP endpoint                                |
| 8546  | Websocket endpoint                           |
| 8551  | Authentication port for the consensus client |
| 6060  | Metrics Port                                 |

## 3.2. Ethereum consensus client: **Prysm**
ref : https://docs.prylabs.network/docs/prysm-usage/p2p-host-ip
| Ports | Process                                                      |
| :---- | :----------------------------------------------------------- |
| 4000  | Port to let the validator access to the consensus using gRPC |
| 3500  | JSON-RPC for API                                             |


## 3.3. Ethereum consensus client: **Validator**
| Ports | Process |
| :---- | :------ |


## 3.4. Block explorer: **Ethereum lite-explorer**
| Ports | Process    |
| :---- | :--------- |
| 8081  | Web Server |

## 3.5. Faucet: **ether-faucet**
| Ports | Process                |
| :---- | :--------------------- |
| 5001  | Frontend web interface |
| 5002  | Backend API            |
| 8001  | Redis web interface    |

For ether-faucet, the configuration to enable the frontend to access the ether-faucet API is carried out on the web interface. 

For more information, see : [ether-faucet by APlonks](https://github.com/APlonks/ether-faucet)
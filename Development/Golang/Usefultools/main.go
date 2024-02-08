package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"tools/functions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	var (
		client                *ethclient.Client
		err                   error
		choice                int
		http_endpoint         string
		rich_account_priv_key string
		richPrivKey           *ecdsa.PrivateKey
		richPubKey            common.Address
	)

	http_endpoint = "http://localhost:8545"
	rich_account_priv_key = "2e0834786285daccd064ca17f1654f67b4aef298acbb82cef9ec422fb4975622"

	richPrivKey, richPubKey, err = functions.RetrieveKeysFromHexHashedPrivateKey(rich_account_priv_key)
	if err != nil {
		log.Fatal("Cannot retrieve Private and Public keys")
	}

	client, err = ethclient.Dial(http_endpoint)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Choose what do u want to do:")
		fmt.Println("1: Create a new account")
		fmt.Println("2: Retrieve information header about a block")
		fmt.Println("3: Retrieve compete information about a block")
		fmt.Println("4: Send Ethers from a rich account to an account")

		fmt.Println()
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("Create a new account")
			functions.CreateWallet()
		case 2:
			fmt.Println("Retrieve information header about a block")
			functions.Blockheader(client)
		case 3:
			fmt.Println("Retrieve compete information about a block")
			functions.Blockfull((client))
		case 4:
			fmt.Println("Send Ethers from a rich account to an account")
			functions.SendEthers(client, richPrivKey, richPubKey)
		default:
			fmt.Println("Functions not implemented")

		}
		// functions.Blockheader(client)
	}

}

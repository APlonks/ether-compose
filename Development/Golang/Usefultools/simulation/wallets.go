package simulation

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"tools/functions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	Key        ecdsa.PrivateKey
	KeyHex     string
	Address    common.Address
	AddressHex string
}

func NewWallets(key *ecdsa.PrivateKey, keyHex string, address common.Address, addressHex string) Wallet {
	wallet := Wallet{Key: *key, KeyHex: keyHex, Address: address, AddressHex: addressHex}
	return wallet
}

func CreateWallets(numWallets int) []Wallet {

	wallets := make([]Wallet, numWallets)

	for i := 0; i < numWallets; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		// fmt.Println("The private key:", hexutil.Encode(privateKeyBytes))

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

		address := crypto.PubkeyToAddress(*publicKeyECDSA)

		hash := sha3.NewLegacyKeccak256()
		hash.Write(publicKeyBytes[1:])
		// fmt.Println("The public key:", hexutil.Encode(hash.Sum(nil)[12:]))

		wallets[i] = NewWallets(privateKey, hexutil.Encode(privateKeyBytes), address, hexutil.Encode(hash.Sum(nil)[12:]))
	}

	return wallets
}

func SendEthers(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, wallets []Wallet, nbEthers int) {

	var (
		nonce uint64
		err   error
	)

	for _, wallet := range wallets {
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		functions.ErrManagement(err)

		// Convert nbEthers (int) en big.Int
		amount := big.NewInt(int64(nbEthers))
		// Convert Ethers to Wei (1 Ether = 1e18 Wei)
		weiValue := new(big.Int).Mul(amount, big.NewInt(1e18))

		value := weiValue         // in wei (1 eth)
		gasLimit := uint64(21000) // in units
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var data []byte
		tx := types.NewTransaction(nonce, wallet.Address, value, gasLimit, gasPrice, data)
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
		fmt.Println()
	}

}

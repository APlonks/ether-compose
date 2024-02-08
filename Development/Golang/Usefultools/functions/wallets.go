package functions

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"hash"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func CreateWallet() {

	var (
		privateKey     *ecdsa.PrivateKey
		err            error
		publicKeyECDSA *ecdsa.PublicKey
		ok             bool
		publicKeyBytes []byte
		address        string
		hash           hash.Hash
	)

	fmt.Println()
	privateKey, err = crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Private key in *ecdsa.PrivateKey format", privateKey)
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private key in hex format:", hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok = publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes = crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println("EDCSA public key in Slice of bytes",hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Public key in common.Address (hex) format:", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash = sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Public key in hex format hashed:", hexutil.Encode(hash.Sum(nil)[12:]))

	// Providing 10 Ethers to the account
	fmt.Println()
}

func SendEthers(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address) {

	var (
		toaddr   string
		err      error
		nbEthers int
	)

	fmt.Println()

	fmt.Print("Which account to send ethers ?:")
	_, err = fmt.Scanf("%s", &toaddr)
	if err != nil {
		log.Fatal("Cannot read the address which will receive ethers")
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("How much ethers to send ?")
	_, err = fmt.Scanf("%d", &nbEthers)

	nbEthers = nbEthers * 1000000000000000000

	value := big.NewInt(int64(nbEthers)) // in wei (1 eth)
	gasLimit := uint64(21000)            // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(toaddr)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

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
	fmt.Println()
}

func RetrieveKeysFromHexHashedPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {

	fmt.Println()

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, fromAddress, nil
}

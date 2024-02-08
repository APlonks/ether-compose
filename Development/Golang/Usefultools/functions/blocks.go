package functions

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Blockheader(client *ethclient.Client) {

	var (
		number *big.Int
		input  int
	)

	fmt.Print("Which block ? : ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée :", err)
		return
	}

	number = big.NewInt(int64(input))

	header, err := client.HeaderByNumber(context.Background(), number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The block number :", header.Number.String()) // 5671744
	fmt.Println("The block hash", header.Hash().Hex())
	fmt.Println("The block size", header.Size())
}

func Blockfull(client *ethclient.Client) {

	var (
		number *big.Int
		input  int
	)

	fmt.Print("Which block ? : ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée :", err)
		return
	}

	number = big.NewInt(int64(input))

	block, err := client.BlockByNumber(context.Background(), number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("The block number :", block.Header().Number) // 5671744
	fmt.Println("The block hash", block.Header().Hash().Hex())
	fmt.Println("The block size", block.Header().Size())

	fmt.Println("Transactions list:", block.Transactions())

	fmt.Println()
}

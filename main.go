package main

import (
	"fmt"

	"github.com/pkw5240/blockchain/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}

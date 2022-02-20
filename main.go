package main

import (
	"fmt"

	"github.com/pkw5240/blockchain/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("second")
	chain.AddBlock("third")

	fmt.Println(chain)

	for _, block := range chain.AllBlocks() {
		fmt.Printf("data : %s \n", block.Data)
		fmt.Printf("hash : %x \n", block.Hash)
		fmt.Printf("PrevHash : %x \n", block.PrevHash)
	}

}

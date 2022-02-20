package main

import (
	"fmt"

	"github.com/pkw5240/blockchain/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}

/*
type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLashHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.prevHash = fmt.Sprintf("%x", hash)

	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) getLashHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Println(block)
	}
}

func main() {

	chain := blockchain{}
	chain.addBlock("genesis block")
	chain.addBlock("second block")
	chain.listBlocks()

	// genesisBlock := block{"Genesis block", "", ""}

	// for _, aByte := range "Genesis block" {
	// 	fmt.Printf("%b", aByte)
	// }
	// hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	// hexHash := fmt.Sprintf("%x", hash)
	// genesisBlock.hash = hexHash

	// // fmt.Println(genesisBlock)

}*/

// package main

// import (
// 	"fmt"

// 	"github.com/pkw5240/blockchain/person"
// )

// func main() {

// 	name := "12"
// 	fmt.Print(name)
// 	fmt.Print("welcome to giwan coin")

// 	fmt.Print(test(3, 4, "Test!@@"))

// 	fmt.Print(testArgu(1, 2, 3, 4, 5, 6, 7))

// 	printWord("awtawetwerwewe")

// 	var x int
// 	x = 134312321

// 	xAsBinary := fmt.Sprintf("%b", x)
// 	fmt.Println(xAsBinary)

// 	foods := [3]string{"potato", "pizza", "pasta"}

// 	for _, food := range foods {
// 		fmt.Println(food)
// 	}

// 	for i := 0; i < len(foods); i++ {
// 		fmt.Println(foods[i])
// 	}

// 	nico := person.Person{}
// 	nico.SetDetails("test", 5)
// 	fmt.Print(nico)
// 	fmt.Print(nico.Name)
// }

// func test(a int, b int, c string) (int, string) {
// 	return a + b, c
// }

// func testArgu(a ...int) int {
// 	var total int
// 	for _, item := range a {
// 		total = total + item
// 	}
// 	return total
// }

// func printWord(data string) {
// 	for _, item := range data {
// 		fmt.Printf("%c", item)
// 	}
// }

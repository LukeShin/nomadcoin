package main

import (
	"fmt"

	"github.com/LukeShin/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("2nd Block")
	chain.AddBlock("3rd Block")
	chain.AddBlock("4th Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev: %s\n", block.PrevHash)
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var blockchain *Blockchain
	args := os.Args

	// Default difficulty
	difficulty := uint64(6)

	if len(args) == 1 {
		blockchain = NewBlockchain(difficulty)
	} else if len(args) > 1 {
		var err error
		difficulty, err = strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			fmt.Println("Invalid difficulty input. Using default difficulty.")
			blockchain = NewBlockchain(difficulty)
		} else {
			blockchain, err = LoadBlockchainFromFile("blockchain.json")
			if err != nil {
				blockchain = NewBlockchain(difficulty)
			}
		}
	}

	var i = 1
	for {
		fmt.Printf("Mining block %d...\n", i)
		blockchain.AddBlock(fmt.Sprintf("Block %d test data", i))
		blockchain.SaveToFile("blockchain.json")
		i++
	}
}

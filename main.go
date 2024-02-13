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

	// Default average time duration (seconds)
	avgTimeDuration := uint64(3)

	if len(args) == 1 {
		blockchain = NewBlockchain(difficulty, avgTimeDuration)
	} else if len(args) > 1 {
		var err error
		difficulty, err = strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			fmt.Println("Invalid difficulty input. Using default difficulty.")
			blockchain = NewBlockchain(difficulty, avgTimeDuration)
		} else {
			blockchain, err = LoadBlockchainFromFile("blockchain.json")
			if err != nil {
				blockchain = NewBlockchain(difficulty, avgTimeDuration)
			}
		}
	}

	blockchain.MineBlocks(avgTimeDuration)
}

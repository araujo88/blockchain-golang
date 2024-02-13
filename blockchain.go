package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Blockchain struct {
	Chain                []Block `json:"chain"`
	difficultyAdjustment int64
	avgTimeDuration      uint64
}

func NewBlockchain(difficulty uint64, avgTimeDuration uint64) *Blockchain {
	blockchain := &Blockchain{
		Chain:           make([]Block, 0),
		avgTimeDuration: avgTimeDuration,
	}
	genesisBlock := NewBlock(0, "Genesis Block", difficulty, avgTimeDuration)
	genesisBlock.Hash = genesisBlock.calculateHash()
	blockchain.Chain = append(blockchain.Chain, *genesisBlock)
	return blockchain
}

func (bc *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(uint64(len(bc.Chain)), data, uint64(int64(bc.GetLastBlock().Difficulty)+bc.difficultyAdjustment), bc.avgTimeDuration)
	newBlock.PreviousHash = bc.GetLastBlock().GetHash()
	bc.difficultyAdjustment = newBlock.MineBlock()
	bc.Chain = append(bc.Chain, *newBlock)
}

func (bc *Blockchain) MineBlocks(avgTimeDuration uint64) {
	var i = 1
	for {
		fmt.Printf("Mining block %d...\n", i)
		bc.AddBlock(fmt.Sprintf("Block %d test data", i))
		bc.SaveToFile("blockchain.json")
		i++
	}
}

func (bc *Blockchain) GetLastBlock() *Block {
	if len(bc.Chain) == 0 {
		return nil // Or handle this case as you see fit
	}
	return &bc.Chain[len(bc.Chain)-1]
}

// Function to save Blockchain as JSON file
func (bc *Blockchain) SaveToFile(filename string) error {
	jsonData, err := json.MarshalIndent(bc, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

// LoadBlockchainFromFile loads a blockchain from a JSON file.
func LoadBlockchainFromFile(filename string) (*Blockchain, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var blockchain Blockchain
	err = json.Unmarshal(bytes, &blockchain)
	if err != nil {
		return nil, err
	}

	return &blockchain, nil
}

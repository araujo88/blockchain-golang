package main

import (
	"encoding/json"
	"io"
	"os"
)

type Blockchain struct {
	Chain      []Block `json:"chain"`
	Difficulty uint64  `json:"difficulty"`
}

func NewBlockchain(difficulty uint64) *Blockchain {
	blockchain := &Blockchain{
		Chain:      make([]Block, 0),
		Difficulty: difficulty,
	}
	genesisBlock := NewBlock(0, "Genesis Block")
	blockchain.Chain = append(blockchain.Chain, *genesisBlock)
	return blockchain
}

func (bc *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(uint64(len(bc.Chain)), data)
	newBlock.PreviousHash = bc.GetLastBlock().GetHash()
	newBlock.MineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, *newBlock)
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

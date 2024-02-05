package main

import (
	"testing"
)

// TestNewBlockchain tests the NewBlockchain function.
func TestNewBlockchain(t *testing.T) {
	difficulty := uint64(1)
	blockchain := NewBlockchain(difficulty)

	if len(blockchain.Chain) != 1 {
		t.Errorf("NewBlockchain should start with one genesis block, got %d", len(blockchain.Chain))
	}

	if blockchain.Chain[0].Data != "Genesis Block" {
		t.Errorf("Genesis block data incorrect, got %s, want Genesis Block", blockchain.Chain[0].Data)
	}
}

// TestAddBlock tests the AddBlock method of Blockchain.
func TestAddBlock(t *testing.T) {
	blockchain := NewBlockchain(1)
	blockchain.AddBlock("Test Block")

	if len(blockchain.Chain) != 2 {
		t.Errorf("AddBlock failed, expected blockchain length 2, got %d", len(blockchain.Chain))
	}

	if blockchain.Chain[1].Data != "Test Block" {
		t.Errorf("AddBlock failed to add correct data, expected 'Test Block', got '%s'", blockchain.Chain[1].Data)
	}

	if blockchain.Chain[1].PreviousHash != blockchain.Chain[0].GetHash() {
		t.Errorf("AddBlock failed to set correct previous hash")
	}
}

// TestGetLastBlock tests the GetLastBlock method of Blockchain.
func TestGetLastBlock(t *testing.T) {
	blockchain := NewBlockchain(1)
	lastBlock := blockchain.GetLastBlock()

	if lastBlock == nil {
		t.Errorf("GetLastBlock returned nil on a new blockchain")
	}

	if lastBlock != nil {
		if lastBlock.Data != "Genesis Block" {
			t.Errorf("GetLastBlock should return genesis block on new blockchain, got %s", lastBlock.Data)
		}
	}

	blockchain.AddBlock("Test Block")
	lastBlock = blockchain.GetLastBlock()

	if lastBlock.Data != "Test Block" {
		t.Errorf("GetLastBlock failed to return the last block after adding a new block, got %s", lastBlock.Data)
	}
}

// TestBlockchainValidity tests the overall validity of the blockchain.
func TestBlockchainValidity(t *testing.T) {
	blockchain := NewBlockchain(1)
	blockchain.AddBlock("Block 1")
	blockchain.AddBlock("Block 2")

	for i := 1; i < len(blockchain.Chain); i++ {
		if blockchain.Chain[i].PreviousHash != blockchain.Chain[i-1].GetHash() {
			t.Errorf("Invalid blockchain: Block %d's previous hash does not match Block %d's hash", i, i-1)
		}
	}
}

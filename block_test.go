package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestNewBlock tests the NewBlock function.
func TestNewBlock(t *testing.T) {
	data := "test data"
	block := NewBlock(0, data, 1, 2)

	if block.Data != data {
		t.Errorf("NewBlock data = %s; want %s", block.Data, data)
	}

	if block.Index != 0 {
		t.Errorf("NewBlock index = %d; want 0", block.Index)
	}

	if block.Timestamp <= 0 {
		t.Error("NewBlock timestamp is not set properly")
	}
}

// TestMineBlock tests the MineBlock function.
func TestMineBlock(t *testing.T) {
	block := NewBlock(0, "test data", 1, 2)

	block.MineBlock()

	if !strings.HasPrefix(block.GetHash(), strings.Repeat("0", int(block.GetDifficulty()))) {
		t.Errorf("MineBlock did not mine correctly with difficulty %d", int(block.GetDifficulty()))
	}

	if block.Nonce <= 0 {
		t.Error("MineBlock did not increment nonce")
	}
}

// TestCalculateHash tests the calculateHash method.
func TestCalculateHash(t *testing.T) {
	block := NewBlock(0, "test data", 4, 2)
	block.Nonce = 0
	block.PreviousHash = "0000"
	expectedHash := block.calculateHash()

	// Manually calculate expected hash
	record := strconv.FormatUint(block.Index, 10) + strconv.FormatInt(block.Timestamp, 10) + block.Data + strconv.Itoa(block.Nonce) + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	expected := fmt.Sprintf("%x", h.Sum(nil))

	if expectedHash != expected {
		t.Errorf("calculateHash = %s; want %s", expectedHash, expected)
	}
}

// TestBlockImmutableTimestamp ensures that the timestamp of a block doesn't change.
func TestBlockImmutableTimestamp(t *testing.T) {
	block := NewBlock(0, "test data", 0, 2)
	initialTimestamp := block.Timestamp

	// Simulate some delay
	time.Sleep(1 * time.Second)

	if block.Timestamp != initialTimestamp {
		t.Errorf("Block timestamp changed from %d to %d", initialTimestamp, block.Timestamp)
	}
}

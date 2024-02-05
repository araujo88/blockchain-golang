package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        uint64 `json:"index"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
	Hash         string `json:"hash"`
	Nonce        int    `json:"nonce"`
	Difficulty   uint64 `json:"difficulty"`
	PreviousHash string `json:"previousHash,omitempty"` // Omit if empty
}

func NewBlock(index uint64, data string, difficulty uint64) *Block {
	block := &Block{
		Index:      index,
		Timestamp:  time.Now().Unix(),
		Data:       data,
		Nonce:      -1,
		Difficulty: difficulty,
	}
	return block
}

func (b *Block) GetHash() string {
	return b.Hash
}

func (b *Block) GetDifficulty() uint64 {
	return b.Difficulty
}

func (b *Block) MineBlock() {
	startTime := time.Now() // Record start time

	var strBuilder strings.Builder
	for i := uint64(0); i < b.Difficulty; i++ {
		strBuilder.WriteString("0")
	}
	str := strBuilder.String()

	for {
		b.Nonce++
		b.Hash = b.calculateHash()
		fmt.Printf("%s\r\r", hexDump(b.Hash))
		if strings.HasPrefix(b.Hash, str) {
			fmt.Printf("\n*** BLOCK FOUND! ***\n\n")
			break
		}
	}

	elapsed := time.Since(startTime) // Calculate duration

	fmt.Println("Block mined:", b.Hash)
	fmt.Println("Nonce found:", b.Nonce)
	fmt.Printf("Time taken to mine the block: %s\n", elapsed)
}

func (b *Block) calculateHash() string {
	record := strconv.FormatUint(uint64(b.Index), 10) + strconv.FormatInt(b.Timestamp, 10) + b.Data + strconv.Itoa(b.Nonce) + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

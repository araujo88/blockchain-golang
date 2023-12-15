package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        uint64
	Timestamp    int64
	Data         string
	Hash         string
	Nonce        int
	PreviousHash string
}

func NewBlock(index uint64, data string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().Unix(),
		Data:      data,
		Nonce:     -1,
	}
	return block
}

func (b *Block) GetHash() string {
	return b.Hash
}

func (b *Block) MineBlock(difficulty uint64) {
	startTime := time.Now() // Record start time

	var strBuilder strings.Builder
	for i := uint64(0); i < difficulty; i++ {
		strBuilder.WriteString("0")
	}
	str := strBuilder.String()

	for {
		b.Nonce++
		b.Hash = b.calculateHash()
		fmt.Printf("%s\r\r", hexDump(b.Hash))
		if strings.HasPrefix(b.Hash, str) {
			fmt.Println("\n*** BLOCK FOUND! ***\n\n")
			break
		}
	}

	elapsed := time.Since(startTime) // Calculate duration

	fmt.Println("Block mined:", b.Hash)
	fmt.Println("Nonce found:", b.Nonce)
	fmt.Printf("Time taken to mine the block: %s\n", elapsed)

	time.Sleep(2 * time.Second) // This line might not be necessary unless you have a specific reason to keep it
}

func (b *Block) calculateHash() string {
	record := strconv.FormatUint(uint64(b.Index), 10) + strconv.FormatInt(b.Timestamp, 10) + b.Data + strconv.Itoa(b.Nonce) + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

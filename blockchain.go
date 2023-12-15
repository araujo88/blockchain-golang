package main

type Blockchain struct {
	Chain      []Block
	Difficulty uint64
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

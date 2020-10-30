package main

type Blockchain struct {
	blocks []*Block
}

// 새로운 블록체인
func NewBlockchain() *Blockchain {
	return &Blockchain{ blocks: []*Block{NewGenesisBlock()}}
}

// 구조체의 메소드
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}




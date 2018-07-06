package core

import (
	"log"
	"fmt"
)

type  BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	bc := BlockChain{}
	genesisBlock := GenerateGenesisBlock()
	bc.Blocks = append(bc.Blocks, &genesisBlock)
	return &bc
}

func (bc *BlockChain) Write(data string) bool {
	preBlock := *bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateBlock(preBlock, data)
	if !isValid(&preBlock, &newBlock) {
		return false
	} else {
		bc.Blocks = append(bc.Blocks, &newBlock)
		return true
	}
}

func (bc *BlockChain) Read() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev.Hash: %s\n", block.PreBlockHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Curr.Data: %s\n", block.Data)
		fmt.Printf("Curr.Timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}

func (bc *BlockChain) appendBlock(newBlock *Block) {
	if isValid(bc.Blocks[len(bc.Blocks) - 1], newBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

func isValid(preBlock, newBlock *Block) bool {

	if newBlock.Index - 1 != preBlock.Index {
		return false
	}
	if newBlock.PreBlockHash != preBlock.Hash {
		return false
	}
	if newBlock.Hash != Hash(*newBlock) {
		return false
	}
	return true
}

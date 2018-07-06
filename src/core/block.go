package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块
type Block struct {
	Index int64
	Timestamp int64
	PreBlockHash string
	Hash string

	Data string
}

//计算区块哈希值
func Hash(block Block) string {
	raw := string(block.Index) + string(block.Timestamp) + block.Data
	hashByte := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hashByte[:])
}

//生成区块
func GenerateBlock(preBlock Block, data string) Block {
	block := Block{}
	block.Index = preBlock.Index + 1
	block.Timestamp = time.Now().Unix()
	block.PreBlockHash = preBlock.Hash
	block.Data = data
	block.Hash = Hash(block)
	return block
}

//生成创始区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateBlock(preBlock, "Genesis Block")
}

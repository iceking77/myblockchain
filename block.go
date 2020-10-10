package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp				int64
	Data 					[]byte
	PrevBlockHash			[]byte
	Hash 					[]byte
}

// 새 블록 만드는 함수
// 타임 스템프, 머클트리 헤더 , 이전, nonce
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 해시값 도출
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 제네시스 블록을 만드는 함수
// 설정값의 난이도
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}


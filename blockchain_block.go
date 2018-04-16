package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// NewBlock with data as string
func (bc *Blockchain) NewBlock(minerWallet Wallet, transactions []*Transaction, data string,
	prevBlockHash []byte, height int, nodeID string) *Block {
	return bc.NewBlockWithData(minerWallet, transactions, []byte(data), prevBlockHash, height, nodeID)
}

// NewBlockWithData with byte array data
func (bc *Blockchain) NewBlockWithData(minerWallet Wallet, transactions []*Transaction, data []byte,
	prevBlockHash []byte, height int, nodeID string) *Block {

	if height > 0 && !bc.validateMinerWallet(nodeID) {
		fmt.Printf("Current node's wallet is not a registered wallet on the genesis block\n")
		os.Exit(1)
	}

	block := &Block{time.Now().Unix(), transactions, data, prevBlockHash, []byte{}, height, []byte{}, []byte{}}
	block.SetHash()
	block.SignByMiner(minerWallet, getMinerAddressFromNodeID(nodeID))
	return block
}

func (bc *Blockchain) validateMinerWallet(nodeID string) bool {
	genesis, err := bc.GetBlockByHeight(0)
	if err != nil {
		fmt.Printf("No genesis block. Please create a blockchain first")
		os.Exit(1)
	}
	var xnodes []XNode
	json.Unmarshal(genesis.Data, &xnodes)

	minerAddress := getMinerAddressFromNodeID(nodeID)
	return ValidateMiner(minerAddress, nodeID, xnodes)
}

func getMinerAddressFromNodeID(nodeID string) string {
	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	return wallets.GetAddresses()[0]
}

// NewGenesisBlock creates and returns genesis Block
func (bc *Blockchain) NewGenesisBlock(nodeID string, address string, coinbase *Transaction) *Block {
	xnodes := getXNodes()
	encodedNodes := toJSON(xnodes)

	wallet := GetMinerWallet(nodeID, address)

	if !ValidateMiner(address, nodeID, xnodes) {
		log.Panic("Invalid miner")
	}
	return bc.NewBlockWithData(wallet, []*Transaction{coinbase},
		[]byte(encodedNodes), []byte{}, 0, nodeID)
}

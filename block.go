package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// AuthorizedNode keeps XNode info
type AuthorizedNode struct {
	IP []byte
}

// XNode contains miners list
type XNode struct {
	URL     string `json:"url"`
	Address string `json:"address"`
}

var genesisXNodes []XNode

// Block keeps block headers
type Block struct {
	Timestamp      int64
	Transactions   []*Transaction
	Data           ByteString
	PrevBlockHash  ByteString
	Hash           ByteString
	Height         int
	MinerSignature ByteString
	MinerPubkey    ByteString
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	hashTrx := b.HashTransactions()
	headers := bytes.Join([][]byte{
		b.PrevBlockHash,
		hashTrx,
		b.Data,
		IntToHex(b.Timestamp),
	}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// SignByMiner signs the block with miner's private key
func (b *Block) SignByMiner(minerWallet Wallet, minerAddress string) {
	r, s, err := ecdsa.Sign(rand.Reader, &minerWallet.PrivateKey, []byte(b.Hash))
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)

	b.MinerSignature = signature
	b.MinerPubkey = getPubkeyHashFromAddress([]byte(minerAddress))
}

func getXNodes() []XNode {
	// Open our jsonFile
	file, e := ioutil.ReadFile("./xnodes.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var xnodes []XNode
	json.Unmarshal(file, &xnodes)

	return xnodes
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func getBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getInterface(bts []byte, data interface{}) error {
	buf := bytes.NewBuffer(bts)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

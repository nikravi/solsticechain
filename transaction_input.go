package main

import "bytes"

// TXInput represents a transaction input
type TXInput struct {
	Txid      ByteString
	Vout      int
	Signature ByteString
	PubKey    ByteString
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

package main

import (
	"log"
)

// ValidateMiner validates the address
func ValidateMiner(pubkey string, nodeID string, xnodes []XNode) bool {
	return ValidateAddress(pubkey) && containsNode(xnodes, XNode{FormatNodeURL(nodeID), pubkey})
}

// GetMinerWallet gets the wallet
func GetMinerWallet(nodeID string, pubkey string) Wallet {
	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	return wallets.GetWallet(pubkey)
}

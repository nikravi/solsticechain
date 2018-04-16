package main

import (
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address string, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()
	UTXOSet := UTXOSet{bc}

	balance := 0
	pubKeyHash := getPubkeyHashFromAddress([]byte(address))
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

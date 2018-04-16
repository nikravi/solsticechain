package main

import (
	"errors"
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, data string, nodeID string, mineNow bool) {
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()
	_, err := addTransactionToBlockchain(from, to, amount, data, nodeID, mineNow, bc)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Success!")
}
func addTransactionToBlockchain(from, to string, amount int, data string,
	nodeID string, mineNow bool, bc *Blockchain) (Transaction, error) {
	var transaction Transaction
	var err error

	if !ValidateAddress(from) {
		err = errors.New("ERROR: Sender address is not valid")
		return transaction, err
	}
	if !ValidateAddress(to) {
		err = errors.New("ERROR: Recipient address is not valid")
		return transaction, err
	}

	UTXOSet := UTXOSet{bc}

	wallets, e := NewWallets(nodeID)
	if e != nil {
		err = e
		return transaction, err
	}
	wallet := wallets.GetWallet(from)

	transaction = *NewUTXOTransaction(&wallet, to, amount, data, &UTXOSet)

	if mineNow {
		cbTx := NewCoinbaseTX(from, "mined ad-hoc")
		txs := []*Transaction{cbTx, &transaction}

		newBlock := bc.MineBlock(wallet, txs, "data set by the current node (as miner)", nodeID)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], &transaction)
	}

	return transaction, nil
}

package main

import (
	"fmt"
)

func (cli *CLI) printChain(nodeID string) {
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("***** Block hash %x *****\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. block:          %x\n", block.PrevBlockHash)
		fmt.Printf("Block Data as string: %s\n", block.Data)
		fmt.Printf("Block MinerSignature: %x\n", block.MinerSignature)
		fmt.Printf("Block MinnerPubkey: %x\n", block.MinerPubkey)
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
